package handler

import (
	"embed"
	"io/fs"
	"net/http"
)

// Dashboard serves the embedded web dashboard.
// The embed.FS root is the web/ package dir, so paths start with "dashboard/".
func Dashboard(webFS embed.FS) http.HandlerFunc {
	sub, _ := fs.Sub(webFS, "dashboard")
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
	sub, _ := fs.Sub(webFS, "static")
	return http.FileServer(http.FS(sub))
}
