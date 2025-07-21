package prod

import "github.com/nicobellanich/migroblogging-platform/internal/platform/repository"

func NewMessageRepository() repository.IMessageRepository {
	return &MessageRepository{}
}

type MessageRepository struct{}

func (mr *MessageRepository) Save() {
	panic("implement") // this should be implemented in real prod code
}

func (mr *MessageRepository) Load() {
	panic("implement") // this should be implemented in real prod code
}

func (mr *MessageRepository) LoadAll() {
	panic("implement") // this should be implemented in real prod code
}
