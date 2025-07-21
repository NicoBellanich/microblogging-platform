package controllers

import (
	"encoding/json"
	"net/http"

	dtos "github.com/nicobellanich/migroblogging-platform/internal/controllers/DTOs"
	"github.com/nicobellanich/migroblogging-platform/internal/usecase"
)

type FollowersController struct {
	UsecaseFollow *usecase.Follow
}

func NewFollowersController(ucf *usecase.Follow) IFollowersController {
	return &FollowersController{
		UsecaseFollow: ucf,
	}
}

func (fc *FollowersController) Follow(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req dtos.FollowRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.UserID == "" {
		http.Error(w, "Invalid request body (user_id)", http.StatusBadRequest)
		return
	}

	if req.NewFollow == "" {
		http.Error(w, "Invalid request body (new_follow)", http.StatusBadRequest)
		return
	}

	if err := fc.UsecaseFollow.Execute(req.UserID, req.NewFollow); err != nil {
		http.Error(w, "Failed to follow", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("New Follower added successfully"))

}
