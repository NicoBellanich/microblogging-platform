package controllers

import (
	"encoding/json"
	"net/http"

	dtos "github.com/nicobellanich/migroblogging-platform/internal/controllers/DTOs"
	"github.com/nicobellanich/migroblogging-platform/internal/domain"
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
		writeJSONError(w, domain.NewAppError("[CONTROLLER]", domain.ErrMethodNotAllowed, "method="+r.Method))
		return
	}

	var req dtos.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, domain.NewAppError("[CONTROLLER]", domain.ErrInvalidRequestBody, ""))
		return
	}

	if req.UserName == "" {
		writeJSONError(w, domain.NewAppError("[CONTROLLER]", domain.ErrInvalidRequestBody, "user_id"))
		return
	}

	if err := controller.UserService.AddUser(req.UserName); err != nil {
		writeJSONError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "User created successfully", "status": http.StatusCreated})

}

func (controller *UsersController) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSONError(w, domain.NewAppError("[CONTROLLER]", domain.ErrMethodNotAllowed, "method="+r.Method))
		return
	}

	var req dtos.GetUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, domain.NewAppError("[CONTROLLER]", domain.ErrInvalidRequestBody, ""))
		return
	}

	if req.UserName == "" {
		writeJSONError(w, domain.NewAppError("[CONTROLLER]", domain.ErrInvalidRequestBody, "user_id"))
		return
	}

	user, err := controller.UserService.GetUser(req.UserName)
	if err != nil {
		writeJSONError(w, err)
		return
	}

	resp := dtos.GetUserResponse{
		Name:         user.Name,
		Following:    user.GetAllFollowingUsers(),
		Publications: user.Publications.GetContents(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

}

func (controller *UsersController) AddPublication(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSONError(w, domain.NewAppError("[CONTROLLER]", domain.ErrMethodNotAllowed, "method="+r.Method))
		return
	}

	var req dtos.CreatePublicationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, domain.NewAppError("[CONTROLLER]", domain.ErrInvalidRequestBody, ""))
		return
	}

	if req.UserName == "" {
		writeJSONError(w, domain.NewAppError("[CONTROLLER]", domain.ErrInvalidRequestBody, "user_id"))
		return
	}

	if req.Content == "" {
		writeJSONError(w, domain.NewAppError("[CONTROLLER]", domain.ErrInvalidRequestBody, "content"))
		return
	}

	if err := controller.UsecasePublishMessage.Execute(req.UserName, req.Content); err != nil {
		writeJSONError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Message published successfully", "status": http.StatusCreated})
}

func (controller *UsersController) AddFollowing(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSONError(w, domain.NewAppError("[CONTROLLER]", domain.ErrMethodNotAllowed, "method="+r.Method))
		return
	}

	var req dtos.FollowRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, domain.NewAppError("[CONTROLLER]", domain.ErrInvalidRequestBody, ""))
		return
	}

	if req.UserName == "" {
		writeJSONError(w, domain.NewAppError("[CONTROLLER]", domain.ErrInvalidRequestBody, "user_id"))
		return
	}

	if req.NewFollow == "" {
		writeJSONError(w, domain.NewAppError("[CONTROLLER]", domain.ErrInvalidRequestBody, "new_follow"))
		return
	}

	if err := controller.UsecaseFollow.Execute(req.UserName, req.NewFollow); err != nil {
		writeJSONError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "New Follower added successfully", "status": http.StatusCreated})
}

func (controller *UsersController) GetTimeline(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSONError(w, domain.NewAppError("[CONTROLLER]", domain.ErrMethodNotAllowed, "method="+r.Method))
		return
	}

	var req dtos.GetUserTimelineRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, domain.NewAppError("[CONTROLLER]", domain.ErrInvalidRequestBody, ""))
		return
	}

	feed, err := controller.UsecaseObtainUserTimeline.Execute(req.UserName)
	if err != nil {
		writeJSONError(w, err)
		return
	}

	messageResponse := make([]dtos.MessageResponse, 0)
	for _, msg := range feed.GetAllMessages() {
		messageResponse = append(messageResponse, dtos.MessageResponse{
			ID:        msg.ID(),
			UserName:  msg.UserName(),
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

func writeJSONError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")

	if appErr, ok := err.(*domain.AppError); ok {
		w.WriteHeader(appErr.Code)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error":    appErr.Message,
			"status":   appErr.Code,
			"resource": appErr.Resource,
			"op":       appErr.Op,
		})
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error":  "internal server error",
		"status": http.StatusInternalServerError,
	})
}
