// Package usecase contains the business logic for the microblogging platform.
// This file defines the use case for obtaining a user's timeline.
package usecase

import (
	"fmt"

	"github.com/nicobellanich/migroblogging-platform/internal/domain"
	"github.com/nicobellanich/migroblogging-platform/internal/services"
)

// ObtainUserTimeline handles the logic for retrieving a user's timeline (feed).
type ObtainUserTimeline struct {
	UserServices services.IUserServices
}

// NewObtainUserTimeline creates a new ObtainUserTimeline use case with the given repositories.
func NewObtainUserTimeline(us services.IUserServices) *ObtainUserTimeline {
	return &ObtainUserTimeline{
		UserServices: us,
	}
}

// Execute retrieves the timeline for the given user ID.
// It loads all users that the given user follows, fetches their messages, sorts them by time, and returns the messages.
func (usecase *ObtainUserTimeline) Execute(userID string) (domain.Feed, error) {

	user, err := usecase.UserServices.GetUser(userID)
	if err != nil {
		return nil, err
	}

	var userFeed domain.Feed
	for _, following := range user.Following {
		userFeed.AddMessageList(&following.Publications)
	}

	userFeed.SortAllMessagesDescending()

	consolePrintTimeline(userID, userFeed.GetAllMessages())

	return userFeed, nil
}

// consolePrintTimeline prints the timeline to the console for debugging/logging purposes.
func consolePrintTimeline(user string, messages []domain.Message) {
	fmt.Printf("👤@%s feed ========= \n", user)
	for _, m := range messages {
		fmt.Printf("💬 %s - by %s \n", m.Content(), m.UserID())
	}

	fmt.Println(" ==================================== \n ")
}
