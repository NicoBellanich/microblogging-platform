package impl

import (
	"errors"

	"github.com/nicobellanich/migroblogging-platform/config"
	"github.com/nicobellanich/migroblogging-platform/internal/platform/repository"
	"github.com/nicobellanich/migroblogging-platform/internal/platform/repository/inmemory"
	"github.com/nicobellanich/migroblogging-platform/internal/platform/repository/prod"
	"github.com/nicobellanich/migroblogging-platform/internal/platform/repository/test"
)

func NewUsersRepository(conf *config.Config) (repository.IUsersRepository, error) {
	var ur repository.IUsersRepository

	switch {
	case conf.IsProdEnv():
		ur = prod.NewUsersRepository()
	case conf.IsTestEnv():
		ur = test.NewUsersRepository()
	case conf.IsLocalEnv():
		ur = inmemory.NewUsersRepository()
	default:
		return nil, errors.New("something went wrong loading environments for creating MessageRepository")
	}

	return ur, nil
}
