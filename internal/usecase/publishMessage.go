// Package usecase contains the business logic for the microblogging platform.
// This file defines the use case for publishing a new message.
package usecase

import (
	"fmt"

	"github.com/nicobellanich/migroblogging-platform/internal/domain"
	"github.com/nicobellanich/migroblogging-platform/internal/platform/repository"
)

// PublishMessage handles the logic for publishing a new message by a user.
type PublishMessage struct {
	MessageRepository repository.IMessageRepository
}

// NewPublishMessage creates a new PublishMessage use case with the given repository.
func NewPublishMessage(mr repository.IMessageRepository) *PublishMessage {
	return &PublishMessage{
		MessageRepository: mr,
	}
}

// Execute creates and saves a new message for the given user.
// Returns an error if validation fails or saving fails.
func (uc *PublishMessage) Execute(userID string, content string) error {
	// Validate and create the message domain object
	newMessage, err := domain.NewMessage(content, userID)
	if err != nil {
		return err
	}

	// Save the message using the repository
	err = uc.MessageRepository.Save(newMessage)
	if err != nil {
		return err
	}

	// Log the publishing event
	fmt.Printf("👤@%s , just published - %s \n ", newMessage.UserID(), newMessage.Content())

	return nil
}
