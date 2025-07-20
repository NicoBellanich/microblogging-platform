package server

import (
	"net/http"
)

func wire() http.Handler {
	mux := http.NewServeMux()

	return mux
}
