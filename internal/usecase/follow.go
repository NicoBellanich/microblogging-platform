package usecase

import (
	"fmt"

	"github.com/nicobellanich/migroblogging-platform/internal/domain"
	"github.com/nicobellanich/migroblogging-platform/internal/platform/repository"
)

type Follow struct {
	FollowersRepository repository.IFollowersRepository
}

func NewFollow(fr repository.IFollowersRepository) *Follow {
	return &Follow{
		FollowersRepository: fr,
	}
}

// Execute runs UseCase Follow
func (uc *Follow) Execute(userID string, newFollow string) error {

	err := uc.FollowersRepository.Save(userID, newFollow)
	if err != nil {
		if err == domain.ErrInvalidArgument {
			return err
		}
		return fmt.Errorf("failed to follow user: %w", err)
	}

	fmt.Printf("👤@%s , now is following  👤@%s \n", userID, newFollow)

	return nil
}
