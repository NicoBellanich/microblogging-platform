package usecase

import (
	"fmt"

	"github.com/nicobellanich/migroblogging-platform/internal/platform/repository"
)

type Follow struct {
	UsersRepository repository.IUsersRepository
}

func NewFollow(ur repository.IUsersRepository) *Follow {
	return &Follow{
		UsersRepository: ur,
	}
}

// Execute runs UseCase Follow
func (uc *Follow) Execute(userID string, newFollow string) error {

	usr, err := uc.UsersRepository.Get(userID)
	if err != nil {
		return err
	}

	usrNewFollow, err := uc.UsersRepository.Get(newFollow)
	if err != nil {
		return err
	}

	usr.AddFollowing(usrNewFollow)

	fmt.Printf("👤@%s , now is following  👤@%s \n", userID, newFollow)

	return nil
}
