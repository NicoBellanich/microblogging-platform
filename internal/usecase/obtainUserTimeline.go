// Package usecase contains the business logic for the microblogging platform.
// This file defines the use case for obtaining a user's timeline.
package usecase

import (
	"fmt"

	"github.com/nicobellanich/migroblogging-platform/internal/domain"
	"github.com/nicobellanich/migroblogging-platform/internal/platform/repository"
)

// ObtainUserTimeline handles the logic for retrieving a user's timeline (feed).
type ObtainUserTimeline struct {
	FollowersRepository repository.IFollowersRepository
	MessageRepository   repository.IMessageRepository
}

// NewObtainUserTimeline creates a new ObtainUserTimeline use case with the given repositories.
func NewObtainUserTimeline(fr repository.IFollowersRepository, mr repository.IMessageRepository) *ObtainUserTimeline {
	return &ObtainUserTimeline{
		FollowersRepository: fr,
		MessageRepository:   mr,
	}
}

// Execute retrieves the timeline for the given user ID.
// It loads all users that the given user follows, fetches their messages, sorts them by time, and returns the contents.
func (uc *ObtainUserTimeline) Execute(userID string) ([]string, error) {

	var timeline []string

	// Load the list of users that userID follows
	followers, err := uc.FollowersRepository.LoadFollowersByUser(userID)
	if err != nil {
		if err == domain.ErrInvalidArgument || err == domain.ErrNoFollowersForUser {
			return nil, err
		}
		return nil, fmt.Errorf("failed to load followers: %w", err)
	}

	// For each followed user, load their messages
	var messages []domain.Message
	var messageList domain.MessageList
	for _, f := range followers {
		userMessages, err := uc.MessageRepository.LoadAllByUser(f)
		if err != nil {
			if err == domain.ErrInvalidArgument || err == domain.ErrNoMessagesForUser {
				return nil, err
			}
			return nil, fmt.Errorf("failed to load messages: %w", err)
		}
		messages = append(messages, userMessages...)
	}

	messageList = domain.MessageList(messages)

	// Sort messages by creation time (descending)
	messageList.SortByCreatedAtDescending()

	// Extract message contents for the timeline
	timeline = messageList.GetContents()

	// Optionally print the timeline to the console (for debugging/logging)
	consolePrintTimeline(userID, timeline)

	return timeline, nil
}

// consolePrintTimeline prints the timeline to the console for debugging/logging purposes.
func consolePrintTimeline(userID string, timeline []string) {
	fmt.Printf("👤@%s feed ========= \n", userID)
	for _, m := range timeline {
		fmt.Printf("💬 %s \n", m)
	}

	fmt.Println(" ==================================== \n ")
}
