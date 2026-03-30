package web

import "embed"

//go:embed all:dashboard
//go:embed all:static
var FS embed.FS
