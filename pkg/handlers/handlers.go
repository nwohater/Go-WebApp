package handlers

import (
	"github.com/nwohater/Go-WebApp/pkg/render"
	"net/http"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}

func SuckIt(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "suckit.page.html")
}
