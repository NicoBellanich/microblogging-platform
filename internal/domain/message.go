package domain

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	iD        string    `json:"id"`
	content   string    `json:"content"`
	userName  string    `json:"user_id"`
	createdAt time.Time `json:"created_at"`
}

// NewMessage add new message with validations
func NewMessage(content, userName string) (*Message, error) {
	if len(content) == 0 {
		return nil, NewAppError("[DOMAIN]", ErrContentEmpty, "")
	}
	if len(content) > 280 {
		return nil, NewAppError("[DOMAIN]", ErrContentTooLong, "")
	}
	if len(userName) == 0 {
		return nil, NewAppError("[DOMAIN]", ErrUserNameEmpty, "")
	}

	return &Message{
		iD:        uuid.New().String(),
		content:   content,
		userName:  userName,
		createdAt: time.Now().UTC(),
	}, nil
}

// ID returns Message ID
func (m *Message) ID() string {
	return m.iD
}

// Content returns Message content
func (m *Message) Content() string {
	return m.content
}

// UserName returns Message user name
func (m *Message) UserName() string {
	return m.userName
}

// CreatedAt returns Message created time.Time
func (m *Message) CreatedAt() time.Time {
	return m.createdAt
}
