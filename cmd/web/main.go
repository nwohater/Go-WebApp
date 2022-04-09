package main

import (
	"fmt"
	"github.com/nwohater/webapp/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

// Funky main
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/suckit", handlers.SuckIt)
	fmt.Printf("Starting Application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
