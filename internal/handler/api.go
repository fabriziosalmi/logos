package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/fabriziosalmi/logos/internal/config"
	"github.com/fabriziosalmi/logos/internal/svg"
	"github.com/go-chi/chi/v5"
)

var svgConfig struct {
	a11y config.A11yConfig
	etag bool
}

func InitHandlers(cfg *config.Config) {
	svgConfig.a11y = cfg.Accessibility
	svgConfig.etag = cfg.Cache.ETag
}

func wantsSaveData(r *http.Request) bool {
	return svgConfig.a11y.RespectSaveData && r.Header.Get("Save-Data") == "on"
}

// parseQueryParams extracts common parametric controls from query string.
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

func parseShape(r *http.Request) string {
	return r.URL.Query().Get("shape")
}

func writeSVG(w http.ResponseWriter, r *http.Request, result svg.RenderResult) {
	if svgConfig.etag && result.ETag != "" {
		w.Header().Set("ETag", result.ETag)
		if match := r.Header.Get("If-None-Match"); match == result.ETag {
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Vary", "Accept-Encoding, Save-Data")
	w.WriteHeader(http.StatusOK)
	w.Write(result.SVG)
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

	result := svg.Render(svg.RenderOptions{
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
	})

	writeSVG(w, r, result)
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

	result := svg.RenderFavicon(
		color1, color2, animation, theme,
		svgConfig.a11y.GenerateTitle,
		svgConfig.a11y.RespectReducedMotion,
	)

	writeSVG(w, r, result)
}

// App handles /app/{name}/{format}.svg with fallback placeholder for unknown apps.
func App(apps map[string]*config.AppConfig, defaults config.DefaultsConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		format := strings.TrimSuffix(chi.URLParam(r, "format"), ".svg")

		app, ok := apps[name]
		if !ok {
			// Fallback: return a placeholder instead of 404
			result := svg.RenderPlaceholder()
			w.Header().Set("X-Logos-Fallback", "true")
			writeSVG(w, r, result)
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

		// Query params can override app defaults
		if qs := parseShape(r); qs != "" {
			shape = qs
		}

		stroke, padding, alpha, badge, variant, hover, decorative := parseQueryParams(r)

		result := svg.Render(svg.RenderOptions{
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
		})

		writeSVG(w, r, result)
	}
}

// Health returns a simple health check.
func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}
