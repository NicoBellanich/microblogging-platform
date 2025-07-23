package domain

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	Username  string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

// NewMessage add new message with validations
func NewMessage(content, userName string) (*Message, error) {
	if len(content) == 0 {
		return nil, ErrContentEmpty
	}
	if len(content) > 280 {
		return nil, ErrContentTooLong
	}
	if len(userName) == 0 {
		return nil, ErrUserNameEmpty
	}

	return &Message{
		ID:        uuid.New().String(),
		Content:   content,
		Username:  userName,
		CreatedAt: time.Now().UTC(),
	}, nil
}
