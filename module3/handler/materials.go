package handler

import "net/http"
import "github.com/vlpolak/swtgo/module3/handler/template"

func MaterialsHandler(w http.ResponseWriter, r *http.Request) {
	template.RenderTemplate(w, "materials")
}
