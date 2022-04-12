package main

import (
	"github.com/bmizerany/pat"
	"github.com/nwohater/Go-WebApp/pkg/config"
	"github.com/nwohater/Go-WebApp/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux.Get("/suckit", http.HandlerFunc(handlers.Repo.SuckIt))
	return mux
}
