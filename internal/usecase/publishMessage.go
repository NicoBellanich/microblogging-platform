package usecase

import (
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
func (pm *PublishMessage) Execute() {
	panic("implement")
}
