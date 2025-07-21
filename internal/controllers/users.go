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
		writeJSONError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req dtos.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.UserID == "" {
		writeJSONError(w, "Invalid request body (user_id)", http.StatusBadRequest)
		return
	}

	if err := controller.UserService.AddUser(req.UserID); err != nil {
		writeJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "User created successfully", "status": http.StatusCreated})

}

func (controller *UsersController) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSONError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req dtos.GetUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.UserID == "" {
		writeJSONError(w, "Invalid request body (user_id)", http.StatusBadRequest)
		return
	}

	user, err := controller.UserService.GetUser(req.UserID)
	if err != nil {
		writeJSONError(w, err.Error(), http.StatusInternalServerError)
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
		writeJSONError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req dtos.CreatePublicationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.UserID == "" {
		writeJSONError(w, "Invalid request body (user_id)", http.StatusBadRequest)
		return
	}

	if req.Content == "" {
		writeJSONError(w, "Invalid request body (content)", http.StatusBadRequest)
		return
	}

	if err := controller.UsecasePublishMessage.Execute(req.UserID, req.Content); err != nil {
		writeJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Message published successfully", "status": http.StatusCreated})
}

func (controller *UsersController) AddFollowing(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSONError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req dtos.FollowRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.UserID == "" {
		writeJSONError(w, "Invalid request body (user_id)", http.StatusBadRequest)
		return
	}

	if req.NewFollow == "" {
		writeJSONError(w, "Invalid request body (new_follow)", http.StatusBadRequest)
		return
	}

	if err := controller.UsecaseFollow.Execute(req.UserID, req.NewFollow); err != nil {
		writeJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "New Follower added successfully", "status": http.StatusCreated})
}

func (controller *UsersController) GetTimeline(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSONError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req dtos.GetUserTimelineRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	feed, err := controller.UsecaseObtainUserTimeline.Execute(req.UserID)
	if err != nil {
		writeJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	messageResponse := make([]dtos.MessageResponse, 0)
	for _, msg := range feed.GetAllMessages() {
		messageResponse = append(messageResponse, dtos.MessageResponse{
			ID:        msg.ID(),
			UserID:    msg.UserID(),
			Content:   msg.Content(),
			CreatedAt: msg.CreatedAt().Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	resp := dtos.GetUserTimelineResponse{
		Feed: messageResponse,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

// writeJSONError writes a JSON error response with the given message and status code
func writeJSONError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error":  message,
		"status": status,
	})
}
