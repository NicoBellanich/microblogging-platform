package server

import (
	"net/http"

	"github.com/nicobellanich/migroblogging-platform/config"
	repository "github.com/nicobellanich/migroblogging-platform/internal/platform/repository/impl"
	"github.com/nicobellanich/migroblogging-platform/internal/usecase"
)

func wire() http.Handler {
	mux := http.NewServeMux()

	conf := config.Load()

	// Infra
	messageRepository := repository.NewMessageRepository(conf)

	// Services

	// UC
	usecase.NewPublishMessage(messageRepository)

	// Controllers

	return mux
}
