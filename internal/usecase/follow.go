package usecase

import "github.com/nicobellanich/migroblogging-platform/internal/platform/repository"

type Follow struct {
	FollowersRepository repository.IFollowersRepository
}

func NewFollow(fr repository.IFollowersRepository) *Follow {
	return &Follow{
		FollowersRepository: fr,
	}
}

// Execute runs UseCase Follow
func (f *Follow) Execute(userID string, newFollow string) error {

	err := f.FollowersRepository.Save(userID, newFollow)
	if err != nil {
		return err
	}

	return nil
}
