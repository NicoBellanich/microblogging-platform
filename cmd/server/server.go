package server

import (
	"fmt"
	"log"
	"net/http"
)

func Run() {
	mux := wire()
	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
