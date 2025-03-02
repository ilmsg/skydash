package handler

import (
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/lafriks/xormstore"
)

type IndexHandler struct {
	Store *xormstore.Store
}

func (h *IndexHandler) Index(w http.ResponseWriter, r *http.Request) {
	Render(w, nil, []string{"index.html"})
}

func Render(w http.ResponseWriter, data any, filenames []string) {
	tmps := []string{filepath.Join("templates", "layout.html")}
	for _, filename := range filenames {
		tmps = append(tmps, filepath.Join("templates", filename))
	}
	tmpl := template.Must(template.ParseFiles(tmps...))
	tmpl.ExecuteTemplate(w, "layout", data)
}
