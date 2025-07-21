package server

import (
	"net/http"

	"github.com/nicobellanich/migroblogging-platform/config"
	"github.com/nicobellanich/migroblogging-platform/internal/controllers"
	repository "github.com/nicobellanich/migroblogging-platform/internal/platform/repository/impl"
	"github.com/nicobellanich/migroblogging-platform/internal/usecase"
)

func wire() http.Handler {
	mux := http.NewServeMux()

	conf := config.Load()

	// Infra
	messageRepository, err := repository.NewMessageRepository(conf)
	if err != nil {
		panic(err)
	}

	// Services
	// ...

	// UC
	useCasePublishMessage := usecase.NewPublishMessage(messageRepository)

	// Controllers
	messageController := controllers.NewMessageController(useCasePublishMessage)

	// Handlers
	mux.HandleFunc("/publish", messageController.Publish)

	return mux
}
