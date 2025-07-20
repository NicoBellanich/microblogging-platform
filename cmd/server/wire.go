package server

import (
	"net/http"

	"github.com/nicobellanich/migroblogging-platform/config"
)

func wire() http.Handler {
	mux := http.NewServeMux()

	config.Load()

	return mux
}
