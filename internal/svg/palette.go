package svg

import "regexp"

var hexPattern = regexp.MustCompile(`^[0-9A-Fa-f]{3,6}$`)

// BaseColors are the primary palette.
var BaseColors = map[string]string{
	"amber": "eab308", "blue": "3b82f6", "cyan": "06b6d4", "green": "22c55e",
	"indigo": "6366f1", "orange": "f97316", "purple": "a855f7", "rose": "f43f5e",
}

// CreativeColors are the extended palette.
var CreativeColors = map[string]string{
	"white": "ffffff", "gray": "9ca3af", "black": "111111", "gold": "d4af37",
	"platinum": "e5e4e2", "champagne": "f7e7ce", "neon": "39ff14", "matrix": "00ff00",
	"cyber": "ff00ff", "laser": "ff0099", "plasma": "00ffff", "void": "8a2be2",
	"emerald": "50c878", "sapphire": "0f52ba", "ruby": "e0115f", "ocean": "006994",
	"sunset": "fd5e53", "magma": "ff3300", "mint": "98ff98", "peach": "ffdab9",
	"lavender": "e6e6fa",
}

// AllColors merges base + creative.
var AllColors map[string]string

func init() {
	AllColors = make(map[string]string, len(BaseColors)+len(CreativeColors))
	for k, v := range BaseColors {
		AllColors[k] = v
	}
	for k, v := range CreativeColors {
		AllColors[k] = v
	}
}

// ResolveHex takes a color name or raw hex and returns a 6-char hex string.
func ResolveHex(val string) string {
	if hex, ok := AllColors[val]; ok {
		return hex
	}
	if hexPattern.MatchString(val) {
		return val
	}
	return "ffffff"
}

// ColorNames returns a sorted list of all palette names for the dashboard.
func ColorNames() []string {
	// Base first, then creative, deterministic order
	names := make([]string, 0, len(AllColors))
	baseOrder := []string{"amber", "blue", "cyan", "green", "indigo", "orange", "purple", "rose"}
	creativeOrder := []string{
		"white", "gray", "black", "gold", "platinum", "champagne",
		"neon", "matrix", "cyber", "laser", "plasma", "void",
		"emerald", "sapphire", "ruby", "ocean", "sunset", "magma",
		"mint", "peach", "lavender",
	}
	names = append(names, baseOrder...)
	names = append(names, creativeOrder...)
	return names
}
