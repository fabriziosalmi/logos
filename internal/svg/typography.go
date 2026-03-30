package svg

import (
	"fmt"
	"html"
)

// TextParams holds typography configuration.
type TextParams struct {
	Title    string
	Subtitle string
}

// RenderTypography returns SVG text elements for title/subtitle.
func RenderTypography(text TextParams, format, scene, primaryHex string) string {
	if text.Title == "" {
		return ""
	}

	title := html.EscapeString(text.Title)
	subtitle := html.EscapeString(text.Subtitle)

	titleSize := 48
	subSize := 24
	x := "50%"
	y := "75%"
	textAnchor := "middle"

	if format == "og-card" && scene == "split" {
		titleSize = 82
		subSize = 36
		x = "80px"
		y = "50%"
		textAnchor = "start"
	} else if format == "hero" {
		titleSize = 120
		subSize = 48
		y = "80%"
	}

	titleYOffset := -10
	if subtitle == "" {
		titleYOffset = 10
	}
	subYOffset := titleSize/2 + 20

	result := fmt.Sprintf(`
    <g class="typography">
      <text x="%s" y="%s" dy="%dpx" text-anchor="%s" fill="#ffffff" font-family="-apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif" font-size="%dpx" font-weight="800" letter-spacing="-0.03em">%s</text>`,
		x, y, titleYOffset, textAnchor, titleSize, title)

	if subtitle != "" {
		result += fmt.Sprintf(`
      <text x="%s" y="%s" dy="%dpx" text-anchor="%s" fill="#%s" font-family="-apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif" font-size="%dpx" font-weight="500" opacity="0.8">%s</text>`,
			x, y, subYOffset, textAnchor, primaryHex, subSize, subtitle)
	}

	result += "\n    </g>"
	return result
}
