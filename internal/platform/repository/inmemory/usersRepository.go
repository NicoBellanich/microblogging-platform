package inmemory

import (
	"sync"

	"github.com/nicobellanich/migroblogging-platform/internal/domain"
	"github.com/nicobellanich/migroblogging-platform/internal/platform/repository"
)

type UsersRepository struct {
	users map[string]*domain.User
	mutex sync.Mutex
}

func NewUsersRepository() repository.IUsersRepository {
	return &UsersRepository{
		users: make(map[string]*domain.User),
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

	ur.users[user.Name] = user
	return nil
}

func (ur *UsersRepository) Update(userName string, user *domain.User) error {
	if user == nil {
		return domain.ErrNilUserProvided
	}
	if userName == "" {
		return domain.ErrUserNameEmpty
	}

	ur.mutex.Lock()
	defer ur.mutex.Unlock()

	if _, exists := ur.users[userName]; !exists {
		return domain.ErrUserNotFound
	}

	ur.users[userName] = user
	return nil
}

func (ur *UsersRepository) Get(userName string) (*domain.User, error) {
	if userName == "" {
		return nil, domain.ErrUserNameEmpty
	}

	ur.mutex.Lock()
	defer ur.mutex.Unlock()

	user, exists := ur.users[userName]
	if !exists {
		return nil, domain.ErrUserNotFound
	}

	return user, nil
}
