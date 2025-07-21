package services

import (
	"github.com/nicobellanich/migroblogging-platform/internal/domain"
	"github.com/nicobellanich/migroblogging-platform/internal/platform/repository"
)

type UserServices struct {
	UsersRepository repository.IUsersRepository
}

func NewUserServices(ur repository.IUsersRepository) IUserServices {
	return &UserServices{
		UsersRepository: ur,
	}
}

func (s *UserServices) AddUser(userName string) error {

	usr := domain.NewUser(userName)

	err := s.UsersRepository.Create(usr)
	if err != nil {
		return err
	}

	return nil

}

func (s *UserServices) GetUser(userName string) (*domain.User, error) {
	user, err := s.UsersRepository.Get(userName)

	if err != nil {
		return nil, err
	}

	return user, nil
}
