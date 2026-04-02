package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/fabriziosalmi/logos/internal/cache"
	"github.com/fabriziosalmi/logos/internal/config"
	"github.com/fabriziosalmi/logos/internal/handler"
	"github.com/fabriziosalmi/logos/internal/middleware"
	"github.com/fabriziosalmi/logos/internal/svg"
	"github.com/fabriziosalmi/logos/web"
	"github.com/go-chi/chi/v5"
)

const version = "1.0.0"

func resolveConfigPath(primary string) string {
	if primary == "" {
		primary = "config.yaml"
	}
	if _, err := os.Stat(primary); err == nil {
		return primary
	}
	if _, err := os.Stat("config.sample.yaml"); err == nil {
		return "config.sample.yaml"
	}
	return primary
}

func main() {
	// CLI subcommands
	if len(os.Args) > 1 && os.Args[1] == "render" {
		cliRender()
		return
	}
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Println(version)
		return
	}

	serve()
}

// cliRender handles: logos render [flags] or logos render <app-name>
func cliRender() {
	fs := flag.NewFlagSet("render", flag.ExitOnError)
	fFormat := fs.String("format", "favicon", "Output format: favicon, avatar, og-card, hero")
	fScene := fs.String("scene", "pure", "Scene: pure, spotlight, grid, split")
	fColor := fs.String("color", "", "Primary color (name or hex)")
	fColor2 := fs.String("color2", "", "Gradient color (optional)")
	fAnimation := fs.String("animation", "static", "Animation name")
	fTheme := fs.String("theme", "dark", "Theme: dark, light, solid, glass, auto")
	fTexture := fs.String("texture", "", "Texture: none, grain, glass, noise, glitch, shadow, neon")
	fShape := fs.String("shape", "atom", "Shape: atom, shield, hexagon, diamond, bolt, cube, wave, gear, eye, leaf, star, circle, triangle, square")
	fTitle := fs.String("title", "", "Title text")
	fSubtitle := fs.String("subtitle", "", "Subtitle text")
	fStroke := fs.Float64("stroke", 0, "Stroke width override")
	fBadge := fs.Int("badge", 0, "Badge notification number")
	fConfig := fs.String("config", "config.yaml", "Config file path")

	fs.Parse(os.Args[2:])

	// If first non-flag arg is an app name, render from config
	args := fs.Args()
	if len(args) > 0 && *fColor == "" {
		renderApp(args[0], *fFormat, *fConfig)
		return
	}

	if *fColor == "" {
		*fColor = "white"
	}

	result := svg.Render(svg.RenderOptions{
		Format:      *fFormat,
		Scene:       *fScene,
		Color1:      *fColor,
		Color2:      *fColor2,
		Animation:   *fAnimation,
		Theme:       *fTheme,
		Texture:     *fTexture,
		Shape:       *fShape,
		StrokeWidth: *fStroke,
		Badge:       *fBadge,
		Text:        svg.TextParams{Title: *fTitle, Subtitle: *fSubtitle},
		A11yTitle:   true,
		A11yDesc:    true,
		ReducedMotion: true,
	})
	os.Stdout.Write(result.SVG)
}

func renderApp(name, format, cfgPath string) {
	cfg, err := config.Load(resolveConfigPath(cfgPath))
	if err != nil {
		fmt.Fprintf(os.Stderr, "config error: %v\n", err)
		os.Exit(1)
	}

	app, ok := cfg.Apps[name]
	if !ok {
		// List available apps
		names := make([]string, 0, len(cfg.Apps))
		for k := range cfg.Apps {
			names = append(names, k)
		}
		fmt.Fprintf(os.Stderr, "app %q not found. available: %s\n", name, strings.Join(names, ", "))
		os.Exit(1)
	}

	anim := app.Animation
	if anim == "" {
		anim = cfg.Defaults.Animation
	}
	theme := app.Theme
	if theme == "" {
		theme = cfg.Defaults.Theme
	}
	scene := app.Scene
	if scene == "" {
		scene = cfg.Defaults.Scene
	}
	texture := app.Texture
	if texture == "" {
		texture = cfg.Defaults.Texture
	}
	if format == "" {
		format = cfg.Defaults.Format
	}

	result := svg.Render(svg.RenderOptions{
		Format:        format,
		Scene:         scene,
		Color1:        app.Color,
		Color2:        app.Color2,
		Animation:     anim,
		Theme:         theme,
		Texture:       texture,
		Text:          svg.TextParams{Title: app.Title, Subtitle: app.Subtitle},
		A11yTitle:     true,
		A11yDesc:      true,
		ReducedMotion: true,
		AppRole:       app.Subtitle,
	})
	os.Stdout.Write(result.SVG)
}

func serve() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	cfgPath := "config.yaml"
	if v := os.Getenv("CONFIG_PATH"); v != "" {
		cfgPath = v
	}
	cfgPath = resolveConfigPath(cfgPath)

	cfg, err := config.Load(cfgPath)
	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	// Validate config
	if err := config.Validate(cfg); err != nil {
		slog.Error("invalid config", "error", err)
		os.Exit(1)
	}

	// Load icon packs (Lucide 1694 + Heroicons 324)
	svg.LoadIconPack("lucide", web.IconsFS, "icons/lucide")
	svg.LoadIconPack("hero", web.IconsFS, "icons/hero")

	// Init multi-tier cache
	tc := cache.NewTiered(cfg.Cache)

	handler.InitHandlers(cfg, tc)

	// Pre-warm: render all registered app favicons into all cache tiers
	for name, app := range cfg.Apps {
		anim := app.Animation
		if anim == "" {
			anim = cfg.Defaults.Animation
		}
		theme := app.Theme
		if theme == "" {
			theme = cfg.Defaults.Theme
		}
		shape := app.Shape
		if shape == "" {
			shape = cfg.Defaults.Shape
		}
		opts := svg.RenderOptions{
			Format:    "favicon",
			Scene:     "pure",
			Color1:    app.Color,
			Animation: anim,
			Theme:     theme,
			Shape:     shape,
			A11yTitle: true,
			Text:      svg.TextParams{Title: app.Title},
		}
		result := svg.Render(opts)
		key := cache.KeyFromRequest(
			opts.Format, opts.Scene, opts.Color1, "", opts.Animation,
			opts.Theme, opts.Shape, "", "", opts.Text.Title, "", "",
			0, 0, 0, 1.0, 0, false, false,
		)
		tc.Set(key, result.SVG, result.ETag)
		slog.Info("pre-warmed", "app", name)
	}

	r := chi.NewRouter()

	// Global middleware
	r.Use(middleware.Recover)
	r.Use(middleware.Logger(logger))
	r.Use(middleware.CORS(cfg.CORS))
	r.Use(middleware.Security(cfg.Security))
	r.Use(middleware.RateLimit(cfg.Security.RateLimit))

	if cfg.Delivery.Compression.Gzip {
		r.Use(middleware.Gzip)
	}

	r.Get("/healthz", handler.Health)
	r.Get("/api/v4/icons", handler.IconList)

	r.With(middleware.AdminOnly(cfg.Server.AdminKey)).Get("/cache/stats", handler.CacheStats)
	r.With(middleware.AdminOnly(cfg.Server.AdminKey)).Post("/cache/purge", handler.CachePurge)

	// Generative: prompt → procedural SVG → SVG
	r.Route("/api/v4/gen", func(r chi.Router) {
		r.Use(middleware.CacheSWR(cfg.Cache))
		r.Get("/render", handler.GenRender)     // ?prompt=cyberpunk+hacker+glowing&format=avatar
		r.Get("/resolve", handler.GenResolve)   // ?prompt=... → JSON params (debug)
	})

	r.Route("/api/v4/render", func(r chi.Router) {
		r.Use(middleware.CacheSWR(cfg.Cache))
		r.Get("/{format}/{scene}/{color1}/{animation}", handler.V4)
		r.Get("/{format}/{scene}/{color1}/{color2}/{animation}", handler.V4)
	})

	r.Route("/api/favicon", func(r chi.Router) {
		r.Use(middleware.CacheSWR(cfg.Cache))
		r.Get("/{color1}/{animation}", handler.V3)
		r.Get("/{color1}/{color2}/{animation}", handler.V3)
	})

	if len(cfg.Apps) > 0 {
		r.Route("/app", func(r chi.Router) {
			r.Use(middleware.CacheSWR(cfg.Cache))
			appHandler := handler.App(cfg.Apps, cfg.Defaults)
			r.Get("/{name}/{format}", appHandler)
		})
	}

	if cfg.Cache.ImmutableStatic {
		r.With(middleware.CacheImmutable(cfg.Cache.StaleWhileRevalidate)).
			Handle("/static/*", http.StripPrefix("/static/", handler.StaticFiles(web.FS)))
	} else {
		r.Handle("/static/*", http.StripPrefix("/static/", handler.StaticFiles(web.FS)))
	}

	dashboardHandler := handler.Dashboard(web.FS)
	r.Get("/", dashboardHandler)
	r.Get("/css/*", dashboardHandler)
	r.Get("/js/*", dashboardHandler)

	srv := &http.Server{
		Addr:         cfg.Addr(),
		Handler:      r,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)

	go func() {
		slog.Info("logos api starting",
			"addr", cfg.Addr(),
			"version", version,
			"apps", len(cfg.Apps),
		)
		fmt.Printf("\n  Logos API v%s\n  http://%s\n\n", version, cfg.Addr())
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("server error", "error", err)
			os.Exit(1)
		}
	}()

	<-done
	slog.Info("shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Server.ShutdownTimeout)
	defer cancel()
	srv.Shutdown(ctx)
}
