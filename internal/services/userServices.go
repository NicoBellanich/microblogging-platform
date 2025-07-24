package services

import (
	"fmt"

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

func (s *UserServices) UpdateUser(userName string, newUser *domain.User) error {

	user, err := s.GetUser(userName)
	if err != nil {
		return err
	}

	user.Name = newUser.Name
	user.Following = newUser.Following
	user.Publications = newUser.Publications

	err = s.UsersRepository.Update(userName, user)
	if err != nil {
		return err
	}

	return nil

}

func (s *UserServices) AddFollowing(userName, newFollowing string) error {

	// get user
	usr, err := s.UsersRepository.Get(userName)
	if err != nil {
		return err
	}

	for _, userFollowing := range usr.Following {
		if newFollowing == userFollowing.Name {
			return domain.NewAppError(
				"[SERVICE]",
				domain.ErrUserAlreadyFollowing,
				fmt.Sprintf("user=%s already follows user=%s", userName, newFollowing))
		}
	}

	// get new following
	usrNewFollow, err := s.UsersRepository.Get(newFollowing)
	if err != nil {
		return err
	}

	// append new following to user
	usr.AddFollowing(usrNewFollow)

	// update repository
	err = s.UsersRepository.Update(usr.Name, usr)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserServices) AddPublication(userName, content string) error {

	newMessage, err := domain.NewMessage(content, userName)
	if err != nil {
		return err
	}

	usr, err := s.UsersRepository.Get(userName)
	if err != nil {
		return err
	}

	usr.AddPublication(*newMessage)

	err = s.UsersRepository.Update(usr.Name, usr)
	if err != nil {
		return err
	}

	return nil
}
