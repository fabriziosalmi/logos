package svg

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sync"
)

var bufPool = sync.Pool{
	New: func() any { return new(bytes.Buffer) },
}

// RenderOptions configures the cinematic SVG renderer.
type RenderOptions struct {
	Format    string
	Scene     string
	Color1    string
	Color2    string
	Animation string
	Theme     string
	Texture   string
	Shape     string // atom, shield, hexagon, diamond, bolt, cube, wave, gear, eye, leaf, star, circle, triangle, square
	Text      TextParams

	// Parametric controls
	StrokeWidth float64
	Padding     float64
	Alpha       float64
	Badge       int
	Variant     string

	// A11y & empathy flags
	A11yTitle     bool
	A11yDesc      bool
	AriaHidden    bool
	ReducedMotion bool
	SaveData      bool
	Hover         bool
	AppRole       string
}

// RenderResult holds the SVG bytes and its content hash for ETag.
type RenderResult struct {
	SVG  []byte
	ETag string
}

// Render generates a complete cinematic SVG.
func Render(opts RenderOptions) RenderResult {
	buf := bufPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufPool.Put(buf)

	dim := Formats[opts.Format]
	if dim.Width == 0 {
		dim = Formats["favicon"]
	}

	primaryHex := ResolveHex(opts.Color1)
	secondaryHex := primaryHex
	if opts.Color2 != "" {
		secondaryHex = ResolveHex(opts.Color2)
	}

	hasGradient := secondaryHex != primaryHex

	fillDef := "#" + primaryHex
	strokeDef := "#" + primaryHex
	if hasGradient {
		fillDef = "url(#gradCore)"
		strokeDef = "url(#gradCore)"
	}

	animCss := Animations[opts.Animation]
	theme := ResolveTheme(opts.Theme, fillDef)

	outerStrokeWidth := theme.StrokeWidth
	if opts.StrokeWidth > 0 {
		outerStrokeWidth = fmt.Sprintf("%.1f", opts.StrokeWidth)
	}
	orbitStrokeWidth := "1.5"
	if opts.StrokeWidth > 0 {
		orbitStrokeWidth = fmt.Sprintf("%.1f", opts.StrokeWidth)
	}

	if opts.Variant == "outline" {
		theme.CoreBg = "transparent"
		if outerStrokeWidth == "0" {
			outerStrokeWidth = "1.5"
		}
	}

	alpha := opts.Alpha
	if alpha <= 0 || alpha > 1 {
		alpha = 1.0
	}

	texture := opts.Texture
	if opts.SaveData || texture == "" {
		texture = "none"
	}

	shape := opts.Shape
	if shape == "" {
		shape = "atom"
	}

	pad := opts.Padding
	vbX, vbY := 0.0, 0.0
	vbW, vbH := float64(dim.Width), float64(dim.Height)
	if pad > 0 {
		vbX = -pad
		vbY = -pad
		vbW += pad * 2
		vbH += pad * 2
	}

	// SVG root
	if opts.AriaHidden {
		fmt.Fprintf(buf, `<svg xmlns="http://www.w3.org/2000/svg" viewBox="%.0f %.0f %.0f %.0f" preserveAspectRatio="xMidYMid meet" aria-hidden="true">`,
			vbX, vbY, vbW, vbH)
	} else {
		fmt.Fprintf(buf, `<svg xmlns="http://www.w3.org/2000/svg" viewBox="%.0f %.0f %.0f %.0f" preserveAspectRatio="xMidYMid meet" role="img" aria-labelledby="logo-title logo-desc">`,
			vbX, vbY, vbW, vbH)
	}

	// A11y tags
	if !opts.AriaHidden {
		if opts.A11yTitle && opts.Text.Title != "" {
			fmt.Fprintf(buf, "\n  <title id=\"logo-title\">%s</title>", opts.Text.Title)
		} else if opts.A11yTitle {
			buf.WriteString("\n  <title id=\"logo-title\">Logos API</title>")
		}
		if opts.A11yDesc {
			desc := "Dynamic SVG logo"
			if opts.Text.Title != "" {
				desc = opts.Text.Title + " logo"
			}
			if opts.AppRole != "" {
				desc += " — " + opts.AppRole
			}
			fmt.Fprintf(buf, "\n  <desc id=\"logo-desc\">%s</desc>", desc)
		}
	}

	// Defs
	buf.WriteString("\n  <defs>")
	if hasGradient {
		fmt.Fprintf(buf, `
    <linearGradient id="gradCore" x1="0%%" y1="0%%" x2="100%%" y2="100%%">
      <stop offset="0%%" stop-color="#%s"/>
      <stop offset="100%%" stop-color="#%s"/>
    </linearGradient>`, primaryHex, secondaryHex)
	}
	buf.WriteString(SceneDefs(opts.Scene, primaryHex, secondaryHex))
	buf.WriteString(TextureDefs(texture))
	buf.WriteString("\n  </defs>")

	// Style
	buf.WriteString("\n  <style>")
	buf.WriteString("\n    :root{--ease:cubic-bezier(.25,1,.5,1);--ease-smooth:cubic-bezier(.65,0,.35,1)}")
	buf.WriteString("\n    .outer,.orbit,.core{transform-origin:16px 16px}")

	if animCss != "" {
		fmt.Fprintf(buf, "\n    %s", animCss)
		if opts.ReducedMotion {
			buf.WriteString("\n    @media(prefers-reduced-motion:reduce){.outer,.orbit,.core{animation:none!important}}")
		}
	}

	buf.WriteString("\n    @media(forced-colors:active){.outer,.orbit{stroke:CanvasText}.core circle,.core polygon{fill:CanvasText}.dot{fill:Canvas}}")

	if opts.Hover {
		buf.WriteString("\n    .logos-api-core:hover{filter:brightness(1.2);transition:filter .2s ease}")
	}

	buf.WriteString("\n  </style>")

	// Background
	buf.WriteString("\n  ")
	buf.WriteString(SceneBackground(opts.Scene))

	// Alpha wrapper
	if alpha < 1.0 {
		fmt.Fprintf(buf, "\n  <g opacity=\"%.2f\">", alpha)
	}

	// Transform + texture wrapper
	transform := logoTransform(opts.Format, opts.Scene, dim)
	filterAttr := TextureFilterAttr(texture)

	if transform != "" || filterAttr != "" {
		buf.WriteString("\n  <g")
		if transform != "" {
			fmt.Fprintf(buf, ` transform="%s"`, transform)
		}
		if filterAttr != "" {
			buf.WriteString(filterAttr)
		}
		buf.WriteString(">")
	}

	// THE SHAPE — this is where it all comes together
	shapeParams := ShapeParams{
		CoreBg:           theme.CoreBg,
		FillDef:          fillDef,
		StrokeDef:        strokeDef,
		DotColor:         theme.DotColor,
		OuterStrokeWidth: outerStrokeWidth,
		OrbitStrokeWidth: orbitStrokeWidth,
	}
	buf.WriteString(RenderShape(shape, shapeParams))

	if transform != "" || filterAttr != "" {
		buf.WriteString("\n  </g>")
	}

	// Badge
	if opts.Badge > 0 && opts.Badge <= 999 {
		renderBadge(buf, opts.Badge, dim, opts.Format)
	}

	// Typography
	buf.WriteString(RenderTypography(opts.Text, opts.Format, opts.Scene, primaryHex))

	if alpha < 1.0 {
		buf.WriteString("\n  </g>")
	}

	buf.WriteString("\n</svg>")

	out := make([]byte, buf.Len())
	copy(out, buf.Bytes())

	hash := sha256.Sum256(out)
	etag := `"` + hex.EncodeToString(hash[:8]) + `"`

	return RenderResult{SVG: out, ETag: etag}
}

// RenderFavicon is a shorthand for simple favicon generation (V3 compat).
func RenderFavicon(color1, color2, animation, theme string, a11y bool, reducedMotion bool) RenderResult {
	return Render(RenderOptions{
		Format:        "favicon",
		Scene:         "pure",
		Color1:        color1,
		Color2:        color2,
		Animation:     animation,
		Theme:         theme,
		Shape:         "atom",
		A11yTitle:     a11y,
		A11yDesc:      a11y,
		ReducedMotion: reducedMotion,
	})
}

// RenderPlaceholder generates a generic placeholder SVG for unknown apps.
func RenderPlaceholder() RenderResult {
	return Render(RenderOptions{
		Format:    "favicon",
		Scene:     "pure",
		Color1:    "gray",
		Animation: "static",
		Theme:     "auto",
		Shape:     "circle",
		A11yTitle: true,
		Text:      TextParams{Title: "Unknown"},
	})
}

func renderBadge(buf *bytes.Buffer, count int, dim Format, format string) {
	var cx, cy, r, fontSize float64
	switch format {
	case "favicon":
		cx, cy, r, fontSize = 26, 6, 5, 7
	case "avatar":
		cx, cy, r, fontSize = float64(dim.Width)-50, 50, 40, 36
	case "og-card":
		cx, cy, r, fontSize = float64(dim.Width)-60, 60, 45, 40
	case "hero":
		cx, cy, r, fontSize = float64(dim.Width)-80, 80, 60, 52
	default:
		cx, cy, r, fontSize = 26, 6, 5, 7
	}

	label := fmt.Sprintf("%d", count)
	if count > 99 {
		label = "99+"
		r *= 1.3
	}

	fmt.Fprintf(buf, `
  <g class="badge">
    <circle cx="%.0f" cy="%.0f" r="%.0f" fill="#ef4444"/>
    <text x="%.0f" y="%.0f" text-anchor="middle" dominant-baseline="central" fill="#fff" font-family="system-ui,sans-serif" font-size="%.0fpx" font-weight="700">%s</text>
  </g>`, cx, cy, r, cx, cy, fontSize, label)
}

func logoTransform(format, scene string, dim Format) string {
	if format == "favicon" {
		return ""
	}

	logoBase := 32.0
	var scale, tx, ty float64

	switch format {
	case "avatar":
		scale = float64(dim.Width) / 64.0
		tx = (float64(dim.Width) - logoBase*scale) / 2
		ty = (float64(dim.Height) - logoBase*scale) / 2
	case "og-card":
		if scene == "split" {
			scale = 12
			tx = float64(dim.Width) - logoBase*scale - 100
			ty = (float64(dim.Height) - logoBase*scale) / 2
		} else {
			scale = 8
			tx = (float64(dim.Width) - logoBase*scale) / 2
			ty = (float64(dim.Height)-logoBase*scale)/2 - 40
		}
	case "hero":
		scale = 15
		tx = (float64(dim.Width) - logoBase*scale) / 2
		ty = (float64(dim.Height)-logoBase*scale)/2 - 80
	}

	return fmt.Sprintf("translate(%.0f, %.0f) scale(%.0f)", tx, ty, scale)
}
