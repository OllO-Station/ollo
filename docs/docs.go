package docs

import (
	"embed"
	"html/template"
	"net/http"
)

//go:embed index.tpl
var Index embed.FS

//go:embed static
var File embed.FS 

var FileHandler = http.FileServer(http.FS(File))

// Handler returns an http handler that servers OpenAPI console for an OpenAPI spec at specURL.
func DocsHandler(title, specURL string) http.HandlerFunc {
	t, _ := template.ParseFS(Index, "index.tpl")

	return func(w http.ResponseWriter, req *http.Request) {
		t.Execute(w, struct {
			Title string
			URL   string
		}{
			title,
			specURL,
		})
	}
}
