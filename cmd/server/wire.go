// Package server contains the dependency wiring and HTTP handler setup for the API server.
// The wire function initializes repositories, use cases, controllers, and HTTP routes.
package server

import (
	"net/http"

	"github.com/nicobellanich/migroblogging-platform/config"
	"github.com/nicobellanich/migroblogging-platform/internal/controllers"
	repository "github.com/nicobellanich/migroblogging-platform/internal/platform/repository/impl"
	"github.com/nicobellanich/migroblogging-platform/internal/usecase"
)

// wire sets up all dependencies and returns the HTTP handler mux.
func wire() http.Handler {
	mux := http.NewServeMux()

	// Load configuration (environment, etc.)
	conf := config.Load()

	// Infrastructure: initialize repositories based on environment
	messageRepository, err := repository.NewMessageRepository(conf)
	if err != nil {
		panic(err)
	}

	followersRepository, err := repository.NewFollowersRepository(conf)
	if err != nil {
		panic(err)
	}

	// Use Cases: business logic
	useCasePublishMessage := usecase.NewPublishMessage(messageRepository)
	usecaseFollow := usecase.NewFollow(followersRepository)
	usecaseObtainUserTimeline := usecase.NewObtainUserTimeline(followersRepository, messageRepository)

	// Controllers: HTTP handlers
	messageController := controllers.NewMessageController(useCasePublishMessage)
	followersController := controllers.NewFollowersController(usecaseFollow)
	userTimelineController := controllers.NewUserTimelineController(usecaseObtainUserTimeline)

	// HTTP Routes
	mux.HandleFunc("/publish", messageController.Publish)
	mux.HandleFunc("/follow", followersController.Follow)
	mux.HandleFunc("/usertimeline", userTimelineController.ObtainUserTimeline)

	return mux
}
