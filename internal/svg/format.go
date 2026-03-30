package svg

// Format defines output dimensions.
type Format struct {
	Width  int
	Height int
}

var Formats = map[string]Format{
	"favicon": {32, 32},
	"avatar":  {512, 512},
	"og-card": {1200, 630},
	"hero":    {1920, 1080},
}

// FormatNames returns valid format keys.
func FormatNames() []string {
	return []string{"favicon", "avatar", "og-card", "hero"}
}

// ValidFormat checks if a format name is valid.
func ValidFormat(name string) bool {
	_, ok := Formats[name]
	return ok
}
