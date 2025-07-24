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
		return domain.NewAppError(
			"[REPOSITORY]",
			domain.ErrNilUserProvided,
			"",
		)
	}

	ur.mutex.Lock()
	defer ur.mutex.Unlock()

	if _, exists := ur.users[user.Name]; exists {
		return domain.NewAppError(
			"[REPOSITORY]",
			domain.ErrUserAlreadyExists,
			"username="+user.Name,
		)
	}

	ur.users[user.Name] = user
	return nil
}

func (ur *UsersRepository) Update(userName string, user *domain.User) error {
	if user == nil {
		return domain.NewAppError("REPOSITORY", domain.ErrNilUserProvided, "")
	}
	if userName == "" {
		return domain.NewAppError("REPOSITORY", domain.ErrUserNameEmpty, "")
	}

	ur.mutex.Lock()
	defer ur.mutex.Unlock()

	if _, exists := ur.users[userName]; !exists {
		return domain.NewAppError("REPOSITORY", domain.ErrUserNotFound, "user="+userName)
	}

	ur.users[userName] = user
	return nil
}

func (ur *UsersRepository) Get(userName string) (*domain.User, error) {
	if userName == "" {
		return nil, domain.NewAppError("[REPOSITORY]", domain.ErrUserNameEmpty, "")
	}

	ur.mutex.Lock()
	defer ur.mutex.Unlock()

	user, exists := ur.users[userName]
	if !exists {
		return nil, domain.NewAppError("[REPOSITORY]", domain.ErrUserNotFound, "username="+userName)
	}

	return user, nil
}
