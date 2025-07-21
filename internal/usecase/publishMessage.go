package usecase

import (
	"github.com/nicobellanich/migroblogging-platform/internal/domain"
	"github.com/nicobellanich/migroblogging-platform/internal/platform/repository"
)

type PublishMessage struct {
	MessageRepository repository.IMessageRepository
}

func NewPublishMessage(mr repository.IMessageRepository) *PublishMessage {
	return &PublishMessage{
		MessageRepository: mr,
	}
}

// Execute runs UseCase PublishMessage
func (pm *PublishMessage) Execute(userID string, content string) error {
	newMessage, err := domain.NewMessage(content, userID)
	if err != nil {
		return err
	}

	err = pm.MessageRepository.Save(newMessage)
	if err != nil {
		return err
	}

	return nil
}
