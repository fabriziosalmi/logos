package svg

import (
	"embed"
	"fmt"
	"io/fs"
	"log/slog"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

// IconPack holds parsed SVG icon content from an embedded icon library.
// Icons are accessed via "pack:name" syntax (e.g. "lucide:rocket").
type IconPack struct {
	mu    sync.RWMutex
	icons map[string]string // "pack:name" -> SVG inner content
}

var globalIconPack = &IconPack{
	icons: make(map[string]string),
}

// svgTagRe matches the outer <svg> tag to extract inner content.
var svgTagRe = regexp.MustCompile(`(?s)<svg[^>]*>(.*)</svg>`)

// LoadIconPack parses all .svg files from an embedded FS directory
// and registers them under the given pack name.
// Lucide icons (24x24, stroke-based) are converted to fill-based shapes.
func LoadIconPack(packName string, iconFS embed.FS, dir string) int {
	sub, err := fs.Sub(iconFS, dir)
	if err != nil {
		slog.Warn("icon pack dir not found", "pack", packName, "dir", dir, "error", err)
		return 0
	}

	count := 0
	fs.WalkDir(sub, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(path, ".svg") {
			return nil
		}

		data, err := fs.ReadFile(sub, path)
		if err != nil {
			return nil
		}

		name := strings.TrimSuffix(filepath.Base(path), ".svg")
		inner := extractSVGInner(string(data))
		if inner == "" {
			return nil
		}

		key := packName + ":" + name
		globalIconPack.mu.Lock()
		globalIconPack.icons[key] = inner
		globalIconPack.mu.Unlock()
		count++
		return nil
	})

	slog.Info("icon pack loaded", "pack", packName, "count", count)
	return count
}

// extractSVGInner extracts everything between <svg> and </svg>.
func extractSVGInner(svgContent string) string {
	matches := svgTagRe.FindStringSubmatch(svgContent)
	if len(matches) < 2 {
		return ""
	}
	return strings.TrimSpace(matches[1])
}

// GetIcon returns the SVG inner content for a "pack:name" shape key.
func GetIcon(key string) (string, bool) {
	globalIconPack.mu.RLock()
	defer globalIconPack.mu.RUnlock()
	inner, ok := globalIconPack.icons[key]
	return inner, ok
}

// IconCount returns the total number of loaded icons across all packs.
func IconCount() int {
	globalIconPack.mu.RLock()
	defer globalIconPack.mu.RUnlock()
	return len(globalIconPack.icons)
}

// IconPackNames returns all loaded icon keys.
func IconPackNames() []string {
	globalIconPack.mu.RLock()
	defer globalIconPack.mu.RUnlock()
	names := make([]string, 0, len(globalIconPack.icons))
	for k := range globalIconPack.icons {
		names = append(names, k)
	}
	return names
}

// RenderIconShape wraps an icon pack's SVG content in our animation-compatible structure.
// Lucide/Heroicons use stroke-based 24x24 paths, so we scale to 32x32 and apply colors.
func RenderIconShape(key string, p ShapeParams) string {
	inner, ok := GetIcon(key)
	if !ok {
		return RenderShape("circle", p) // fallback
	}

	// Replace stroke="currentColor" and fill="none" with our dynamic colors
	inner = strings.ReplaceAll(inner, `stroke="currentColor"`, fmt.Sprintf(`stroke="%s"`, p.StrokeDef))
	inner = strings.ReplaceAll(inner, `fill="currentColor"`, fmt.Sprintf(`fill="%s"`, p.FillDef))

	// Wrap in our animation-compatible structure
	// Lucide icons are 24x24, we transform to center in our 32x32 grid
	return fmt.Sprintf(`
    <g class="logos-api-core">
      <g class="outer" transform="translate(4, 4)" fill="none" stroke="%s" stroke-width="%s" stroke-linecap="round" stroke-linejoin="round" vector-effect="non-scaling-stroke">
        %s
      </g>
    </g>`, p.StrokeDef, p.OuterStrokeWidth, inner)
}
