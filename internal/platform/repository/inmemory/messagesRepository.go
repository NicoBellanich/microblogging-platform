package inmemory

import "github.com/nicobellanich/migroblogging-platform/internal/repository"

func NewMessageRepository() repository.IMessageRepository {
	return &MessageRepository{}
}

type MessageRepository struct{}

func (mr *MessageRepository) Save() {
	panic("implement")
}

func (mr *MessageRepository) Load() {
	panic("implement")
}

func (mr *MessageRepository) LoadAll() {
	panic("implement")
}
