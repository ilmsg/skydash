package handler

import (
	"net/http"
	"path/filepath"
	"text/template"
)

type IndexHandler struct {
}

func (h *IndexHandler) Index(w http.ResponseWriter, r *http.Request) {
	Render(w, nil, []string{"index.html"})
}

func (h *IndexHandler) Table(w http.ResponseWriter, r *http.Request) {
	Render(w, nil, []string{"table.html"})
}

func (h *IndexHandler) UIButton(w http.ResponseWriter, r *http.Request) {
	Render(w, nil, []string{filepath.Join("ui-element", "button.html")})
}

func (h *IndexHandler) UIDropdown(w http.ResponseWriter, r *http.Request) {
	Render(w, nil, []string{filepath.Join("ui-element", "dropdown.html")})
}

func (h *IndexHandler) UITypography(w http.ResponseWriter, r *http.Request) {
	Render(w, nil, []string{filepath.Join("ui-element", "typography.html")})
}

func (h *IndexHandler) Notfound(w http.ResponseWriter, r *http.Request) {
	Render(w, nil, []string{"404.html"})
}

func (h *IndexHandler) Error(w http.ResponseWriter, r *http.Request) {
	Render(w, nil, []string{"500.html"})
}

func Render(w http.ResponseWriter, data any, filenames []string) {
	tmps := []string{
		filepath.Join("templates", "layout.html"),
		filepath.Join("templates/partials", "container.html"),
		filepath.Join("templates/partials", "sidebar.html"),
		filepath.Join("templates/partials", "footer.html"),
	}
	for _, filename := range filenames {
		tmps = append(tmps, filepath.Join("templates", filename))
	}
	tmpl := template.Must(template.ParseFiles(tmps...))
	tmpl.ExecuteTemplate(w, "layout", data)
}
