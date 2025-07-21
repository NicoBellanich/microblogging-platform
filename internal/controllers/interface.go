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

type IUserController interface {
	Create(http.ResponseWriter, *http.Request)
	GetUserByUsername(http.ResponseWriter, *http.Request)
	AddPublication(http.ResponseWriter, *http.Request)
	GetTimeline(http.ResponseWriter, *http.Request)
	AddFollowing(http.ResponseWriter, *http.Request)
}
