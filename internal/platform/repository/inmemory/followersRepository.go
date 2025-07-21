package inmemory

import (
	"sync"

	"github.com/nicobellanich/migroblogging-platform/internal/domain"
	"github.com/nicobellanich/migroblogging-platform/internal/platform/repository"
)

type FollowersRepository struct {
	followers map[string][]string
	mutex     sync.Mutex
}

func NewFollowersRepository() repository.IFollowersRepository {
	return &FollowersRepository{
		followers: make(map[string][]string),
	}
}

func (fr *FollowersRepository) Save(userID, followerID string) error {
	if userID == "" || followerID == "" {
		return domain.ErrInvalidArgument
	}

	fr.mutex.Lock()
	defer fr.mutex.Unlock()

	if _, exists := fr.followers[userID]; !exists {
		fr.followers[userID] = make([]string, 0)
	}

	// avoid duplicate following
	for _, id := range fr.followers[userID] {
		if id == followerID {
			return nil
		}
	}

	fr.followers[userID] = append(fr.followers[userID], followerID)
	return nil
}

func (fr *FollowersRepository) LoadFollowersByUser(userID string) ([]string, error) {
	if userID == "" {
		return nil, domain.ErrInvalidArgument
	}

	fr.mutex.Lock()
	defer fr.mutex.Unlock()

	followers, exists := fr.followers[userID]
	if !exists {
		return nil, domain.ErrNoFollowersForUser
	}

	// create copy to avoid external modifications
	followersCopy := make([]string, len(followers))
	copy(followersCopy, followers)
	return followersCopy, nil
}
