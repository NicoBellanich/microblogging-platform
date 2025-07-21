package inmemory

import (
	"sync"

	"github.com/nicobellanich/migroblogging-platform/internal/domain"
	"github.com/nicobellanich/migroblogging-platform/internal/platform/repository"
)

type UsersRepository struct {
	users map[string]domain.User
	mutex sync.Mutex
}

func NewUsersRepository() repository.IUsersRepository {
	return &UsersRepository{
		users: make(map[string]domain.User),
	}
}

func (ur *UsersRepository) Create(user *domain.User) error {
	if user == nil {
		return domain.ErrNilUserProvided
	}

	ur.mutex.Lock()
	defer ur.mutex.Unlock()

	if _, exists := ur.users[user.Name]; exists {
		return domain.ErrUserAlreadyExists
	}

	userCopy := *user
	ur.users[user.Name] = userCopy
	return nil
}

func (ur *UsersRepository) Update(userID string, user *domain.User) error {
	if user == nil {
		return domain.ErrNilUserProvided
	}
	if userID == "" {
		return domain.ErrUserIDEmpty
	}

	ur.mutex.Lock()
	defer ur.mutex.Unlock()

	if _, exists := ur.users[userID]; !exists {
		return domain.ErrUserNotFound
	}

	userCopy := *user
	ur.users[userID] = userCopy
	return nil
}

func (ur *UsersRepository) Get(userID string) (*domain.User, error) {
	if userID == "" {
		return nil, domain.ErrUserIDEmpty
	}

	ur.mutex.Lock()
	defer ur.mutex.Unlock()

	user, exists := ur.users[userID]
	if !exists {
		return nil, domain.ErrUserNotFound
	}

	userCopy := user
	return &userCopy, nil
}
