package inmemory

import (
	"sync"

	"github.com/nicobellanich/migroblogging-platform/internal/domain"
	"github.com/nicobellanich/migroblogging-platform/internal/platform/repository"
)

func NewMessageRepository() repository.IMessageRepository {
	return &MessageRepository{
		messages: make([]domain.Message, 0),
	}
}

type MessageRepository struct {
	messages []domain.Message
	mutex    sync.Mutex
}

func (mr *MessageRepository) Save(msg *domain.Message) error {
	if msg == nil {
		return domain.ErrInvalidArgument
	}

	mr.mutex.Lock()
	defer mr.mutex.Unlock()

	newMsg := *msg
	mr.messages = append(mr.messages, newMsg)
	return nil
}

func (mr *MessageRepository) LoadAllByUser(userID string) ([]domain.Message, error) {
	if userID == "" {
		return nil, domain.ErrInvalidArgument
	}

	mr.mutex.Lock()
	defer mr.mutex.Unlock()

	var userMessages []domain.Message
	for _, msg := range mr.messages {
		if msg.UserID() == userID {
			msgCopy := msg
			userMessages = append(userMessages, msgCopy)
		}
	}

	if len(userMessages) == 0 {
		return nil, domain.ErrNoMessagesForUser
	}

	return userMessages, nil
}
