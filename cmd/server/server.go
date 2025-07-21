// Package server provides the HTTP server setup for the microblogging platform.
package server

import (
	"fmt"
	"log"
	"net/http"
)

// Run initializes the HTTP handlers and starts the server on port 8080.
func Run() {
	mux := wire()
	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
