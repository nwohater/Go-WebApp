package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Funky main
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/suckit", SuckIt)
	fmt.Printf("Starting Application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
