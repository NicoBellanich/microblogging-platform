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

// NewMessage add new message with validations
func NewMessage(content, userID string) (*Message, error) {
	if len(content) == 0 {
		return nil, errors.New("content cannot be empty")
	}
	if len(content) > 280 {
		return nil, errors.New("content exceeds 280 characters")
	}
	if len(userID) == 0 {
		return nil, errors.New("userID cannot be empty")
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

// SetContent updates the message content with validation
func (m *Message) SetContent(content string) error {
	if len(content) == 0 {
		return errors.New("content cannot be empty")
	}
	if len(content) > 280 {
		return errors.New("content exceeds 280 characters")
	}
	m.content = content
	return nil
}

// SetCreatedAt updates the message creation time
func (m *Message) SetCreatedAt(createdAt time.Time) {
	m.createdAt = createdAt
}
