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
	UsersRepository repository.IUsersRepository
}

// NewObtainUserTimeline creates a new ObtainUserTimeline use case with the given repositories.
func NewObtainUserTimeline(ur repository.IUsersRepository) *ObtainUserTimeline {
	return &ObtainUserTimeline{
		UsersRepository: ur,
	}
}

// Execute retrieves the timeline for the given user ID.
// It loads all users that the given user follows, fetches their messages, sorts them by time, and returns the contents.
func (uc *ObtainUserTimeline) Execute(userID string) ([]string, error) {

	var timeline []string

	// get user
	usr, err := uc.UsersRepository.Get(userID)
	if err != nil {
		return nil, err
	}

	// build feed
	var userFeed domain.Feed
	for _, following := range usr.Following {
		userFeed.AddMessageList(&following.Publications)
	}

	userFeed.SortAllMessagesDescending()

	timeline = userFeed.GetMessagesContent()

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
