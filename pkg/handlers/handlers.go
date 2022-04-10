package handlers

import (
	"github.com/nwohater/Go-WebApp/pkg/config"
	"github.com/nwohater/Go-WebApp/pkg/render"
	"net/http"
)

var Repo *Repository

// Repository Is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers Sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}

func (m *Repository) SuckIt(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "suckit.page.tmpl")
}
