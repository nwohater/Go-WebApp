package main

import (
	"fmt"
	"github.com/nwohater/Go-WebApp/pkg/config"
	"github.com/nwohater/Go-WebApp/pkg/handlers"
	"github.com/nwohater/Go-WebApp/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

// Funky main
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	http.HandleFunc("/suckit", handlers.Repo.SuckIt)
	fmt.Printf("Starting Application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
