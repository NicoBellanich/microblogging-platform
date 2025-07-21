package controllers

import "net/http"

type IMessageController interface {
	Publish(http.ResponseWriter, *http.Request)
}

type IFollowersController interface {
	Follow(http.ResponseWriter, *http.Request)
}

type IUserTimeline interface {
	ObtainUserTimeline(http.ResponseWriter, *http.Request)
}
