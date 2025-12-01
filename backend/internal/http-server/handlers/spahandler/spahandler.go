package spahandler

import (
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// New serves files from staticDir and falls back to indexFile for SPA routes.
func New(log *slog.Logger, staticDir string, indexFile string) http.HandlerFunc {
	root := filepath.Clean(staticDir)
	indexPath := filepath.Join(root, indexFile)

	return func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Clean(strings.TrimPrefix(r.URL.Path, "/"))

		if path == "." || path == "/" || path == "" {
			http.ServeFile(w, r, indexPath)
			return
		}

		requested := filepath.Join(root, path)
		if !strings.HasPrefix(requested, root+string(os.PathSeparator)) && requested != root {
			http.NotFound(w, r)
			return
		}

		info, err := os.Stat(requested)
		if err == nil && !info.IsDir() {
			http.ServeFile(w, r, requested)
			return
		}

		if err != nil && !os.IsNotExist(err) {
			log.Warn("failed to stat asset", slog.String("path", requested), slog.String("err", err.Error()))
		}

		http.ServeFile(w, r, indexPath)
	}
}
