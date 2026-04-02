package handler

import (
	"embed"
	"io/fs"
	"log/slog"
	"net/http"
)

// Dashboard serves the embedded web dashboard.
// embed.FS paths start with "dashboard/" since the FS root is the web/ package.
// fs.Sub errors are safe to ignore: embed.FS subdirs are compile-time guaranteed.
func Dashboard(webFS embed.FS) http.HandlerFunc {
	sub, err := fs.Sub(webFS, "dashboard")
	if err != nil {
		slog.Error("embedded dashboard dir missing", "error", err)
		panic("dashboard embed broken")
	}
	fileServer := http.FileServer(http.FS(sub))

	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			data, err := webFS.ReadFile("dashboard/index.html")
			if err != nil {
				http.Error(w, "dashboard not found", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(data)
			return
		}
		fileServer.ServeHTTP(w, r)
	}
}

// StaticFiles serves pre-generated SVGs from the embedded static/ dir.
func StaticFiles(webFS embed.FS) http.Handler {
	sub, err := fs.Sub(webFS, "static")
	if err != nil {
		slog.Error("embedded static dir missing", "error", err)
		panic("static embed broken")
	}
	return http.FileServer(http.FS(sub))
}
