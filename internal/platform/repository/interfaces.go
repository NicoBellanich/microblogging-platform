//go:generate mockgen -source=interfaces.go -destination=../../mocks/repository/repository_mock.go -package=mocks
package repository

import "github.com/nicobellanich/migroblogging-platform/internal/domain"

type IUsersRepository interface {
	Create(*domain.User) error
	Update(string, *domain.User) error
	Get(string) (*domain.User, error)
}
