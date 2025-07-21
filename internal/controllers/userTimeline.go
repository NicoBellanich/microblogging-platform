package controllers

import (
	"encoding/json"
	"net/http"

	dtos "github.com/nicobellanich/migroblogging-platform/internal/controllers/DTOs"
	"github.com/nicobellanich/migroblogging-platform/internal/usecase"
)

type UserTimelineController struct {
	UsecaseObtainUserTimeline *usecase.ObtainUserTimeline
}

func NewUserTimelineController(ucutc *usecase.ObtainUserTimeline) IUserTimeline {
	return &UserTimelineController{
		UsecaseObtainUserTimeline: ucutc,
	}
}

func (ctrlr *UserTimelineController) ObtainUserTimeline(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req dtos.GetUserTimelineRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	messages, err := ctrlr.UsecaseObtainUserTimeline.Execute(req.UserID)
	if err != nil {
		http.Error(w, "Failed to obtain user timeline", http.StatusInternalServerError)
		return
	}

	resp := dtos.GetUserTimelineResponse{
		Messages: messages,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
