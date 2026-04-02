package svg

// Format defines output dimensions.
type Format struct {
	Width  int
	Height int
}

// Formats: 40 standard sizes covering favicons, social, devices, print.
var Formats = map[string]Format{
	// Core
	"favicon":    {32, 32},
	"avatar":     {512, 512},
	"og-card":    {1200, 630},
	"hero":       {1920, 1080},

	// Favicons & Icons
	"icon-16":    {16, 16},
	"icon-32":    {32, 32},
	"icon-48":    {48, 48},
	"icon-64":    {64, 64},
	"icon-96":    {96, 96},
	"icon-128":   {128, 128},
	"icon-192":   {192, 192},
	"icon-256":   {256, 256},
	"icon-512":   {512, 512},

	// Apple
	"apple-touch":     {180, 180},
	"apple-splash":    {2048, 2732},

	// Social Media
	"twitter-card":    {1200, 628},
	"twitter-avatar":  {400, 400},
	"facebook-cover":  {1640, 856},
	"facebook-post":   {1200, 630},
	"instagram-post":  {1080, 1080},
	"instagram-story": {1080, 1920},
	"linkedin-cover":  {1584, 396},
	"linkedin-post":   {1200, 627},
	"youtube-thumb":   {1280, 720},

	// Chat & Collab
	"slack-emoji":  {128, 128},
	"discord-icon": {512, 512},
	"teams-icon":   {192, 192},
	"notion-icon":  {280, 280},

	// Devices
	"desktop-hd":  {1920, 1080},
	"desktop-4k":  {3840, 2160},
	"tablet":      {1024, 768},
	"mobile":      {375, 812},

	// Standard Squares
	"square-sm":  {64, 64},
	"square-md":  {256, 256},
	"square-lg":  {512, 512},
	"square-xl":  {1024, 1024},

	// Banners
	"banner-sm":  {728, 90},
	"banner-md":  {1200, 300},
	"banner-lg":  {1920, 480},

	// PWA Manifest
	"pwa-192": {192, 192},
	"pwa-512": {512, 512},
}

// FormatCount returns the total number of formats.
func FormatCount() int {
	return len(Formats)
}
