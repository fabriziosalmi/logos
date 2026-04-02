package svg

import "fmt"

// SceneDefs returns SVG <defs> content for a given scene.
func SceneDefs(scene, primaryHex, secondaryHex string) string {
	switch scene {
	case "spotlight":
		return fmt.Sprintf(`
      <radialGradient id="spotlightGrad" cx="50%%" cy="50%%" r="50%%" fx="50%%" fy="50%%">
        <stop offset="0%%" stop-color="#%s" stop-opacity="0.3" />
        <stop offset="100%%" stop-color="#%s" stop-opacity="0" />
      </radialGradient>`, primaryHex, secondaryHex)
	case "grid":
		return fmt.Sprintf(`
      <pattern id="gridPattern" width="40" height="40" patternUnits="userSpaceOnUse">
        <path d="M 40 0 L 0 0 0 40" fill="none" stroke="rgba(255,255,255,0.05)" stroke-width="1"/>
      </pattern>
      <radialGradient id="gridGlow" cx="50%%" cy="100%%" r="70%%">
        <stop offset="0%%" stop-color="#%s" stop-opacity="0.15" />
        <stop offset="100%%" stop-color="#000000" stop-opacity="0" />
      </radialGradient>`, primaryHex)
	case "dots":
		return fmt.Sprintf(`
      <pattern id="dotsPattern" width="20" height="20" patternUnits="userSpaceOnUse">
        <circle cx="10" cy="10" r="1" fill="#%s" opacity="0.1"/>
      </pattern>`, primaryHex)
	case "diagonal":
		return fmt.Sprintf(`
      <pattern id="diagPattern" width="10" height="10" patternUnits="userSpaceOnUse" patternTransform="rotate(45)">
        <line x1="0" y1="0" x2="0" y2="10" stroke="#%s" stroke-width="0.5" opacity="0.08"/>
      </pattern>`, primaryHex)
	case "noise-bg":
		return `
      <filter id="bgNoise" x="0" y="0" width="100%" height="100%">
        <feTurbulence type="fractalNoise" baseFrequency="0.8" numOctaves="4" stitchTiles="stitch"/>
        <feColorMatrix type="saturate" values="0"/>
        <feComponentTransfer><feFuncA type="linear" slope="0.03"/></feComponentTransfer>
      </filter>`
	case "gradient":
		return fmt.Sprintf(`
      <linearGradient id="bgGrad" x1="0%%" y1="0%%" x2="100%%" y2="100%%">
        <stop offset="0%%" stop-color="#%s" stop-opacity="0.15"/>
        <stop offset="100%%" stop-color="#%s" stop-opacity="0.05"/>
      </linearGradient>`, primaryHex, secondaryHex)
	case "radial":
		return fmt.Sprintf(`
      <radialGradient id="radialBg" cx="50%%" cy="50%%" r="60%%">
        <stop offset="0%%" stop-color="#%s" stop-opacity="0.2"/>
        <stop offset="100%%" stop-color="#000" stop-opacity="0"/>
      </radialGradient>`, primaryHex)
	case "vignette":
		return `
      <radialGradient id="vignetteBg" cx="50%" cy="50%" r="70%">
        <stop offset="0%" stop-color="#000" stop-opacity="0"/>
        <stop offset="100%" stop-color="#000" stop-opacity="0.6"/>
      </radialGradient>`
	case "hexgrid":
		return fmt.Sprintf(`
      <pattern id="hexPattern" width="28" height="49" patternUnits="userSpaceOnUse">
        <path d="M14,0 L28,8.5 L28,24.5 L14,33 L0,24.5 L0,8.5 Z M14,16.5 L28,25 L28,41 L14,49.5 L0,41 L0,25 Z" fill="none" stroke="#%s" stroke-width="0.5" opacity="0.06"/>
      </pattern>`, primaryHex)
	case "circuit":
		return fmt.Sprintf(`
      <pattern id="circuitPattern" width="50" height="50" patternUnits="userSpaceOnUse">
        <path d="M25,0 V15 M25,35 V50 M0,25 H15 M35,25 H50" fill="none" stroke="#%s" stroke-width="0.5" opacity="0.06"/>
        <circle cx="25" cy="25" r="2" fill="none" stroke="#%s" stroke-width="0.5" opacity="0.08"/>
      </pattern>`, primaryHex, primaryHex)
	// === NEW SCENES (18 more) ===

	case "crosshatch":
		return fmt.Sprintf(`
      <pattern id="crosshatchP" width="8" height="8" patternUnits="userSpaceOnUse" patternTransform="rotate(45)">
        <line x1="0" y1="0" x2="0" y2="8" stroke="#%s" stroke-width="0.3" opacity="0.06"/>
      </pattern>
      <pattern id="crosshatchP2" width="8" height="8" patternUnits="userSpaceOnUse" patternTransform="rotate(-45)">
        <line x1="0" y1="0" x2="0" y2="8" stroke="#%s" stroke-width="0.3" opacity="0.06"/>
      </pattern>`, primaryHex, primaryHex)
	case "stars":
		return fmt.Sprintf(`
      <pattern id="starsP" width="50" height="50" patternUnits="userSpaceOnUse">
        <circle cx="5" cy="5" r="0.5" fill="#%s" opacity="0.15"/>
        <circle cx="25" cy="30" r="0.3" fill="#%s" opacity="0.1"/>
        <circle cx="40" cy="10" r="0.4" fill="#%s" opacity="0.12"/>
        <circle cx="15" cy="40" r="0.3" fill="#%s" opacity="0.08"/>
        <circle cx="45" cy="45" r="0.5" fill="#%s" opacity="0.1"/>
      </pattern>`, primaryHex, primaryHex, primaryHex, primaryHex, primaryHex)
	case "waves-bg":
		return fmt.Sprintf(`
      <pattern id="wavesP" width="60" height="20" patternUnits="userSpaceOnUse">
        <path d="M0,10 C15,0 15,20 30,10 C45,0 45,20 60,10" fill="none" stroke="#%s" stroke-width="0.5" opacity="0.06"/>
      </pattern>`, primaryHex)
	case "diamond-grid":
		return fmt.Sprintf(`
      <pattern id="diamondP" width="20" height="20" patternUnits="userSpaceOnUse">
        <path d="M10,0 L20,10 L10,20 L0,10 Z" fill="none" stroke="#%s" stroke-width="0.4" opacity="0.06"/>
      </pattern>`, primaryHex)
	case "brick":
		return fmt.Sprintf(`
      <pattern id="brickP" width="30" height="15" patternUnits="userSpaceOnUse">
        <rect width="30" height="15" fill="none" stroke="#%s" stroke-width="0.4" opacity="0.05"/>
        <line x1="15" y1="0" x2="15" y2="7.5" stroke="#%s" stroke-width="0.4" opacity="0.05"/>
      </pattern>`, primaryHex, primaryHex)
	case "triangle-grid":
		return fmt.Sprintf(`
      <pattern id="triP" width="20" height="17.3" patternUnits="userSpaceOnUse">
        <path d="M0,17.3 L10,0 L20,17.3 Z M10,17.3 L20,0 L30,17.3 Z M-10,17.3 L0,0 L10,17.3 Z" fill="none" stroke="#%s" stroke-width="0.3" opacity="0.05"/>
      </pattern>`, primaryHex)
	case "concentric":
		return fmt.Sprintf(`
      <radialGradient id="concentricBg" cx="50%%" cy="50%%" r="50%%">
        <stop offset="0%%" stop-color="#%s" stop-opacity="0.08"/>
        <stop offset="25%%" stop-color="#000" stop-opacity="0"/>
        <stop offset="50%%" stop-color="#%s" stop-opacity="0.05"/>
        <stop offset="75%%" stop-color="#000" stop-opacity="0"/>
        <stop offset="100%%" stop-color="#%s" stop-opacity="0.03"/>
      </radialGradient>`, primaryHex, primaryHex, primaryHex)
	case "scanlines":
		return `
      <pattern id="scanlinesP" width="4" height="4" patternUnits="userSpaceOnUse">
        <line x1="0" y1="0" x2="4" y2="0" stroke="rgba(255,255,255,0.03)" stroke-width="1"/>
      </pattern>`
	case "topography":
		return fmt.Sprintf(`
      <pattern id="topoP" width="80" height="80" patternUnits="userSpaceOnUse">
        <circle cx="40" cy="40" r="30" fill="none" stroke="#%s" stroke-width="0.3" opacity="0.04"/>
        <circle cx="40" cy="40" r="20" fill="none" stroke="#%s" stroke-width="0.3" opacity="0.04"/>
        <circle cx="40" cy="40" r="10" fill="none" stroke="#%s" stroke-width="0.3" opacity="0.04"/>
        <circle cx="10" cy="10" r="15" fill="none" stroke="#%s" stroke-width="0.3" opacity="0.03"/>
      </pattern>`, primaryHex, primaryHex, primaryHex, primaryHex)
	case "plus-grid":
		return fmt.Sprintf(`
      <pattern id="plusP" width="24" height="24" patternUnits="userSpaceOnUse">
        <path d="M12,8 V16 M8,12 H16" fill="none" stroke="#%s" stroke-width="0.5" opacity="0.06"/>
      </pattern>`, primaryHex)
	case "isometric":
		return fmt.Sprintf(`
      <pattern id="isoP" width="28" height="48" patternUnits="userSpaceOnUse">
        <path d="M14,0 L28,8 L14,16 L0,8 Z M14,16 L28,24 L14,32 L0,24 Z M14,32 L28,40 L14,48 L0,40 Z" fill="none" stroke="#%s" stroke-width="0.3" opacity="0.05"/>
      </pattern>`, primaryHex)
	case "aurora-bg":
		return fmt.Sprintf(`
      <linearGradient id="auroraBg1" x1="0%%" y1="0%%" x2="100%%" y2="100%%">
        <stop offset="0%%" stop-color="#%s" stop-opacity="0.1"/>
        <stop offset="50%%" stop-color="#%s" stop-opacity="0.05"/>
        <stop offset="100%%" stop-color="#%s" stop-opacity="0.15"/>
      </linearGradient>`, primaryHex, secondaryHex, primaryHex)
	case "spotlight-color":
		return fmt.Sprintf(`
      <radialGradient id="spotColorBg" cx="50%%" cy="50%%" r="40%%">
        <stop offset="0%%" stop-color="#%s" stop-opacity="0.2"/>
        <stop offset="100%%" stop-color="#000" stop-opacity="0"/>
      </radialGradient>`, primaryHex)
	case "corner-glow":
		return fmt.Sprintf(`
      <radialGradient id="cornerGlow" cx="0%%" cy="0%%" r="60%%">
        <stop offset="0%%" stop-color="#%s" stop-opacity="0.15"/>
        <stop offset="100%%" stop-color="#000" stop-opacity="0"/>
      </radialGradient>`, primaryHex)
	case "dual-glow":
		return fmt.Sprintf(`
      <radialGradient id="dualGlow1" cx="20%%" cy="30%%" r="40%%">
        <stop offset="0%%" stop-color="#%s" stop-opacity="0.12"/>
        <stop offset="100%%" stop-color="#000" stop-opacity="0"/>
      </radialGradient>
      <radialGradient id="dualGlow2" cx="80%%" cy="70%%" r="40%%">
        <stop offset="0%%" stop-color="#%s" stop-opacity="0.12"/>
        <stop offset="100%%" stop-color="#000" stop-opacity="0"/>
      </radialGradient>`, primaryHex, secondaryHex)
	case "paper":
		return `
      <filter id="paperBg" x="0" y="0" width="100%" height="100%">
        <feTurbulence type="fractalNoise" baseFrequency="0.9" numOctaves="4" stitchTiles="stitch"/>
        <feColorMatrix type="saturate" values="0"/>
        <feComponentTransfer><feFuncA type="linear" slope="0.02"/></feComponentTransfer>
      </filter>`
	case "white":
		return ""
	case "black":
		return ""
	default:
		return ""
	}
}

// SceneBackground returns background SVG elements for a scene.
func SceneBackground(scene string) string {
	switch scene {
	case "spotlight":
		return `<rect width="100%" height="100%" fill="#050505"/><rect width="100%" height="100%" fill="url(#spotlightGrad)"/>`
	case "grid":
		return `<rect width="100%" height="100%" fill="#000"/><rect width="100%" height="100%" fill="url(#gridPattern)"/><rect width="100%" height="100%" fill="url(#gridGlow)"/>`
	case "pure":
		return `<rect width="100%" height="100%" fill="transparent"/>`
	case "split":
		return `<rect width="100%" height="100%" fill="#0a0a0c"/>`
	case "dots":
		return `<rect width="100%" height="100%" fill="#0a0a0c"/><rect width="100%" height="100%" fill="url(#dotsPattern)"/>`
	case "diagonal":
		return `<rect width="100%" height="100%" fill="#0a0a0c"/><rect width="100%" height="100%" fill="url(#diagPattern)"/>`
	case "noise-bg":
		return `<rect width="100%" height="100%" fill="#0a0a0c"/><rect width="100%" height="100%" filter="url(#bgNoise)"/>`
	case "gradient":
		return `<rect width="100%" height="100%" fill="#0a0a0c"/><rect width="100%" height="100%" fill="url(#bgGrad)"/>`
	case "radial":
		return `<rect width="100%" height="100%" fill="#0a0a0c"/><rect width="100%" height="100%" fill="url(#radialBg)"/>`
	case "vignette":
		return `<rect width="100%" height="100%" fill="#0a0a0c"/><rect width="100%" height="100%" fill="url(#vignetteBg)"/>`
	case "hexgrid":
		return `<rect width="100%" height="100%" fill="#0a0a0c"/><rect width="100%" height="100%" fill="url(#hexPattern)"/>`
	case "circuit":
		return `<rect width="100%" height="100%" fill="#0a0a0c"/><rect width="100%" height="100%" fill="url(#circuitPattern)"/>`
	case "crosshatch":
		return `<rect width="100%" height="100%" fill="#0a0a0c"/><rect width="100%" height="100%" fill="url(#crosshatchP)"/><rect width="100%" height="100%" fill="url(#crosshatchP2)"/>`
	case "stars":
		return `<rect width="100%" height="100%" fill="#050510"/><rect width="100%" height="100%" fill="url(#starsP)"/>`
	case "waves-bg":
		return `<rect width="100%" height="100%" fill="#0a0a0c"/><rect width="100%" height="100%" fill="url(#wavesP)"/>`
	case "diamond-grid":
		return `<rect width="100%" height="100%" fill="#0a0a0c"/><rect width="100%" height="100%" fill="url(#diamondP)"/>`
	case "brick":
		return `<rect width="100%" height="100%" fill="#0a0a0c"/><rect width="100%" height="100%" fill="url(#brickP)"/>`
	case "triangle-grid":
		return `<rect width="100%" height="100%" fill="#0a0a0c"/><rect width="100%" height="100%" fill="url(#triP)"/>`
	case "concentric":
		return `<rect width="100%" height="100%" fill="#0a0a0c"/><rect width="100%" height="100%" fill="url(#concentricBg)"/>`
	case "scanlines":
		return `<rect width="100%" height="100%" fill="#0a0a0c"/><rect width="100%" height="100%" fill="url(#scanlinesP)"/>`
	case "topography":
		return `<rect width="100%" height="100%" fill="#0a0a0c"/><rect width="100%" height="100%" fill="url(#topoP)"/>`
	case "plus-grid":
		return `<rect width="100%" height="100%" fill="#0a0a0c"/><rect width="100%" height="100%" fill="url(#plusP)"/>`
	case "isometric":
		return `<rect width="100%" height="100%" fill="#0a0a0c"/><rect width="100%" height="100%" fill="url(#isoP)"/>`
	case "aurora-bg":
		return `<rect width="100%" height="100%" fill="#050510"/><rect width="100%" height="100%" fill="url(#auroraBg1)"/>`
	case "spotlight-color":
		return `<rect width="100%" height="100%" fill="#050505"/><rect width="100%" height="100%" fill="url(#spotColorBg)"/>`
	case "corner-glow":
		return `<rect width="100%" height="100%" fill="#0a0a0c"/><rect width="100%" height="100%" fill="url(#cornerGlow)"/>`
	case "dual-glow":
		return `<rect width="100%" height="100%" fill="#050510"/><rect width="100%" height="100%" fill="url(#dualGlow1)"/><rect width="100%" height="100%" fill="url(#dualGlow2)"/>`
	case "paper":
		return `<rect width="100%" height="100%" fill="#f5f0e8"/><rect width="100%" height="100%" filter="url(#paperBg)"/>`
	case "white":
		return `<rect width="100%" height="100%" fill="#ffffff"/>`
	case "black":
		return `<rect width="100%" height="100%" fill="#000000"/>`
	default:
		return `<rect width="100%" height="100%" fill="#0a0a0c"/>`
	}
}

// SceneCount returns the total number of scenes.
func SceneCount() int { return 12 }
