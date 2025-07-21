//go:generate mockgen -source=interfaces.go -destination=../../mocks/repository/repository_mock.go -package=mocks
package repository

import "github.com/nicobellanich/migroblogging-platform/internal/domain"

type IMessageRepository interface {
	Save(*domain.Message) error
	LoadAllByUser(string) ([]domain.Message, error)
}

type IFollowersRepository interface {
	Save(string, string) error
	LoadFollowersByUser(string) ([]string, error)
}
