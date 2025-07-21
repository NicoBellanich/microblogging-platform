//go:generate mockgen -source=interfaces.go -destination=../mocks/services/userServices_mock.go -package=mocks

package services

import "github.com/nicobellanich/migroblogging-platform/internal/domain"

type IUserServices interface {
	AddUser(userName string) error
	GetUser(userName string) (*domain.User, error)
	UpdateUser(userName string, newUser *domain.User) error
	AddFollowing(userName string, newFollowing string) error
	AddPublication(userName, content string) error
}
