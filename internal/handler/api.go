package handler

import (
	"net/http"
	"strconv"
	"strings"

	"encoding/json"
	"sort"

	"github.com/fabriziosalmi/logos/internal/gen"
	"github.com/fabriziosalmi/logos/internal/cache"
	"github.com/fabriziosalmi/logos/internal/config"
	"github.com/fabriziosalmi/logos/internal/svg"
	"github.com/go-chi/chi/v5"
)

var svgConfig struct {
	a11y  config.A11yConfig
	etag  bool
	cache *cache.Tiered
}

func InitHandlers(cfg *config.Config, tc *cache.Tiered) {
	svgConfig.a11y = cfg.Accessibility
	svgConfig.etag = cfg.Cache.ETag
	svgConfig.cache = tc
}

func wantsSaveData(r *http.Request) bool {
	return svgConfig.a11y.RespectSaveData && r.Header.Get("Save-Data") == "on"
}

func parseQueryParams(r *http.Request) (stroke, padding, alpha float64, badge int, variant string, hover, decorative bool) {
	q := r.URL.Query()
	if v := q.Get("stroke"); v != "" {
		stroke, _ = strconv.ParseFloat(v, 64)
		if stroke < 0 || stroke > 10 {
			stroke = 0
		}
	}
	if v := q.Get("padding"); v != "" {
		padding, _ = strconv.ParseFloat(v, 64)
		if padding < 0 || padding > 20 {
			padding = 0
		}
	}
	if v := q.Get("alpha"); v != "" {
		alpha, _ = strconv.ParseFloat(v, 64)
	}
	if v := q.Get("badge"); v != "" {
		badge, _ = strconv.Atoi(v)
		if badge < 0 || badge > 999 {
			badge = 0
		}
	}
	variant = q.Get("variant")
	hover = q.Get("hover") == "true"
	decorative = q.Get("decorative") == "true"
	return
}

func parseSpeed(r *http.Request) float64 {
	if v := r.URL.Query().Get("speed"); v != "" {
		s, _ := strconv.ParseFloat(v, 64)
		if s >= 0.1 && s <= 10 {
			return s
		}
	}
	return 1.0
}

func parseDirection(r *http.Request) string {
	d := r.URL.Query().Get("direction")
	switch d {
	case "reverse", "alternate", "alternate-reverse":
		return d
	}
	return ""
}

func parseShape(r *http.Request) string {
	return r.URL.Query().Get("shape")
}

// writeSVG writes an SVG response with ETag + cache tier headers.
func writeSVG(w http.ResponseWriter, r *http.Request, svgBytes []byte, etag string, tier string, appSlug string) {
	// CDN headers
	if svgConfig.cache != nil {
		svgConfig.cache.SetCDNHeaders(w, tier, appSlug)
	}

	// ETag / 304
	if svgConfig.etag && etag != "" {
		w.Header().Set("ETag", etag)
		if match := r.Header.Get("If-None-Match"); match == etag {
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Vary", "Accept-Encoding, Save-Data")
	w.WriteHeader(http.StatusOK)
	w.Write(svgBytes)
}

// renderCached checks all cache tiers, renders on miss, stores result.
func renderCached(opts svg.RenderOptions, r *http.Request, appSlug string) ([]byte, string, string) {
	key := cache.KeyFromRequest(
		opts.Format, opts.Scene, opts.Color1, opts.Color2,
		opts.Animation, opts.Theme, opts.Shape, opts.Texture, opts.Variant,
		opts.Text.Title, opts.Text.Subtitle, opts.Direction,
		opts.StrokeWidth, opts.Padding, opts.Alpha, opts.Speed,
		opts.Badge, opts.Hover, opts.AriaHidden,
	)

	// Check cache
	if svgConfig.cache != nil {
		if svgBytes, etag, tier := svgConfig.cache.Get(key); tier != "miss" {
			return svgBytes, etag, tier
		}
	}

	// Cache miss: render
	result := svg.Render(opts)

	// Store in all tiers
	if svgConfig.cache != nil {
		svgConfig.cache.Set(key, result.SVG, result.ETag)
	}

	return result.SVG, result.ETag, "render"
}

// V4 handles /api/v4/render/{format}/{scene}/{color1}/{animation}.svg
func V4(w http.ResponseWriter, r *http.Request) {
	format := chi.URLParam(r, "format")
	scene := chi.URLParam(r, "scene")
	color1 := chi.URLParam(r, "color1")
	color2 := chi.URLParam(r, "color2")
	animation := strings.TrimSuffix(chi.URLParam(r, "animation"), ".svg")

	theme := r.URL.Query().Get("theme")
	if theme == "" {
		theme = "auto"
	}
	texture := r.URL.Query().Get("texture")
	shape := parseShape(r)
	stroke, padding, alpha, badge, variant, hover, decorative := parseQueryParams(r)

	opts := svg.RenderOptions{
		Format:      format,
		Scene:       scene,
		Color1:      color1,
		Color2:      color2,
		Animation:   animation,
		Theme:       theme,
		Texture:     texture,
		Shape:       shape,
		StrokeWidth: stroke,
		Padding:     padding,
		Alpha:       alpha,
		Badge:       badge,
		Variant:     variant,
		Speed:       parseSpeed(r),
		Direction:   parseDirection(r),
		Text: svg.TextParams{
			Title:    r.URL.Query().Get("title"),
			Subtitle: r.URL.Query().Get("subtitle"),
		},
		A11yTitle:     svgConfig.a11y.GenerateTitle && !decorative,
		A11yDesc:      svgConfig.a11y.GenerateDesc && !decorative,
		AriaHidden:    decorative,
		ReducedMotion: svgConfig.a11y.RespectReducedMotion,
		SaveData:      wantsSaveData(r),
		Hover:         hover,
	}

	svgBytes, etag, tier := renderCached(opts, r, "")
	writeSVG(w, r, svgBytes, etag, tier, "")
}

// V3 handles legacy /api/favicon/{color1}/{animation}.svg
func V3(w http.ResponseWriter, r *http.Request) {
	color1 := chi.URLParam(r, "color1")
	color2 := chi.URLParam(r, "color2")
	animation := strings.TrimSuffix(chi.URLParam(r, "animation"), ".svg")

	theme := r.URL.Query().Get("theme")
	if theme == "" {
		theme = "auto"
	}

	opts := svg.RenderOptions{
		Format:        "favicon",
		Scene:         "pure",
		Color1:        color1,
		Color2:        color2,
		Animation:     animation,
		Theme:         theme,
		Shape:         "atom",
		A11yTitle:     svgConfig.a11y.GenerateTitle,
		A11yDesc:      svgConfig.a11y.GenerateDesc,
		ReducedMotion: svgConfig.a11y.RespectReducedMotion,
	}

	svgBytes, etag, tier := renderCached(opts, r, "")
	writeSVG(w, r, svgBytes, etag, tier, "")
}

// App handles /app/{name}/{format}.svg with fallback placeholder.
func App(apps map[string]*config.AppConfig, defaults config.DefaultsConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		format := strings.TrimSuffix(chi.URLParam(r, "format"), ".svg")

		app, ok := apps[name]
		if !ok {
			result := svg.RenderPlaceholder()
			w.Header().Set("X-Logos-Fallback", "true")
			writeSVG(w, r, result.SVG, result.ETag, "render", "")
			return
		}

		animation := app.Animation
		if animation == "" {
			animation = defaults.Animation
		}
		theme := app.Theme
		if theme == "" {
			theme = defaults.Theme
		}
		scene := app.Scene
		if scene == "" {
			scene = defaults.Scene
		}
		texture := app.Texture
		if texture == "" {
			texture = defaults.Texture
		}
		shape := app.Shape
		if shape == "" {
			shape = defaults.Shape
		}
		if format == "" {
			format = defaults.Format
		}

		if qs := parseShape(r); qs != "" {
			shape = qs
		}

		stroke, padding, alpha, badge, variant, hover, decorative := parseQueryParams(r)

		opts := svg.RenderOptions{
			Format:      format,
			Scene:       scene,
			Color1:      app.Color,
			Color2:      app.Color2,
			Animation:   animation,
			Theme:       theme,
			Texture:     texture,
			Shape:       shape,
			StrokeWidth: stroke,
			Padding:     padding,
			Alpha:       alpha,
			Badge:       badge,
			Variant:     variant,
			Text: svg.TextParams{
				Title:    app.Title,
				Subtitle: app.Subtitle,
			},
			A11yTitle:     svgConfig.a11y.GenerateTitle && !decorative,
			A11yDesc:      svgConfig.a11y.GenerateDesc && !decorative,
			AriaHidden:    decorative,
			ReducedMotion: svgConfig.a11y.RespectReducedMotion,
			SaveData:      wantsSaveData(r),
			Hover:         hover,
			AppRole:       app.Subtitle,
		}

		svgBytes, etag, tier := renderCached(opts, r, name)
		writeSVG(w, r, svgBytes, etag, tier, name)
	}
}

// Health returns health check.
func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}

// IconList returns all available icon pack names as JSON.
func IconList(w http.ResponseWriter, r *http.Request) {
	pack := r.URL.Query().Get("pack") // "lucide", "hero", or "" for all
	search := r.URL.Query().Get("q")

	allNames := svg.IconPackNames()

	var filtered []string
	for _, name := range allNames {
		if pack != "" && !strings.HasPrefix(name, pack+":") {
			continue
		}
		if search != "" && !strings.Contains(name, strings.ToLower(search)) {
			continue
		}
		filtered = append(filtered, name)
	}

	// Sort for deterministic output
	sort.Strings(filtered)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"count": len(filtered),
		"icons": filtered,
	})
}

// CacheStats returns cache tier stats as JSON.
func CacheStats(w http.ResponseWriter, r *http.Request) {
	if svgConfig.cache == nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{}`))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(svgConfig.cache.StatsJSON())
}

// GenRender handles /api/v4/gen/render?prompt=...
// Generates a unique procedural SVG shape from a text prompt via
// mathematical generators (hash-seeded), styled through the standard pipeline.
func GenRender(w http.ResponseWriter, r *http.Request) {
	prompt := r.URL.Query().Get("prompt")
	if prompt == "" {
		http.Error(w, `{"error":"prompt required"}`, http.StatusBadRequest)
		return
	}

	// Resolve prompt to params (semantic mapping)
	resolved := gen.ResolvePrompt(prompt)

	// Generate unique procedural geometry from prompt
	generatedPath := gen.GeneratePath(prompt)

	// Allow query param overrides (user can fine-tune AI suggestions)
	format := r.URL.Query().Get("format")
	if format == "" {
		format = "favicon"
	}

	// ?generate=false skips procedural path, uses semantic shape instead
	useGenerated := r.URL.Query().Get("generate") != "false"

	// Override resolved params with explicit query params if provided
	if v := r.URL.Query().Get("shape"); v != "" {
		resolved.Shape = v
	}
	if v := r.URL.Query().Get("color"); v != "" {
		resolved.Color = v
	}
	if v := r.URL.Query().Get("animation"); v != "" {
		resolved.Animation = v
	}
	if v := r.URL.Query().Get("scene"); v != "" {
		resolved.Scene = v
	}
	if v := r.URL.Query().Get("theme"); v != "" {
		resolved.Theme = v
	}
	if v := r.URL.Query().Get("texture"); v != "" {
		resolved.Texture = v
	}
	if v := r.URL.Query().Get("variant"); v != "" {
		resolved.Variant = v
	}

	stroke, padding, alpha, badge, _, hover, decorative := parseQueryParams(r)

	customPath := ""
	if useGenerated {
		customPath = generatedPath
	}

	opts := svg.RenderOptions{
		Format:      format,
		Scene:       resolved.Scene,
		Color1:      resolved.Color,
		Color2:      resolved.Color2,
		Animation:   resolved.Animation,
		Theme:       resolved.Theme,
		Texture:     resolved.Texture,
		Shape:       resolved.Shape,
		Variant:     resolved.Variant,
		CustomPath:  customPath,
		StrokeWidth: stroke,
		Padding:     padding,
		Alpha:       alpha,
		Badge:       badge,
		Speed:       parseSpeed(r),
		Direction:   parseDirection(r),
		Text: svg.TextParams{
			Title:    r.URL.Query().Get("title"),
			Subtitle: r.URL.Query().Get("subtitle"),
		},
		A11yTitle:     svgConfig.a11y.GenerateTitle && !decorative,
		A11yDesc:      svgConfig.a11y.GenerateDesc && !decorative,
		AriaHidden:    decorative,
		ReducedMotion: svgConfig.a11y.RespectReducedMotion,
		SaveData:      wantsSaveData(r),
		Hover:         hover,
	}

	svgBytes, etag, tier := renderCached(opts, r, "")

	// Set extra header showing what the AI resolved
	w.Header().Set("X-Logos-Gen-Resolved", resolved.Shape+"/"+resolved.Color+"/"+resolved.Animation)
	writeSVG(w, r, svgBytes, etag, tier, "")
}

// GenResolve returns the resolved params as JSON (debug/preview endpoint).
func GenResolve(w http.ResponseWriter, r *http.Request) {
	prompt := r.URL.Query().Get("prompt")
	if prompt == "" {
		http.Error(w, `{"error":"prompt required"}`, http.StatusBadRequest)
		return
	}
	resolved := gen.ResolvePrompt(prompt)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resolved)
}

// CachePurge clears all cache tiers.
func CachePurge(w http.ResponseWriter, r *http.Request) {
	if svgConfig.cache != nil {
		svgConfig.cache.Purge()
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"purged":true}`))
}
