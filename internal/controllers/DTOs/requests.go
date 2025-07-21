package dtos

type PublishRequest struct {
	UserID  string `json:"user_id"`
	Content string `json:"content"`
}

type FollowRequest struct {
	UserID    string `json:"user_id"`
	NewFollow string `json:"new_follow"`
}

type GetUserTimelineRequest struct {
	UserID string `json:"user_id"`
}
