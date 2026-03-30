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
	default:
		return ""
	}
}

// SceneBackground returns background SVG elements for a scene.
func SceneBackground(scene string) string {
	switch scene {
	case "spotlight":
		return `
      <rect width="100%" height="100%" fill="#050505" />
      <rect width="100%" height="100%" fill="url(#spotlightGrad)" />`
	case "grid":
		return `
      <rect width="100%" height="100%" fill="#000000" />
      <rect width="100%" height="100%" fill="url(#gridPattern)" />
      <rect width="100%" height="100%" fill="url(#gridGlow)" />`
	case "pure":
		return `<rect width="100%" height="100%" fill="transparent" />`
	default:
		return `<rect width="100%" height="100%" fill="#0a0a0c" />`
	}
}

// ValidScene checks if a scene name is valid.
func ValidScene(name string) bool {
	switch name {
	case "pure", "spotlight", "grid", "split":
		return true
	}
	return false
}
