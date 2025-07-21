package dtos

type FollowRequest struct {
	UserID    string `json:"user_id"`
	NewFollow string `json:"new_follow"`
}

type GetUserTimelineRequest struct {
	UserID string `json:"user_id"`
}

type CreateUserRequest struct {
	UserID string `json:"user_id"`
}

type GetUserRequest struct {
	UserID string `json:"user_id"`
}

type CreatePublicationRequest struct {
	UserID  string `json:"user_id"`
	Content string `json:"content"`
}
