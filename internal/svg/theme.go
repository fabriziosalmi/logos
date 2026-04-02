package svg

// ThemeParams holds resolved theme values for SVG rendering.
type ThemeParams struct {
	CoreBg      string
	DotColor    string
	StrokeWidth string
}

// ResolveTheme returns fill/stroke params for a given theme.
// 15 themes: 5 core + 10 presets inspired by popular editor/terminal themes.
func ResolveTheme(theme, fillDef string) ThemeParams {
	switch theme {
	// Core themes
	case "dark":
		return ThemeParams{CoreBg: "#0a0a0c", DotColor: "#ffffff", StrokeWidth: "1.5"}
	case "light":
		return ThemeParams{CoreBg: "#ffffff", DotColor: "#0a0a0c", StrokeWidth: "1.5"}
	case "solid":
		return ThemeParams{CoreBg: fillDef, DotColor: "#ffffff", StrokeWidth: "0"}
	case "glass":
		return ThemeParams{CoreBg: "rgba(255,255,255,0.1)", DotColor: "#ffffff", StrokeWidth: "1.5"}

	// Editor/Terminal presets
	case "monokai":
		return ThemeParams{CoreBg: "#272822", DotColor: "#f8f8f2", StrokeWidth: "1.5"}
	case "dracula":
		return ThemeParams{CoreBg: "#282a36", DotColor: "#f8f8f2", StrokeWidth: "1.5"}
	case "nord":
		return ThemeParams{CoreBg: "#2e3440", DotColor: "#eceff4", StrokeWidth: "1.5"}
	case "solarized-dark":
		return ThemeParams{CoreBg: "#002b36", DotColor: "#fdf6e3", StrokeWidth: "1.5"}
	case "solarized-light":
		return ThemeParams{CoreBg: "#fdf6e3", DotColor: "#002b36", StrokeWidth: "1.5"}
	case "gruvbox":
		return ThemeParams{CoreBg: "#282828", DotColor: "#ebdbb2", StrokeWidth: "1.5"}
	case "catppuccin":
		return ThemeParams{CoreBg: "#1e1e2e", DotColor: "#cdd6f4", StrokeWidth: "1.5"}
	case "tokyo-night":
		return ThemeParams{CoreBg: "#1a1b26", DotColor: "#c0caf5", StrokeWidth: "1.5"}
	case "one-dark":
		return ThemeParams{CoreBg: "#282c34", DotColor: "#abb2bf", StrokeWidth: "1.5"}
	case "github-dark":
		return ThemeParams{CoreBg: "#0d1117", DotColor: "#e6edf3", StrokeWidth: "1.5"}
	case "github-light":
		return ThemeParams{CoreBg: "#ffffff", DotColor: "#1f2328", StrokeWidth: "1.5"}
	case "ayu-dark":
		return ThemeParams{CoreBg: "#0b0e14", DotColor: "#bfbdb6", StrokeWidth: "1.5"}
	case "ayu-light":
		return ThemeParams{CoreBg: "#fcfcfc", DotColor: "#5c6166", StrokeWidth: "1.5"}
	case "rose-pine":
		return ThemeParams{CoreBg: "#191724", DotColor: "#e0def4", StrokeWidth: "1.5"}
	case "everforest":
		return ThemeParams{CoreBg: "#2d353b", DotColor: "#d3c6aa", StrokeWidth: "1.5"}
	case "kanagawa":
		return ThemeParams{CoreBg: "#1f1f28", DotColor: "#dcd7ba", StrokeWidth: "1.5"}

	// === NEW THEMES (30 more) ===

	// Terminal/Editor
	case "material":
		return ThemeParams{CoreBg: "#263238", DotColor: "#eeffff", StrokeWidth: "1.5"}
	case "material-light":
		return ThemeParams{CoreBg: "#fafafa", DotColor: "#546e7a", StrokeWidth: "1.5"}
	case "night-owl":
		return ThemeParams{CoreBg: "#011627", DotColor: "#d6deeb", StrokeWidth: "1.5"}
	case "poimandres":
		return ThemeParams{CoreBg: "#1b1e28", DotColor: "#a6accd", StrokeWidth: "1.5"}
	case "vesper":
		return ThemeParams{CoreBg: "#101010", DotColor: "#b0b0b0", StrokeWidth: "1.5"}
	case "synthwave-84":
		return ThemeParams{CoreBg: "#262335", DotColor: "#e0def4", StrokeWidth: "1.5"}
	case "cobalt2":
		return ThemeParams{CoreBg: "#193549", DotColor: "#e1efff", StrokeWidth: "1.5"}
	case "palenight":
		return ThemeParams{CoreBg: "#292d3e", DotColor: "#a6accd", StrokeWidth: "1.5"}
	case "shades-of-purple":
		return ThemeParams{CoreBg: "#2d2b55", DotColor: "#e0def4", StrokeWidth: "1.5"}
	case "atom-one-light":
		return ThemeParams{CoreBg: "#fafafa", DotColor: "#383a42", StrokeWidth: "1.5"}
	case "high-contrast":
		return ThemeParams{CoreBg: "#000000", DotColor: "#ffffff", StrokeWidth: "2"}
	case "high-contrast-light":
		return ThemeParams{CoreBg: "#ffffff", DotColor: "#000000", StrokeWidth: "2"}

	// Aesthetic/Social
	case "midnight":
		return ThemeParams{CoreBg: "#0f0f1a", DotColor: "#8888cc", StrokeWidth: "1.5"}
	case "sunset-theme":
		return ThemeParams{CoreBg: "#1a0a0a", DotColor: "#ff8866", StrokeWidth: "1.5"}
	case "ocean-theme":
		return ThemeParams{CoreBg: "#0a1a2a", DotColor: "#66ccff", StrokeWidth: "1.5"}
	case "forest-theme":
		return ThemeParams{CoreBg: "#0a1a0a", DotColor: "#88cc88", StrokeWidth: "1.5"}
	case "lavender-theme":
		return ThemeParams{CoreBg: "#1a1020", DotColor: "#cc99ff", StrokeWidth: "1.5"}
	case "ember":
		return ThemeParams{CoreBg: "#1a0800", DotColor: "#ff6633", StrokeWidth: "1.5"}
	case "arctic":
		return ThemeParams{CoreBg: "#f0f5fa", DotColor: "#2255aa", StrokeWidth: "1.5"}
	case "sandstorm":
		return ThemeParams{CoreBg: "#2a2010", DotColor: "#ddc088", StrokeWidth: "1.5"}
	case "cherry":
		return ThemeParams{CoreBg: "#1a0010", DotColor: "#ff3366", StrokeWidth: "1.5"}
	case "matrix-theme":
		return ThemeParams{CoreBg: "#000800", DotColor: "#00ff41", StrokeWidth: "1.5"}

	// Platform
	case "slack":
		return ThemeParams{CoreBg: "#1a1d21", DotColor: "#d1d2d3", StrokeWidth: "1.5"}
	case "discord":
		return ThemeParams{CoreBg: "#36393f", DotColor: "#dcddde", StrokeWidth: "1.5"}
	case "notion":
		return ThemeParams{CoreBg: "#191919", DotColor: "#e0e0e0", StrokeWidth: "1.5"}
	case "notion-light":
		return ThemeParams{CoreBg: "#ffffff", DotColor: "#37352f", StrokeWidth: "1.5"}
	case "linear":
		return ThemeParams{CoreBg: "#191a23", DotColor: "#f2f2f2", StrokeWidth: "1.5"}
	case "vercel":
		return ThemeParams{CoreBg: "#000000", DotColor: "#ffffff", StrokeWidth: "1.5"}
	case "stripe":
		return ThemeParams{CoreBg: "#0a2540", DotColor: "#e3e8ee", StrokeWidth: "1.5"}
	case "supabase":
		return ThemeParams{CoreBg: "#1c1c1c", DotColor: "#3ecf8e", StrokeWidth: "1.5"}

	default: // auto
		return ThemeParams{CoreBg: "transparent", DotColor: "#fff", StrokeWidth: "1.5"}
	}
}

// ThemeCount returns the total number of themes.
func ThemeCount() int { return 51 }
