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

	case "pentagon":
		return fmt.Sprintf(`
    <g class="logos-api-core">
      <polygon points="16,2 29.5,12 24.5,28 7.5,28 2.5,12" fill="%s" stroke="%s" stroke-width="%s" stroke-linejoin="round" vector-effect="non-scaling-stroke" class="outer"/>
      <polygon points="16,7 24,13.5 21,24.5 11,24.5 8,13.5" fill="none" stroke="%s" stroke-width="%s" stroke-linejoin="round" vector-effect="non-scaling-stroke" class="orbit"/>
      <g class="core">
        <circle cx="16" cy="17" r="3.5" fill="%s"/>
        <circle cx="17.2" cy="15.8" r="1.2" fill="%s" opacity="0.9" class="dot"/>
      </g>
    </g>`,
			p.CoreBg, p.StrokeDef, p.OuterStrokeWidth,
			p.StrokeDef, p.OrbitStrokeWidth,
			p.FillDef, p.DotColor)

	case "octagon":
		return fmt.Sprintf(`
    <g class="logos-api-core">
      <polygon points="10,2 22,2 30,10 30,22 22,30 10,30 2,22 2,10" fill="%s" stroke="%s" stroke-width="%s" stroke-linejoin="round" vector-effect="non-scaling-stroke" class="outer"/>
      <polygon points="12,6 20,6 26,12 26,20 20,26 12,26 6,20 6,12" fill="none" stroke="%s" stroke-width="%s" stroke-linejoin="round" vector-effect="non-scaling-stroke" class="orbit"/>
      <g class="core">
        <circle cx="16" cy="16" r="3.5" fill="%s"/>
        <circle cx="17.2" cy="14.8" r="1.2" fill="%s" opacity="0.9" class="dot"/>
      </g>
    </g>`,
			p.CoreBg, p.StrokeDef, p.OuterStrokeWidth,
			p.StrokeDef, p.OrbitStrokeWidth,
			p.FillDef, p.DotColor)

	case "cross":
		return fmt.Sprintf(`
    <g class="logos-api-core">
      <path d="M12,2 H20 V12 H30 V20 H20 V30 H12 V20 H2 V12 H12 Z" fill="%s" stroke="%s" stroke-width="%s" stroke-linejoin="round" vector-effect="non-scaling-stroke" class="outer"/>
      <path d="M13.5,6 H18.5 V13.5 H25 V18.5 H18.5 V25 H13.5 V18.5 H7 V13.5 H13.5 Z" fill="none" stroke="%s" stroke-width="%s" stroke-linejoin="round" vector-effect="non-scaling-stroke" class="orbit"/>
      <g class="core">
        <circle cx="16" cy="16" r="3" fill="%s"/>
        <circle cx="17" cy="14.8" r="1" fill="%s" opacity="0.9" class="dot"/>
      </g>
    </g>`,
			p.CoreBg, p.StrokeDef, p.OuterStrokeWidth,
			p.StrokeDef, p.OrbitStrokeWidth,
			p.FillDef, p.DotColor)

	case "heart":
		return fmt.Sprintf(`
    <g class="logos-api-core">
      <path d="M16,28 C12,24 2,18 2,11 C2,6 6,2 10.5,2 C13,2 15,3.5 16,5.5 C17,3.5 19,2 21.5,2 C26,2 30,6 30,11 C30,18 20,24 16,28 Z" fill="%s" stroke="%s" stroke-width="%s" stroke-linejoin="round" vector-effect="non-scaling-stroke" class="outer"/>
      <path d="M16,23 C13.5,20.5 7,16.5 7,12 C7,9 9,7 11,7 C13,7 14.5,8 16,10 C17.5,8 19,7 21,7 C23,7 25,9 25,12 C25,16.5 18.5,20.5 16,23 Z" fill="none" stroke="%s" stroke-width="%s" stroke-linejoin="round" vector-effect="non-scaling-stroke" class="orbit"/>
      <g class="core">
        <circle cx="16" cy="15" r="2.5" fill="%s"/>
        <circle cx="17" cy="13.8" r="0.9" fill="%s" opacity="0.9" class="dot"/>
      </g>
    </g>`,
			p.CoreBg, p.StrokeDef, p.OuterStrokeWidth,
			p.StrokeDef, p.OrbitStrokeWidth,
			p.FillDef, p.DotColor)

	case "cloud":
		return fmt.Sprintf(`
    <g class="logos-api-core">
      <path d="M8,24 C4,24 2,21 2,18.5 C2,16 3.5,14 6,13.5 C6,9.5 9,6 13,6 C16,6 18.5,7.5 20,10 C21,9 22.5,8.5 24,8.5 C28,8.5 30,11.5 30,14.5 C30,14.5 30.5,15 30,15.5 C31,16.5 31,18 30.5,19.5 C30,21.5 28,24 24,24 Z" fill="%s" stroke="%s" stroke-width="%s" stroke-linejoin="round" vector-effect="non-scaling-stroke" class="outer"/>
      <path d="M10,21 C7.5,21 6,19.5 6,17.5 C6,16 7,14.5 9,14.5 C9,12 11,10 13.5,10 C15.5,10 17,11 18,12.5 C19,11.5 20,11 21.5,11 C24,11 25.5,13 25.5,15 C26.5,15.5 27,17 26.5,18.5 C26,20 24.5,21 22,21 Z" fill="none" stroke="%s" stroke-width="%s" stroke-linejoin="round" vector-effect="non-scaling-stroke" class="orbit"/>
      <g class="core">
        <circle cx="16" cy="16.5" r="2.5" fill="%s"/>
        <circle cx="17" cy="15.3" r="0.9" fill="%s" opacity="0.9" class="dot"/>
      </g>
    </g>`,
			p.CoreBg, p.StrokeDef, p.OuterStrokeWidth,
			p.StrokeDef, p.OrbitStrokeWidth,
			p.FillDef, p.DotColor)

	case "flame":
		return fmt.Sprintf(`
    <g class="logos-api-core">
      <path d="M16,2 C16,2 24,10 24,18 C24,24 20.5,30 16,30 C11.5,30 8,24 8,18 C8,10 16,2 16,2 Z" fill="%s" stroke="%s" stroke-width="%s" stroke-linejoin="round" vector-effect="non-scaling-stroke" class="outer"/>
      <path d="M16,10 C16,10 20,15 20,19 C20,22.5 18.5,25 16,25 C13.5,25 12,22.5 12,19 C12,15 16,10 16,10 Z" fill="none" stroke="%s" stroke-width="%s" stroke-linejoin="round" vector-effect="non-scaling-stroke" class="orbit"/>
      <g class="core">
        <circle cx="16" cy="20" r="2.5" fill="%s"/>
        <circle cx="17" cy="18.8" r="0.9" fill="%s" opacity="0.9" class="dot"/>
      </g>
    </g>`,
			p.CoreBg, p.StrokeDef, p.OuterStrokeWidth,
			p.StrokeDef, p.OrbitStrokeWidth,
			p.FillDef, p.DotColor)

	case "droplet":
		return fmt.Sprintf(`
    <g class="logos-api-core">
      <path d="M16,2 C16,2 26,14 26,20 C26,25.5 21.5,30 16,30 C10.5,30 6,25.5 6,20 C6,14 16,2 16,2 Z" fill="%s" stroke="%s" stroke-width="%s" stroke-linejoin="round" vector-effect="non-scaling-stroke" class="outer"/>
      <path d="M16,10 C16,10 22,17 22,21 C22,24.3 19.3,27 16,27 C12.7,27 10,24.3 10,21 C10,17 16,10 16,10 Z" fill="none" stroke="%s" stroke-width="%s" stroke-linejoin="round" vector-effect="non-scaling-stroke" class="orbit"/>
      <g class="core">
        <circle cx="16" cy="20" r="3" fill="%s"/>
        <circle cx="17.2" cy="18.8" r="1" fill="%s" opacity="0.9" class="dot"/>
      </g>
    </g>`,
			p.CoreBg, p.StrokeDef, p.OuterStrokeWidth,
			p.StrokeDef, p.OrbitStrokeWidth,
			p.FillDef, p.DotColor)

	case "moon":
		return fmt.Sprintf(`
    <g class="logos-api-core">
      <circle cx="16" cy="16" r="14" fill="%s" stroke="%s" stroke-width="%s" vector-effect="non-scaling-stroke" class="outer"/>
      <path d="M20,6 C15,8 12,13 12,18 C12,23 15,26 20,26 C15,28 8,24 8,16 C8,8 15,4 20,6 Z" fill="none" stroke="%s" stroke-width="%s" vector-effect="non-scaling-stroke" class="orbit"/>
      <g class="core">
        <circle cx="14" cy="16" r="3" fill="%s"/>
        <circle cx="15" cy="14.8" r="1" fill="%s" opacity="0.9" class="dot"/>
      </g>
    </g>`,
			p.CoreBg, p.StrokeDef, p.OuterStrokeWidth,
			p.StrokeDef, p.OrbitStrokeWidth,
			p.FillDef, p.DotColor)

	case "sun":
		return fmt.Sprintf(`
    <g class="logos-api-core">
      <g class="outer">
        <circle cx="16" cy="16" r="8" fill="%s" stroke="%s" stroke-width="%s" vector-effect="non-scaling-stroke"/>
        <line x1="16" y1="2" x2="16" y2="6" stroke="%s" stroke-width="%s" stroke-linecap="round" vector-effect="non-scaling-stroke"/>
        <line x1="16" y1="26" x2="16" y2="30" stroke="%s" stroke-width="%s" stroke-linecap="round" vector-effect="non-scaling-stroke"/>
        <line x1="2" y1="16" x2="6" y2="16" stroke="%s" stroke-width="%s" stroke-linecap="round" vector-effect="non-scaling-stroke"/>
        <line x1="26" y1="16" x2="30" y2="16" stroke="%s" stroke-width="%s" stroke-linecap="round" vector-effect="non-scaling-stroke"/>
        <line x1="6" y1="6" x2="9" y2="9" stroke="%s" stroke-width="%s" stroke-linecap="round" vector-effect="non-scaling-stroke"/>
        <line x1="23" y1="23" x2="26" y2="26" stroke="%s" stroke-width="%s" stroke-linecap="round" vector-effect="non-scaling-stroke"/>
        <line x1="6" y1="26" x2="9" y2="23" stroke="%s" stroke-width="%s" stroke-linecap="round" vector-effect="non-scaling-stroke"/>
        <line x1="23" y1="9" x2="26" y2="6" stroke="%s" stroke-width="%s" stroke-linecap="round" vector-effect="non-scaling-stroke"/>
      </g>
      <circle cx="16" cy="16" r="5" fill="none" stroke="%s" stroke-width="%s" vector-effect="non-scaling-stroke" class="orbit"/>
      <g class="core">
        <circle cx="16" cy="16" r="2.5" fill="%s"/>
        <circle cx="17" cy="15" r="0.9" fill="%s" opacity="0.9" class="dot"/>
      </g>
    </g>`,
			p.CoreBg, p.StrokeDef, p.OuterStrokeWidth,
			p.StrokeDef, p.OuterStrokeWidth,
			p.StrokeDef, p.OuterStrokeWidth,
			p.StrokeDef, p.OuterStrokeWidth,
			p.StrokeDef, p.OuterStrokeWidth,
			p.StrokeDef, p.OuterStrokeWidth,
			p.StrokeDef, p.OuterStrokeWidth,
			p.StrokeDef, p.OuterStrokeWidth,
			p.StrokeDef, p.OuterStrokeWidth,
			p.StrokeDef, p.OrbitStrokeWidth,
			p.FillDef, p.DotColor)

	case "arrow":
		return fmt.Sprintf(`
    <g class="logos-api-core">
      <circle cx="16" cy="16" r="14" fill="%s" stroke="%s" stroke-width="%s" vector-effect="non-scaling-stroke" class="outer"/>
      <path d="M10,16 L22,16 M17,10 L23,16 L17,22" fill="none" stroke="%s" stroke-width="%s" stroke-linecap="round" stroke-linejoin="round" vector-effect="non-scaling-stroke" class="orbit"/>
      <g class="core">
        <circle cx="23" cy="16" r="2" fill="%s"/>
      </g>
    </g>`,
			p.CoreBg, p.StrokeDef, p.OuterStrokeWidth,
			p.StrokeDef, p.OrbitStrokeWidth,
			p.FillDef)

	case "lock":
		return fmt.Sprintf(`
    <g class="logos-api-core">
      <rect x="5" y="14" width="22" height="16" rx="3" fill="%s" stroke="%s" stroke-width="%s" stroke-linejoin="round" vector-effect="non-scaling-stroke" class="outer"/>
      <path d="M10,14 V9 C10,5 12.5,2 16,2 C19.5,2 22,5 22,9 V14" fill="none" stroke="%s" stroke-width="%s" stroke-linecap="round" vector-effect="non-scaling-stroke" class="orbit"/>
      <g class="core">
        <circle cx="16" cy="22" r="3" fill="%s"/>
        <circle cx="17" cy="21" r="1" fill="%s" opacity="0.9" class="dot"/>
      </g>
    </g>`,
			p.CoreBg, p.StrokeDef, p.OuterStrokeWidth,
			p.StrokeDef, p.OrbitStrokeWidth,
			p.FillDef, p.DotColor)

	case "infinity":
		return fmt.Sprintf(`
    <g class="logos-api-core">
      <circle cx="16" cy="16" r="14" fill="%s" stroke="%s" stroke-width="%s" vector-effect="non-scaling-stroke" class="outer"/>
      <path d="M8,16 C8,12 10,10 13,10 C15,10 16,12 16,16 C16,20 17,22 19,22 C22,22 24,20 24,16 C24,12 22,10 19,10 C17,10 16,12 16,16 C16,20 15,22 13,22 C10,22 8,20 8,16 Z" fill="none" stroke="%s" stroke-width="%s" stroke-linecap="round" vector-effect="non-scaling-stroke" class="orbit"/>
      <g class="core">
        <circle cx="16" cy="16" r="2" fill="%s"/>
      </g>
    </g>`,
			p.CoreBg, p.StrokeDef, p.OuterStrokeWidth,
			p.StrokeDef, p.OrbitStrokeWidth,
			p.FillDef)

	case "crown":
		return fmt.Sprintf(`
    <g class="logos-api-core">
      <path d="M3,24 L3,12 L10,18 L16,6 L22,18 L29,12 L29,24 Z" fill="%s" stroke="%s" stroke-width="%s" stroke-linejoin="round" vector-effect="non-scaling-stroke" class="outer"/>
      <line x1="3" y1="27" x2="29" y2="27" stroke="%s" stroke-width="%s" stroke-linecap="round" vector-effect="non-scaling-stroke" class="orbit"/>
      <g class="core">
        <circle cx="16" cy="19" r="2.5" fill="%s"/>
        <circle cx="17" cy="17.8" r="0.9" fill="%s" opacity="0.9" class="dot"/>
      </g>
    </g>`,
			p.CoreBg, p.StrokeDef, p.OuterStrokeWidth,
			p.StrokeDef, p.OrbitStrokeWidth,
			p.FillDef, p.DotColor)

	case "pill":
		return fmt.Sprintf(`
    <g class="logos-api-core">
      <rect x="2" y="8" width="28" height="16" rx="8" fill="%s" stroke="%s" stroke-width="%s" vector-effect="non-scaling-stroke" class="outer"/>
      <rect x="7" y="11" width="18" height="10" rx="5" fill="none" stroke="%s" stroke-width="%s" vector-effect="non-scaling-stroke" class="orbit"/>
      <g class="core">
        <circle cx="16" cy="16" r="3" fill="%s"/>
        <circle cx="17" cy="14.8" r="1" fill="%s" opacity="0.9" class="dot"/>
      </g>
    </g>`,
			p.CoreBg, p.StrokeDef, p.OuterStrokeWidth,
			p.StrokeDef, p.OrbitStrokeWidth,
			p.FillDef, p.DotColor)

	case "target":
		return fmt.Sprintf(`
    <g class="logos-api-core">
      <circle cx="16" cy="16" r="14" fill="%s" stroke="%s" stroke-width="%s" vector-effect="non-scaling-stroke" class="outer"/>
      <circle cx="16" cy="16" r="9" fill="none" stroke="%s" stroke-width="%s" vector-effect="non-scaling-stroke" class="orbit"/>
      <circle cx="16" cy="16" r="5" fill="none" stroke="%s" stroke-width="%s" vector-effect="non-scaling-stroke"/>
      <g class="core">
        <circle cx="16" cy="16" r="2" fill="%s"/>
      </g>
    </g>`,
			p.CoreBg, p.StrokeDef, p.OuterStrokeWidth,
			p.StrokeDef, p.OrbitStrokeWidth,
			p.StrokeDef, p.OrbitStrokeWidth,
			p.FillDef)

	default:
		return RenderShape("atom", p)
	}
}

// ShapeNames returns all available shape names.
func ShapeNames() []string {
	return []string{
		"atom", "shield", "hexagon", "diamond", "bolt", "cube",
		"wave", "gear", "eye", "leaf", "star", "circle",
		"triangle", "square",
		"pentagon", "octagon", "cross", "heart", "cloud", "flame",
		"droplet", "moon", "sun", "arrow", "lock", "infinity",
		"crown", "pill", "target",
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
