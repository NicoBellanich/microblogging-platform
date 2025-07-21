// Package main is the entry point for the microblogging platform API server.
// It initializes and starts the HTTP server.
package main

import (
	"github.com/nicobellanich/migroblogging-platform/cmd/server"
)

// main starts the API server.
func main() {
	server.Run()
}
