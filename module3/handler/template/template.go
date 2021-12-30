package template

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("static/template/materials.html"))

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	err := templates.Execute(w, "materials.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
