package impl

import (
	"errors"

	"github.com/nicobellanich/migroblogging-platform/config"
	"github.com/nicobellanich/migroblogging-platform/internal/platform/repository"
	"github.com/nicobellanich/migroblogging-platform/internal/platform/repository/inmemory"
	"github.com/nicobellanich/migroblogging-platform/internal/platform/repository/prod"
	"github.com/nicobellanich/migroblogging-platform/internal/platform/repository/test"
)

func NewFollowersRepository(conf *config.Config) (repository.IFollowersRepository, error) {
	var fr repository.IFollowersRepository

	switch {
	case conf.IsProdEnv():
		fr = prod.NewFollowersRepository()
	case conf.IsTestEnv():
		fr = test.NewFollowersRepository()
	case conf.IsLocalEnv():
		fr = inmemory.NewFollowersRepository()
	default:
		return nil, errors.New("something went wrong loading enviroments for creating MessageRepository")
	}

	return fr, nil
}
