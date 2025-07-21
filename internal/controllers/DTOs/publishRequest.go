package dtos

type PublishRequest struct {
	UserID  string `json:"user_id"`
	Content string `json:"content"`
}
