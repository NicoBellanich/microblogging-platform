package impl

import (
	"github.com/nicobellanich/migroblogging-platform/config"
	"github.com/nicobellanich/migroblogging-platform/internal/platform/repository"
	"github.com/nicobellanich/migroblogging-platform/internal/platform/repository/inmemory"
	"github.com/nicobellanich/migroblogging-platform/internal/platform/repository/prod"
	"github.com/nicobellanich/migroblogging-platform/internal/platform/repository/test"
)

func NewMessageRepository(conf *config.Config) repository.IMessageRepository {
	var mr repository.IMessageRepository

	if conf.IsProdEnv() {
		mr = prod.NewMessageRepository()
	}

	if conf.IsTestEnv() {
		mr = test.NewMessageRepository()
	}

	if conf.IsLocalEnv() {
		mr = inmemory.NewMessageRepository()
	}

	return mr
}
