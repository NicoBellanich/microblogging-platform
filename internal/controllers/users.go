package controllers

import (
	"encoding/json"
	"net/http"

	dtos "github.com/nicobellanich/migroblogging-platform/internal/controllers/DTOs"
	"github.com/nicobellanich/migroblogging-platform/internal/services"
	"github.com/nicobellanich/migroblogging-platform/internal/usecase"
)

type UsersController struct {
	UsecaseFollow             *usecase.Follow
	UsecasePublishMessage     *usecase.PublishMessage
	UsecaseObtainUserTimeline *usecase.ObtainUserTimeline
	UserService               services.IUserServices
}

func NewUsersController(
	ucf *usecase.Follow,
	ucp *usecase.PublishMessage,
	ucoutl *usecase.ObtainUserTimeline,
	us services.IUserServices,
) IUserController {
	return &UsersController{
		UsecaseFollow:             ucf,
		UsecasePublishMessage:     ucp,
		UsecaseObtainUserTimeline: ucoutl,
		UserService:               us,
	}
}

func (controller *UsersController) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req dtos.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.UserID == "" {
		http.Error(w, "Invalid request body (user_id)", http.StatusBadRequest)
		return
	}

	if err := controller.UserService.AddUser(req.UserID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))

}

func (controller *UsersController) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req dtos.GetUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.UserID == "" {
		http.Error(w, "Invalid request body (user_id)", http.StatusBadRequest)
		return
	}

	user, err := controller.UserService.GetUser(req.UserID)
	if err != nil {
		http.Error(w, "Failed to obtain user timeline", http.StatusInternalServerError)
		return
	}

	var following []string

	for _, f := range user.Following {
		following = append(following, f.Name)
	}

	resp := dtos.GetUserResponse{
		Name:         user.Name,
		Following:    following,
		Publications: user.Publications.GetContents(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

}

func (controller *UsersController) AddPublication(w http.ResponseWriter, r *http.Request) {
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

	if err := controller.UsecasePublishMessage.Execute(req.UserID, req.Content); err != nil {
		http.Error(w, "Failed to publish message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Message published successfully"))
}

func (controller *UsersController) AddFollowing(w http.ResponseWriter, r *http.Request) {
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

	if err := controller.UsecaseFollow.Execute(req.UserID, req.NewFollow); err != nil {
		http.Error(w, "Failed to follow", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("New Follower added successfully"))
}

func (controller *UsersController) GetTimeline(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req dtos.GetUserTimelineRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	messages, err := controller.UsecaseObtainUserTimeline.Execute(req.UserID)
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
