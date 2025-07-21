package services_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nicobellanich/migroblogging-platform/internal/domain"
	mocks "github.com/nicobellanich/migroblogging-platform/internal/mocks/repository"
	"github.com/nicobellanich/migroblogging-platform/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestAddUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIUsersRepository(ctrl)
	service := services.NewUserServices(mockRepo)

	username := "nico"
	mockRepo.EXPECT().Create(gomock.Any()).Return(nil)

	err := service.AddUser(username)

	assert.NoError(t, err)
}

func TestAddUser_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIUsersRepository(ctrl)
	service := services.NewUserServices(mockRepo)

	username := "nico"
	mockRepo.EXPECT().Create(gomock.Any()).Return(errors.New("create failed"))

	err := service.AddUser(username)

	assert.Error(t, err)
}

func TestGetUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIUsersRepository(ctrl)
	service := services.NewUserServices(mockRepo)

	username := "nico"
	expected := domain.NewUser(username)
	mockRepo.EXPECT().Get(username).Return(expected, nil)

	user, err := service.GetUser(username)

	assert.NoError(t, err)
	assert.Equal(t, expected, user)
}

func TestGetUser_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIUsersRepository(ctrl)
	service := services.NewUserServices(mockRepo)

	mockRepo.EXPECT().Get("nope").Return(nil, errors.New("not found"))

	user, err := service.GetUser("nope")

	assert.Nil(t, user)
	assert.Error(t, err)
}

func TestUpdateUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIUsersRepository(ctrl)
	service := services.NewUserServices(mockRepo)

	oldUser := domain.NewUser("nico")
	newUser := domain.NewUser("nuevo-nico")

	mockRepo.EXPECT().Get("nico").Return(oldUser, nil)
	mockRepo.EXPECT().Update("nico", gomock.Any()).Return(nil)

	err := service.UpdateUser("nico", newUser)

	assert.NoError(t, err)
}

func TestAddFollowing_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIUsersRepository(ctrl)
	service := services.NewUserServices(mockRepo)

	user := domain.NewUser("nico")
	target := domain.NewUser("mateo")

	mockRepo.EXPECT().Get("nico").Return(user, nil)
	mockRepo.EXPECT().Get("mateo").Return(target, nil)
	mockRepo.EXPECT().Update("nico", gomock.Any()).Return(nil)

	err := service.AddFollowing("nico", "mateo")

	assert.NoError(t, err)
}

func TestAddPublication_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIUsersRepository(ctrl)
	service := services.NewUserServices(mockRepo)

	user := domain.NewUser("nico")
	mockRepo.EXPECT().Get("nico").Return(user, nil)
	mockRepo.EXPECT().Update("nico", gomock.Any()).Return(nil)

	err := service.AddPublication("nico", "¡Hola mundo!")

	assert.NoError(t, err)
}

func TestAddPublication_InvalidMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIUsersRepository(ctrl)
	service := services.NewUserServices(mockRepo)

	err := service.AddPublication("nico", "")

	assert.ErrorIs(t, err, domain.ErrContentEmpty)
}
