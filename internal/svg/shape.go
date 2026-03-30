package svg

import (
	"fmt"
)

// ShapeParams holds resolved visual params for shape rendering.
type ShapeParams struct {
	CoreBg           string
	FillDef          string
	StrokeDef        string
	DotColor         string
	OuterStrokeWidth string
	OrbitStrokeWidth string
}

// RenderShape returns the SVG elements for a given shape name.
// All shapes use .outer, .orbit, .core classes for animation compatibility.
// Everything is designed on a 32x32 grid with transform-origin: 16px 16px.
func RenderShape(shape string, p ShapeParams) string {
	switch shape {

	case "atom":
		// The original: circle + orbital ellipse + core dot (RESERVED)
		return fmt.Sprintf(`
    <g class="logos-api-core">
      <circle cx="16" cy="16" r="14" fill="%s" stroke="%s" stroke-width="%s" vector-effect="non-scaling-stroke" class="outer"/>
      <ellipse cx="16" cy="16" rx="10" ry="7" fill="none" stroke="%s" stroke-width="%s" vector-effect="non-scaling-stroke" class="orbit"/>
      <g class="core">
        <circle cx="16" cy="16" r="3.5" fill="%s"/>
        <circle cx="17.2" cy="14.8" r="1.2" fill="%s" opacity="0.9" class="dot"/>
      </g>
    </g>`,
			p.CoreBg, p.StrokeDef, p.OuterStrokeWidth,
			p.StrokeDef, p.OrbitStrokeWidth,
			p.FillDef, p.DotColor)

	case "shield":
		// Security shield: rounded shield outline + inner shield + core dot
		return fmt.Sprintf(`
    <g class="logos-api-core">
      <path d="M16 2L4 8v8c0 6.6 5.1 12.8 12 14 6.9-1.2 12-7.4 12-14V8L16 2z" fill="%s" stroke="%s" stroke-width="%s" stroke-linejoin="round" vector-effect="non-scaling-stroke" class="outer"/>
      <path d="M16 6L8 10v5c0 4.4 3.4 8.5 8 9.3 4.6-.8 8-4.9 8-9.3v-5L16 6z" fill="none" stroke="%s" stroke-width="%s" stroke-linejoin="round" vector-effect="non-scaling-stroke" class="orbit"/>
      <g class="core">
        <circle cx="16" cy="15" r="3.5" fill="%s"/>
        <circle cx="17.2" cy="13.8" r="1.2" fill="%s" opacity="0.9" class="dot"/>
      </g>
    </g>`,
			p.CoreBg, p.StrokeDef, p.OuterStrokeWidth,
			p.StrokeDef, p.OrbitStrokeWidth,
			p.FillDef, p.DotColor)

	case "hexagon":
		// Tech hexagon: hex border + inner rotated hex + core
		return fmt.Sprintf(`
    <g class="logos-api-core">
      <polygon points="16,2 28.1,9 28.1,23 16,30 3.9,23 3.9,9" fill="%s" stroke="%s" stroke-width="%s" stroke-linejoin="round" vector-effect="non-scaling-stroke" class="outer"/>
      <polygon points="16,7 23.8,11.5 23.8,20.5 16,25 8.2,20.5 8.2,11.5" fill="none" stroke="%s" stroke-width="%s" stroke-linejoin="round" vector-effect="non-scaling-stroke" class="orbit"/>
      <g class="core">
        <circle cx="16" cy="16" r="3.5" fill="%s"/>
        <circle cx="17.2" cy="14.8" r="1.2" fill="%s" opacity="0.9" class="dot"/>
      </g>
    </g>`,
			p.CoreBg, p.StrokeDef, p.OuterStrokeWidth,
			p.StrokeDef, p.OrbitStrokeWidth,
			p.FillDef, p.DotColor)

	case "diamond":
		// Premium diamond: rotated square + inner diamond + core
		return fmt.Sprintf(`
    <g class="logos-api-core">
      <rect x="5.3" y="5.3" width="21.4" height="21.4" rx="2" transform="rotate(45 16 16)" fill="%s" stroke="%s" stroke-width="%s" vector-effect="non-scaling-stroke" class="outer"/>
      <rect x="8.8" y="8.8" width="14.4" height="14.4" rx="1" transform="rotate(45 16 16)" fill="none" stroke="%s" stroke-width="%s" vector-effect="non-scaling-stroke" class="orbit"/>
      <g class="core">
        <circle cx="16" cy="16" r="3" fill="%s"/>
        <circle cx="17" cy="14.8" r="1" fill="%s" opacity="0.9" class="dot"/>
      </g>
    </g>`,
			p.CoreBg, p.StrokeDef, p.OuterStrokeWidth,
			p.StrokeDef, p.OrbitStrokeWidth,
			p.FillDef, p.DotColor)

	case "bolt":
		// Energy bolt: lightning + circle halo + spark core
		return fmt.Sprintf(`
    <g class="logos-api-core">
      <circle cx="16" cy="16" r="14" fill="%s" stroke="%s" stroke-width="%s" vector-effect="non-scaling-stroke" class="outer"/>
      <polygon points="18,3 10,17 15,17 14,29 22,15 17,15" fill="none" stroke="%s" stroke-width="%s" stroke-linejoin="round" vector-effect="non-scaling-stroke" class="orbit"/>
      <g class="core">
        <polygon points="18,9 13,17 16,17 14,23 19,15 16,15" fill="%s"/>
      </g>
    </g>`,
			p.CoreBg, p.StrokeDef, p.OuterStrokeWidth,
			p.StrokeDef, p.OrbitStrokeWidth,
			p.FillDef)

	case "cube":
		// Infra cube: isometric box outline + inner structure + core
		return fmt.Sprintf(`
    <g class="logos-api-core">
      <polygon points="16,2 30,9.5 30,22.5 16,30 2,22.5 2,9.5" fill="%s" stroke="%s" stroke-width="%s" stroke-linejoin="round" vector-effect="non-scaling-stroke" class="outer"/>
      <line x1="16" y1="16" x2="16" y2="30" stroke="%s" stroke-width="%s" vector-effect="non-scaling-stroke" class="orbit"/>
      <line x1="16" y1="16" x2="2" y2="9.5" stroke="%s" stroke-width="%s" vector-effect="non-scaling-stroke"/>
      <line x1="16" y1="16" x2="30" y2="9.5" stroke="%s" stroke-width="%s" vector-effect="non-scaling-stroke"/>
      <g class="core">
        <circle cx="16" cy="16" r="2.5" fill="%s"/>
        <circle cx="17" cy="15" r="0.9" fill="%s" opacity="0.9" class="dot"/>
      </g>
    </g>`,
			p.CoreBg, p.StrokeDef, p.OuterStrokeWidth,
			p.StrokeDef, p.OrbitStrokeWidth,
			p.StrokeDef, p.OrbitStrokeWidth,
			p.StrokeDef, p.OrbitStrokeWidth,
			p.FillDef, p.DotColor)

	case "wave":
		// Data wave: sine wave lines + floating core
		return fmt.Sprintf(`
    <g class="logos-api-core">
      <circle cx="16" cy="16" r="14" fill="%s" stroke="%s" stroke-width="%s" vector-effect="non-scaling-stroke" class="outer"/>
      <path d="M4,16 C7,10 11,10 16,16 C21,22 25,22 28,16" fill="none" stroke="%s" stroke-width="%s" stroke-linecap="round" vector-effect="non-scaling-stroke" class="orbit"/>
      <path d="M4,16 C7,22 11,22 16,16 C21,10 25,10 28,16" fill="none" stroke="%s" stroke-width="0.8" stroke-linecap="round" vector-effect="non-scaling-stroke" opacity="0.4"/>
      <g class="core">
        <circle cx="16" cy="16" r="3" fill="%s"/>
        <circle cx="17" cy="14.8" r="1" fill="%s" opacity="0.9" class="dot"/>
      </g>
    </g>`,
			p.CoreBg, p.StrokeDef, p.OuterStrokeWidth,
			p.StrokeDef, p.OrbitStrokeWidth,
			p.StrokeDef,
			p.FillDef, p.DotColor)

	case "gear":
		// Ops gear: gear outline + inner ring + core
		return fmt.Sprintf(`
    <g class="logos-api-core">
      <path d="M16,2 L18,5.5 L22,4.5 L22.5,8.5 L26.5,9.5 L25,13 L28.5,15.5 L26,18.5 L28,22 L24.5,23 L24,27 L20.5,26 L18,29.5 L16,26.5 L14,29.5 L11.5,26 L8,27 L7.5,23 L4,22 L6,18.5 L3.5,15.5 L7,13 L5.5,9.5 L9.5,8.5 L10,4.5 L14,5.5 Z" fill="%s" stroke="%s" stroke-width="%s" stroke-linejoin="round" vector-effect="non-scaling-stroke" class="outer"/>
      <circle cx="16" cy="16" r="7" fill="none" stroke="%s" stroke-width="%s" vector-effect="non-scaling-stroke" class="orbit"/>
      <g class="core">
        <circle cx="16" cy="16" r="3.5" fill="%s"/>
        <circle cx="17.2" cy="14.8" r="1.2" fill="%s" opacity="0.9" class="dot"/>
      </g>
    </g>`,
			p.CoreBg, p.StrokeDef, p.OuterStrokeWidth,
			p.StrokeDef, p.OrbitStrokeWidth,
			p.FillDef, p.DotColor)

	case "eye":
		// Monitoring eye: eye shape + iris ring + pupil core
		return fmt.Sprintf(`
    <g class="logos-api-core">
      <path d="M2,16 C6,6 26,6 30,16 C26,26 6,26 2,16 Z" fill="%s" stroke="%s" stroke-width="%s" stroke-linejoin="round" vector-effect="non-scaling-stroke" class="outer"/>
      <circle cx="16" cy="16" r="6" fill="none" stroke="%s" stroke-width="%s" vector-effect="non-scaling-stroke" class="orbit"/>
      <g class="core">
        <circle cx="16" cy="16" r="3.5" fill="%s"/>
        <circle cx="17.5" cy="14.8" r="1.3" fill="%s" opacity="0.9" class="dot"/>
      </g>
    </g>`,
			p.CoreBg, p.StrokeDef, p.OuterStrokeWidth,
			p.StrokeDef, p.OrbitStrokeWidth,
			p.FillDef, p.DotColor)

	case "leaf":
		// Organic leaf: leaf outline + center vein + dewdrop core
		return fmt.Sprintf(`
    <g class="logos-api-core">
      <path d="M16,2 C26,6 28,20 16,30 C4,20 6,6 16,2 Z" fill="%s" stroke="%s" stroke-width="%s" stroke-linejoin="round" vector-effect="non-scaling-stroke" class="outer"/>
      <path d="M16,6 C16,6 16,26 16,26 M10,14 Q16,10 22,14 M11,20 Q16,16 21,20" fill="none" stroke="%s" stroke-width="%s" stroke-linecap="round" vector-effect="non-scaling-stroke" class="orbit"/>
      <g class="core">
        <circle cx="16" cy="13" r="2.5" fill="%s"/>
        <circle cx="17" cy="12" r="0.9" fill="%s" opacity="0.9" class="dot"/>
      </g>
    </g>`,
			p.CoreBg, p.StrokeDef, p.OuterStrokeWidth,
			p.StrokeDef, p.OrbitStrokeWidth,
			p.FillDef, p.DotColor)

	case "star":
		// Star: 5-point star + inner star + core
		return fmt.Sprintf(`
    <g class="logos-api-core">
      <polygon points="16,1 20.5,11.5 31,12.5 23,20 25.5,30.5 16,25 6.5,30.5 9,20 1,12.5 11.5,11.5" fill="%s" stroke="%s" stroke-width="%s" stroke-linejoin="round" vector-effect="non-scaling-stroke" class="outer"/>
      <polygon points="16,7 18.8,13.2 25.5,14 20.5,18.5 22,25 16,21.5 10,25 11.5,18.5 6.5,14 13.2,13.2" fill="none" stroke="%s" stroke-width="%s" stroke-linejoin="round" vector-effect="non-scaling-stroke" class="orbit"/>
      <g class="core">
        <circle cx="16" cy="16" r="3" fill="%s"/>
        <circle cx="17" cy="14.8" r="1" fill="%s" opacity="0.9" class="dot"/>
      </g>
    </g>`,
			p.CoreBg, p.StrokeDef, p.OuterStrokeWidth,
			p.StrokeDef, p.OrbitStrokeWidth,
			p.FillDef, p.DotColor)

	case "circle":
		// Minimal: double circle + core (simple, universal)
		return fmt.Sprintf(`
    <g class="logos-api-core">
      <circle cx="16" cy="16" r="14" fill="%s" stroke="%s" stroke-width="%s" vector-effect="non-scaling-stroke" class="outer"/>
      <circle cx="16" cy="16" r="9" fill="none" stroke="%s" stroke-width="%s" vector-effect="non-scaling-stroke" class="orbit"/>
      <g class="core">
        <circle cx="16" cy="16" r="4" fill="%s"/>
        <circle cx="17.5" cy="14.5" r="1.3" fill="%s" opacity="0.9" class="dot"/>
      </g>
    </g>`,
			p.CoreBg, p.StrokeDef, p.OuterStrokeWidth,
			p.StrokeDef, p.OrbitStrokeWidth,
			p.FillDef, p.DotColor)

	case "triangle":
		// Minimal triangle: equilateral + inner triangle + core
		return fmt.Sprintf(`
    <g class="logos-api-core">
      <polygon points="16,2 30,28 2,28" fill="%s" stroke="%s" stroke-width="%s" stroke-linejoin="round" vector-effect="non-scaling-stroke" class="outer"/>
      <polygon points="16,9 24,24 8,24" fill="none" stroke="%s" stroke-width="%s" stroke-linejoin="round" vector-effect="non-scaling-stroke" class="orbit"/>
      <g class="core">
        <circle cx="16" cy="20" r="3" fill="%s"/>
        <circle cx="17" cy="18.8" r="1" fill="%s" opacity="0.9" class="dot"/>
      </g>
    </g>`,
			p.CoreBg, p.StrokeDef, p.OuterStrokeWidth,
			p.StrokeDef, p.OrbitStrokeWidth,
			p.FillDef, p.DotColor)

	case "square":
		// Rounded square: rounded rect + inner rect + core
		return fmt.Sprintf(`
    <g class="logos-api-core">
      <rect x="2" y="2" width="28" height="28" rx="5" fill="%s" stroke="%s" stroke-width="%s" vector-effect="non-scaling-stroke" class="outer"/>
      <rect x="7" y="7" width="18" height="18" rx="3" fill="none" stroke="%s" stroke-width="%s" vector-effect="non-scaling-stroke" class="orbit"/>
      <g class="core">
        <circle cx="16" cy="16" r="3.5" fill="%s"/>
        <circle cx="17.2" cy="14.8" r="1.2" fill="%s" opacity="0.9" class="dot"/>
      </g>
    </g>`,
			p.CoreBg, p.StrokeDef, p.OuterStrokeWidth,
			p.StrokeDef, p.OrbitStrokeWidth,
			p.FillDef, p.DotColor)

	default:
		// Fallback to atom
		return RenderShape("atom", p)
	}
}

// ShapeNames returns all available shape names.
func ShapeNames() []string {
	return []string{
		"atom", "shield", "hexagon", "diamond", "bolt", "cube",
		"wave", "gear", "eye", "leaf", "star", "circle",
		"triangle", "square",
	}
}

// ValidShape checks if a shape name exists.
func ValidShape(name string) bool {
	for _, s := range ShapeNames() {
		if s == name {
			return true
		}
	}
	return name == ""
}
