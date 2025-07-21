package usecase_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nicobellanich/migroblogging-platform/internal/domain"
	mocks "github.com/nicobellanich/migroblogging-platform/internal/mocks/services"
	"github.com/nicobellanich/migroblogging-platform/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func TestObtainUserTimeline_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserServices := mocks.NewMockIUserServices(ctrl)

	// Arrange
	userID := "nicolas"

	pub1, _ := domain.NewMessage("hola mundo", "followed1")
	pub2, _ := domain.NewMessage("otro mensaje", "followed1")

	followedUser := &domain.User{
		Name:         "followed1",
		Publications: []domain.Message{*pub1, *pub2},
	}

	mainUser := &domain.User{
		Name:      userID,
		Following: []*domain.User{followedUser},
	}

	mockUserServices.
		EXPECT().
		GetUser(userID).
		Return(mainUser, nil)

	useCase := usecase.NewObtainUserTimeline(mockUserServices)

	// Act
	feed, err := useCase.Execute(userID)

	// Assert
	assert.NoError(t, err)
	allMessages := feed.GetAllMessages()
	assert.Len(t, allMessages, 2)
	assert.Equal(t, "otro mensaje", allMessages[0].Content())
	assert.Equal(t, "hola mundo", allMessages[1].Content())
}
