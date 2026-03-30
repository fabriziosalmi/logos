package svg

// ThemeParams holds resolved theme values for SVG rendering.
type ThemeParams struct {
	CoreBg      string
	DotColor    string
	StrokeWidth string
}

// ResolveTheme returns fill/stroke params for a given theme.
func ResolveTheme(theme, fillDef string) ThemeParams {
	switch theme {
	case "dark":
		return ThemeParams{CoreBg: "#0a0a0c", DotColor: "#ffffff", StrokeWidth: "1.5"}
	case "light":
		return ThemeParams{CoreBg: "#ffffff", DotColor: "#0a0a0c", StrokeWidth: "1.5"}
	case "solid":
		return ThemeParams{CoreBg: fillDef, DotColor: "#ffffff", StrokeWidth: "0"}
	case "glass":
		return ThemeParams{CoreBg: "rgba(255,255,255,0.1)", DotColor: "#ffffff", StrokeWidth: "1.5"}
	default: // auto
		return ThemeParams{CoreBg: "transparent", DotColor: "#fff", StrokeWidth: "1.5"}
	}
}

// ValidTheme checks if a theme name is valid.
func ValidTheme(name string) bool {
	switch name {
	case "dark", "light", "solid", "glass", "auto":
		return true
	}
	return false
}
