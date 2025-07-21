// Package server contains the dependency wiring and HTTP handler setup for the API server.
// The wire function initializes repositories, use cases, controllers, and HTTP routes.
package server

import (
	"net/http"

	"github.com/nicobellanich/migroblogging-platform/config"
	"github.com/nicobellanich/migroblogging-platform/internal/controllers"
	repository "github.com/nicobellanich/migroblogging-platform/internal/platform/repository/impl"
	"github.com/nicobellanich/migroblogging-platform/internal/services"
	"github.com/nicobellanich/migroblogging-platform/internal/usecase"
)

// wire sets up all dependencies and returns the HTTP handler mux.
func wire() http.Handler {
	mux := http.NewServeMux()

	// Load configuration (environment, etc.)
	conf := config.Load()

	// Infrastructure: initialize repositories based on environment
	usersRepository, err := repository.NewUsersRepository(conf)
	if err != nil {
		panic(err)
	}

	// Services : different services uses to help Usecases
	usersService := services.NewUserServices(usersRepository)

	// Use Cases: business logic
	useCasePublishMessage := usecase.NewPublishMessage(usersService)
	usecaseFollow := usecase.NewFollow(usersService)
	usecaseObtainUserTimeline := usecase.NewObtainUserTimeline(usersService)

	// Controllers: HTTP handlers
	usersController := controllers.NewUsersController(usecaseFollow, useCasePublishMessage, usecaseObtainUserTimeline, usersService)

	// HTTP Routes
	mux.HandleFunc("/user", usersController.GetUserByUsername)
	mux.HandleFunc("/user/create", usersController.Create)
	mux.HandleFunc("/user/timeline", usersController.GetTimeline)
	mux.HandleFunc("/user/publish", usersController.AddPublication)
	mux.HandleFunc("/user/following", usersController.AddFollowing)

	return mux
}
