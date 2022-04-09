package main

import (
	"net/http"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "aboot.page.html")
}

func SuckIt(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "suckit.page.html")
}
