// Package usecase contains the business logic for the microblogging platform.
// This file defines the use case for publishing a new message.
package usecase

import (
	"fmt"

	"github.com/nicobellanich/migroblogging-platform/internal/services"
)

// PublishMessage handles the logic for publishing a new message by a user.
type PublishMessage struct {
	UsersService services.IUserServices
}

// NewPublishMessage creates a new PublishMessage use case with the given repository.
func NewPublishMessage(us services.IUserServices) *PublishMessage {
	return &PublishMessage{
		UsersService: us,
	}
}

// Execute creates and saves a new message for the given user.
func (uc *PublishMessage) Execute(userName string, content string) error {

	err := uc.UsersService.AddPublication(userName, content)
	if err != nil {
		return err
	}

	fmt.Printf("👤@%s , just published - %s \n ", userName, content)

	return nil
}
