package usecase

import (
	"fmt"

	"github.com/nicobellanich/migroblogging-platform/internal/domain"
	"github.com/nicobellanich/migroblogging-platform/internal/platform/repository"
)

type ObtainUserTimeline struct {
	FollowersRepository repository.IFollowersRepository
	MessageRepository   repository.IMessageRepository
}

func NewObtainUserTimeline(fr repository.IFollowersRepository, mr repository.IMessageRepository) *ObtainUserTimeline {
	return &ObtainUserTimeline{
		FollowersRepository: fr,
		MessageRepository:   mr,
	}
}

// Execute runs UseCase Execute
func (uc *ObtainUserTimeline) Execute(userID string) ([]string, error) {

	var timeline []string

	followers, err := uc.FollowersRepository.LoadFollowersByUser(userID)
	if err != nil {
		return nil, err
	}

	// get messages of all the pople userID follows
	var messages []domain.Message
	var messageList domain.MessageList
	for _, f := range followers {
		userMessages, err := uc.MessageRepository.LoadAllByUser(f)
		if err != nil {
			return nil, err
		}
		messages = append(messages, userMessages...)
	}

	messageList = domain.MessageList(messages)

	// order messages by time
	messageList.SortByCreatedAtDescending()

	// get all messages content
	timeline = messageList.GetContents()

	consolePrintTimeline(userID, timeline)

	return timeline, nil
}

func consolePrintTimeline(userID string, timeline []string) {
	fmt.Printf("👤@%s feed ========= \n", userID)
	for _, m := range timeline {
		fmt.Printf("💬 %s \n", m)
	}

	fmt.Println(" ==================================== \n ")
}
