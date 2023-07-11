package handlers

import (
	"net/http"

	render "github.com/GitEagleY/WebPrjctPractice/pkg/Render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}
