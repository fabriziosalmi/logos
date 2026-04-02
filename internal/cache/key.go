package cache

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

// Key generates a deterministic cache key from render parameters.
// Same params = same key = same SVG bytes. Always.
func Key(parts ...string) string {
	canonical := strings.Join(parts, "|")
	hash := sha256.Sum256([]byte(canonical))
	return hex.EncodeToString(hash[:16]) // 32-char hex
}

// KeyFromRequest builds a cache key from the typical render params.
func KeyFromRequest(format, scene, color1, color2, animation, theme, shape, texture, variant, title, subtitle, direction string, stroke, padding, alpha, speed float64, badge int, hover, decorative bool) string {
	return Key(
		format, scene, color1, color2, animation, theme, shape, texture, variant,
		title, subtitle, direction,
		fmt.Sprintf("%.1f", stroke),
		fmt.Sprintf("%.1f", padding),
		fmt.Sprintf("%.2f", alpha),
		fmt.Sprintf("%.2f", speed),
		fmt.Sprintf("%d", badge),
		fmt.Sprintf("%v", hover),
		fmt.Sprintf("%v", decorative),
	)
}
