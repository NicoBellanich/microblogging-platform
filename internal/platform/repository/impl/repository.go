package impl

import (
	"errors"

	"github.com/nicobellanich/migroblogging-platform/config"
	"github.com/nicobellanich/migroblogging-platform/internal/platform/repository"
	"github.com/nicobellanich/migroblogging-platform/internal/platform/repository/inmemory"
	"github.com/nicobellanich/migroblogging-platform/internal/platform/repository/prod"
	"github.com/nicobellanich/migroblogging-platform/internal/platform/repository/test"
)

func NewMessageRepository(conf *config.Config) (repository.IMessageRepository, error) {
	var mr repository.IMessageRepository

	switch {
	case conf.IsProdEnv():
		mr = prod.NewMessageRepository()
	case conf.IsTestEnv():
		mr = test.NewMessageRepository()
	case conf.IsLocalEnv():
		mr = inmemory.NewMessageRepository()
	default:
		return nil, errors.New("something went wrong loading enviroments for creating MessageRepository")
	}

	return mr, nil
}
