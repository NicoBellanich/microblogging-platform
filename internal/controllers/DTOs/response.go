package dtos

import "github.com/nicobellanich/migroblogging-platform/internal/domain"

type MessageResponse struct {
	ID        string `json:"id"`
	UserName  string `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

type GetUserTimelineResponse struct {
	Feed []MessageResponse `json:"feed"`
}

type GetUserResponse struct {
	Name         string             `json:"username"`
	Following    []*domain.User     `json:"following"`
	Publications domain.MessageList `json:"publications"` // aca deberia devolver []MessageResponse tambien
}
