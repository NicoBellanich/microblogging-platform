package controllers

import (
	"encoding/json"
	"net/http"

	dtos "github.com/nicobellanich/migroblogging-platform/internal/controllers/DTOs"
	"github.com/nicobellanich/migroblogging-platform/internal/usecase"
)

type MessageController struct {
	UsecasePublishMessage *usecase.PublishMessage
}

func NewMessageController(ucpm *usecase.PublishMessage) IMessageController {
	return &MessageController{UsecasePublishMessage: ucpm}
}

func (mc *MessageController) Publish(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req dtos.CreatePublicationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.UserID == "" {
		http.Error(w, "Invalid request body (user_id)", http.StatusBadRequest)
		return
	}

	if req.Content == "" {
		http.Error(w, "Invalid request body (content)", http.StatusBadRequest)
		return
	}

	if err := mc.UsecasePublishMessage.Execute(req.UserID, req.Content); err != nil {
		http.Error(w, "Failed to publish message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Message published successfully"))

}
