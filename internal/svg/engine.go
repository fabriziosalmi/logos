package svg

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
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
	CustomPath string // AI-generated SVG path data (overrides Shape if non-empty)
	Text      TextParams

	// Parametric controls
	StrokeWidth float64
	Padding     float64
	Alpha       float64
	Badge       int
	Variant     string
	Speed       float64 // animation speed multiplier (0.25-4.0, default 1.0)
	Direction   string  // normal, reverse, alternate

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

	// AI-generated paths use 256x256 grid — override viewBox
	if opts.CustomPath != "" {
		dim = Format{Width: 256, Height: 256}
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

	// Variant overrides (applied before rendering)
	variantCss := ""
	switch opts.Variant {
	case "outline":
		// Strokes only, no fills
		theme.CoreBg = "transparent"
		if outerStrokeWidth == "0" {
			outerStrokeWidth = "1.5"
		}
	case "solid":
		// Filled background = primary color, white internals, no strokes
		theme.CoreBg = fillDef
		theme.DotColor = "#ffffff"
		outerStrokeWidth = "0"
		orbitStrokeWidth = "0"
		variantCss = "\n    .orbit{display:none}"
	case "duotone":
		// Outer filled at low opacity, core full
		variantCss = "\n    .outer{fill-opacity:0.15;stroke-opacity:0.4}.orbit{stroke-opacity:0.3}"
	case "ghost":
		// Everything faded — for inactive/disabled states
		variantCss = "\n    .logos-api-core{opacity:0.25}"
	case "ring":
		// Only the outer ring visible
		variantCss = "\n    .orbit,.core{display:none}"
	case "minimal":
		// Only the core dot, no outer/orbit
		variantCss = "\n    .outer,.orbit{display:none}"
	case "inverted":
		// Colored bg, white strokes and shapes
		theme.CoreBg = fillDef
		fillDef = "#ffffff"
		strokeDef = "#ffffff"
		theme.DotColor = fillDef
	case "badge":
		// Filled circle bg + white icon (iOS-style)
		theme.CoreBg = fillDef
		theme.DotColor = theme.CoreBg
		fillDef = "#ffffff"
		strokeDef = "#ffffff"
		outerStrokeWidth = "0"
	case "glow":
		// Normal + outer glow effect
		variantCss = fmt.Sprintf("\n    .outer{filter:drop-shadow(0 0 4px #%s) drop-shadow(0 0 12px #%s)}", primaryHex, primaryHex)

	case "dotted":
		// Dashed/dotted strokes — blueprint/technical feel
		variantCss = "\n    .outer{stroke-dasharray:3 3}.orbit{stroke-dasharray:4 4}"

	case "sketch":
		// Hand-drawn rough strokes via irregular dash + slight rotation jitter
		variantCss = "\n    .outer{stroke-dasharray:8 2 2 2;stroke-linecap:round;stroke-linejoin:round}" +
			"\n    .orbit{stroke-dasharray:6 3 1 3;stroke-linecap:round}"

	case "neon-outline":
		// Outline + glow combined — cyberpunk neon sign
		theme.CoreBg = "transparent"
		if outerStrokeWidth == "0" {
			outerStrokeWidth = "1.5"
		}
		variantCss = fmt.Sprintf("\n    .outer{filter:drop-shadow(0 0 3px #%s) drop-shadow(0 0 8px #%s)}", primaryHex, primaryHex) +
			fmt.Sprintf("\n    .orbit{filter:drop-shadow(0 0 2px #%s)}", primaryHex) +
			fmt.Sprintf("\n    .core{filter:drop-shadow(0 0 4px #%s)}", primaryHex)

	case "retro":
		// 80s offset shadow — colored shadow offset bottom-right
		variantCss = fmt.Sprintf("\n    .logos-api-core{filter:drop-shadow(2px 2px 0 #%s)}", primaryHex)

	case "sticker":
		// White border cutout — like a die-cut sticker
		outerStrokeWidth = "3"
		variantCss = "\n    .outer{stroke:#ffffff;paint-order:stroke fill}"

	case "emboss":
		// Raised 3D look via lighting filter
		variantCss = "\n    .logos-api-core{filter:drop-shadow(1px 1px 0 rgba(255,255,255,0.3)) drop-shadow(-1px -1px 0 rgba(0,0,0,0.4))}"

	case "xray":
		// High contrast monochrome — inverted luminance
		variantCss = "\n    .logos-api-core{filter:contrast(2) brightness(1.5)}.outer{fill:transparent}"
		theme.CoreBg = "transparent"

	case "double":
		// Double concentric stroke on outer
		variantCss = "\n    .outer{stroke-dasharray:none;paint-order:stroke fill;stroke-width:4;stroke-opacity:0.3}" +
			fmt.Sprintf("\n    .outer+circle{display:none}")
		outerStrokeWidth = "1.5"

	case "half":
		// Half-filled — clip left half, like a loading/progress indicator
		variantCss = "\n    .logos-api-core{clip-path:inset(0 50% 0 0)}"

	case "stamp":
		// Thick bold strokes, rounded — rubber stamp aesthetic
		outerStrokeWidth = "3"
		orbitStrokeWidth = "2.5"
		theme.CoreBg = "transparent"
		variantCss = "\n    .outer,.orbit{stroke-linecap:round;stroke-linejoin:round}"

	// === NEW VARIANTS (30 more) ===

	case "flat":
		// No stroke, solid fill, no inner detail
		outerStrokeWidth = "0"
		orbitStrokeWidth = "0"
		variantCss = "\n    .orbit,.dot{display:none}"
	case "thin":
		// Ultra-thin strokes (hairline)
		outerStrokeWidth = "0.5"
		orbitStrokeWidth = "0.5"
		theme.CoreBg = "transparent"
	case "thick":
		// Extra-thick strokes
		outerStrokeWidth = "4"
		orbitStrokeWidth = "3"
		theme.CoreBg = "transparent"
	case "shadow-lift":
		// Floating card shadow below
		variantCss = "\n    .logos-api-core{filter:drop-shadow(0 4px 8px rgba(0,0,0,0.3))}"
	case "long-shadow":
		// Material design long shadow
		variantCss = fmt.Sprintf("\n    .logos-api-core{filter:drop-shadow(3px 3px 0 rgba(0,0,0,0.15)) drop-shadow(6px 6px 0 rgba(0,0,0,0.08)) drop-shadow(9px 9px 0 rgba(0,0,0,0.03))}")
	case "cutout":
		// Outer filled, inner shapes cut out (negative space)
		theme.CoreBg = fillDef
		fillDef = theme.DotColor
		strokeDef = theme.DotColor
	case "duotone-dark":
		// Dark duotone — outer at 30% opacity
		variantCss = "\n    .outer{fill-opacity:0.3;stroke-opacity:0.5}.orbit{stroke-opacity:0.2}.core{opacity:0.9}"
	case "pastel":
		// Soft pastel look — reduced saturation via filter
		variantCss = "\n    .logos-api-core{filter:saturate(0.5) brightness(1.3)}"
	case "monochrome":
		// Grayscale
		variantCss = "\n    .logos-api-core{filter:grayscale(1)}"
	case "sepia":
		// Vintage sepia tone
		variantCss = "\n    .logos-api-core{filter:sepia(0.8) saturate(1.2)}"
	case "invert-colors":
		// CSS invert filter
		variantCss = "\n    .logos-api-core{filter:invert(1)}"
	case "blur":
		// Soft blur (dreamy/frosted)
		variantCss = "\n    .logos-api-core{filter:blur(1px)}"
	case "pixel":
		// Pixelated via CSS (works in modern browsers)
		variantCss = "\n    .logos-api-core{image-rendering:pixelated;filter:contrast(1.5)}"
	case "mirror":
		// Horizontal mirror flip
		variantCss = "\n    .logos-api-core{transform:scaleX(-1);transform-origin:16px 16px}"
	case "flip-v":
		// Vertical flip
		variantCss = "\n    .logos-api-core{transform:scaleY(-1);transform-origin:16px 16px}"
	case "rotate-45":
		// 45 degree rotation
		variantCss = "\n    .logos-api-core{transform:rotate(45deg);transform-origin:16px 16px}"
	case "rotate-90":
		variantCss = "\n    .logos-api-core{transform:rotate(90deg);transform-origin:16px 16px}"
	case "rotate-180":
		variantCss = "\n    .logos-api-core{transform:rotate(180deg);transform-origin:16px 16px}"
	case "scale-sm":
		// Scaled down (with breathing room)
		variantCss = "\n    .logos-api-core{transform:scale(0.7);transform-origin:16px 16px}"
	case "scale-lg":
		// Scaled up (fills more)
		variantCss = "\n    .logos-api-core{transform:scale(1.2);transform-origin:16px 16px}"
	case "shadow-colored":
		// Colored shadow matching the primary color
		variantCss = fmt.Sprintf("\n    .logos-api-core{filter:drop-shadow(0 3px 6px #%s88)}", primaryHex)
	case "hue-rotate":
		// Shift hue by 180 degrees (complementary color)
		variantCss = "\n    .logos-api-core{filter:hue-rotate(180deg)}"
	case "warm":
		// Warm color shift
		variantCss = "\n    .logos-api-core{filter:hue-rotate(-20deg) saturate(1.3) brightness(1.05)}"
	case "cool":
		// Cool color shift
		variantCss = "\n    .logos-api-core{filter:hue-rotate(20deg) saturate(0.9) brightness(1.05)}"
	case "disco":
		// Animated hue rotation
		variantCss = "\n    .logos-api-core{animation:disco 3s linear infinite}" +
			"\n    @keyframes disco{0%{filter:hue-rotate(0deg)}100%{filter:hue-rotate(360deg)}}"
	case "glitch-shift":
		// Animated glitch (translateX jitter)
		variantCss = "\n    .logos-api-core{animation:glitchshift 0.5s steps(3) infinite}" +
			"\n    @keyframes glitchshift{0%,100%{transform:translateX(0)}25%{transform:translateX(-2px)}75%{transform:translateX(2px)}}"
	case "scan-line":
		// CRT scan line effect via repeating gradient
		variantCss = "\n    .logos-api-core{background:repeating-linear-gradient(transparent,transparent 2px,rgba(0,0,0,0.1) 2px,rgba(0,0,0,0.1) 4px)}"
	case "neon-text":
		// Double glow — inner white + outer colored
		theme.CoreBg = "transparent"
		outerStrokeWidth = "2"
		variantCss = fmt.Sprintf("\n    .outer,.orbit{filter:drop-shadow(0 0 2px #fff) drop-shadow(0 0 6px #%s) drop-shadow(0 0 12px #%s)}", primaryHex, primaryHex)
	case "glass-morph":
		// Glassmorphism — frosted background with translucency
		theme.CoreBg = "rgba(255,255,255,0.1)"
		variantCss = "\n    .outer{-webkit-backdrop-filter:blur(8px);backdrop-filter:blur(8px);fill-opacity:0.3}"
	case "wireframe-3d":
		// 3D wireframe — perspective transform + thin strokes
		outerStrokeWidth = "0.8"
		orbitStrokeWidth = "0.6"
		theme.CoreBg = "transparent"
		variantCss = "\n    .logos-api-core{transform:perspective(100px) rotateY(15deg) rotateX(-5deg);transform-origin:16px 16px}"
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

		// Speed multiplier: scales all animation durations
		speed := opts.Speed
		if speed <= 0 || speed > 10 {
			speed = 1.0
		}
		if speed != 1.0 {
			fmt.Fprintf(buf, "\n    .outer,.orbit,.core{animation-duration:calc(var(--d, 1s) * %.2f)}", 1.0/speed)
		}

		// Direction override
		dir := opts.Direction
		if dir == "reverse" || dir == "alternate" || dir == "alternate-reverse" {
			fmt.Fprintf(buf, "\n    .outer,.orbit,.core{animation-direction:%s}", dir)
		}

		if opts.ReducedMotion {
			buf.WriteString("\n    @media(prefers-reduced-motion:reduce){.outer,.orbit,.core{animation:none!important}}")
		}
	}

	buf.WriteString("\n    @media(forced-colors:active){.outer,.orbit{stroke:CanvasText}.core circle,.core polygon{fill:CanvasText}.dot{fill:Canvas}}")

	if opts.Hover {
		buf.WriteString("\n    .logos-api-core:hover{filter:brightness(1.2);transition:filter .2s ease}")
	}

	if variantCss != "" {
		buf.WriteString(variantCss)
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

	// THE SHAPE — predefined, icon pack, or procedural
	shapeParams := ShapeParams{
		CoreBg:           theme.CoreBg,
		FillDef:          fillDef,
		StrokeDef:        strokeDef,
		DotColor:         theme.DotColor,
		OuterStrokeWidth: outerStrokeWidth,
		OrbitStrokeWidth: orbitStrokeWidth,
	}

	if opts.CustomPath != "" {
		// Procedural-generated path (256x256 grid)
		fmt.Fprintf(buf, `
    <g class="logos-api-core">
      <path d="%s" fill="%s" stroke="%s" stroke-width="%s" vector-effect="non-scaling-stroke" class="outer"/>
    </g>`, opts.CustomPath, fillDef, strokeDef, outerStrokeWidth)
	} else if strings.Contains(shape, ":") {
		// Icon pack shape (e.g. "lucide:rocket", "hero:fire")
		buf.WriteString(RenderIconShape(shape, shapeParams))
	} else {
		// Built-in shape
		buf.WriteString(RenderShape(shape, shapeParams))
	}

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
