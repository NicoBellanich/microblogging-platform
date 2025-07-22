package usecase

import (
	"fmt"

	"github.com/nicobellanich/migroblogging-platform/internal/services"
)

type Follow struct {
	UsersService services.IUserServices
}

func NewFollow(us services.IUserServices) *Follow {
	return &Follow{
		UsersService: us,
	}
}

// Execute runs UseCase Follow
func (usecase *Follow) Execute(userName string, newFollow string) error {

	err := usecase.UsersService.AddFollowing(userName, newFollow)
	if err != nil {
		return err
	}

	fmt.Printf("👤@%s , now is following  👤@%s \n", userName, newFollow)

	return nil
}
