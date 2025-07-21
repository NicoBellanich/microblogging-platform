package controllers

import "net/http"

type IMessageController interface {
	Publish(http.ResponseWriter, *http.Request)
}
