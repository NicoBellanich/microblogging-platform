package services

import "github.com/nicobellanich/migroblogging-platform/internal/domain"

type IUserServices interface {
	AddUser(userName string) error
	GetUser(userName string) (*domain.User, error)
}
