package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Message struct {
	iD        string    `json:"id"`
	content   string    `json:"content"`
	userID    string    `json:"user_id"`
	createdAt time.Time `json:"created_at"`
}

var (
	ErrContentEmpty       = errors.New("content cannot be empty")
	ErrContentTooLong     = errors.New("content exceeds 280 characters")
	ErrUserIDEmpty        = errors.New("userID cannot be empty")
	ErrInvalidArgument    = errors.New("invalid argument")
	ErrNoMessagesForUser  = errors.New("user doesn't have any post yet")
	ErrNoFollowersForUser = errors.New("user doesn't have any followers yet")
)

// NewMessage add new message with validations
func NewMessage(content, userID string) (*Message, error) {
	if len(content) == 0 {
		return nil, ErrContentEmpty
	}
	if len(content) > 280 {
		return nil, ErrContentTooLong
	}
	if len(userID) == 0 {
		return nil, ErrUserIDEmpty
	}

	return &Message{
		iD:        uuid.New().String(),
		content:   content,
		userID:    userID,
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

// UserID returns Message user ID
func (m *Message) UserID() string {
	return m.userID
}

// CreatedAt returns Message created time.Time
func (m *Message) CreatedAt() time.Time {
	return m.createdAt
}
