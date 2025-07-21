package prod

import (
	"github.com/nicobellanich/migroblogging-platform/internal/domain"
	"github.com/nicobellanich/migroblogging-platform/internal/platform/repository"
)

func NewMessageRepository() repository.IMessageRepository {
	return &MessageRepository{}
}

type MessageRepository struct{}

func (mr *MessageRepository) Save(msg *domain.Message) error {
	panic("implement") // this should be implemented in real prod code
}

func (mr *MessageRepository) LoadAllByUser(userID string) ([]domain.Message, error) {
	panic("implement") // this should be implemented in real prod code
}
