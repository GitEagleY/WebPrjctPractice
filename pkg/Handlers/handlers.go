package handlers

import (
	"net/http"

	render "github.com/GitEagleY/WebPrjctPractice/pkg/Render"
	"github.com/GitEagleY/WebPrjctPractice/pkg/config"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(a *config.AppConfig) *Repository { //need for creating new Repo
	return &Repository{
		App: a,
	}
}
func NewHandlers(r *Repository) { //need for taking here repo to work with
	Repo = r
}
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}
