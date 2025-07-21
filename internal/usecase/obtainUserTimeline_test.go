package usecase_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	"github.com/nicobellanich/migroblogging-platform/internal/domain"
	mocks "github.com/nicobellanich/migroblogging-platform/internal/mocks/repository"
	"github.com/nicobellanich/migroblogging-platform/internal/usecase"
)

type ObtainUserTimelineTestSuite struct {
	suite.Suite
	ctrl              *gomock.Controller
	mockUsersRepo     *mocks.MockIUsersRepository
	mockFollowersRepo *mocks.MockIFollowersRepository
	mockMessageRepo   *mocks.MockIMessageRepository
	usecase           *usecase.ObtainUserTimeline
}

func (suite *ObtainUserTimelineTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.mockUsersRepo = mocks.NewMockIUsersRepository(suite.ctrl)
	suite.usecase = usecase.NewObtainUserTimeline(suite.mockUsersRepo)
}

func (suite *ObtainUserTimelineTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *ObtainUserTimelineTestSuite) TestExecute_Success() {
	userID := "nicolas"
	followingName1 := "maria"
	followingName2 := "juan"

	msgJuan, _ := domain.NewMessage("Hola soy Juan", "juan")
	juan := &domain.User{Name: followingName2}
	juan.Publications.AddMessage(msgJuan)

	msgMaria, _ := domain.NewMessage("Hola soy Maria", "maria")
	maria := &domain.User{Name: followingName1}
	maria.Publications.AddMessage(msgMaria)

	user := &domain.User{Name: userID, Following: []*domain.User{maria, juan}}

	suite.mockUsersRepo.
		EXPECT().
		Get(userID).
		Return(user, nil)

	timeline, err := suite.usecase.Execute(userID)

	suite.NoError(err)
	suite.Len(timeline, 2)

	// Check that the correct messages are present
	var foundMaria, foundJuan bool
	for _, msg := range timeline {
		if msg.UserID() == "maria" && msg.Content() == "Hola soy Maria" {
			foundMaria = true
		}
		if msg.UserID() == "juan" && msg.Content() == "Hola soy Juan" {
			foundJuan = true
		}
	}
	suite.True(foundMaria, "Maria's message was not found")
	suite.True(foundJuan, "Juan's message was not found")
}

func (suite *ObtainUserTimelineTestSuite) TestExecute_FollowersRepoError() {
	userID := "nicolas"
	expectedErr := domain.ErrNoFollowersForUser

	suite.mockUsersRepo.
		EXPECT().
		Get(userID).
		Return(nil, expectedErr)

	timeline, err := suite.usecase.Execute(userID)

	suite.Error(err)
	suite.Nil(timeline)
	suite.Equal(expectedErr, err)
}

func (suite *ObtainUserTimelineTestSuite) TestExecute_MessageRepoError() {
	userID := "nicolas"
	followingName := "maria"

	maria := &domain.User{Name: followingName} // No publications
	user := &domain.User{Name: userID, Following: []*domain.User{maria}}

	suite.mockUsersRepo.
		EXPECT().
		Get(userID).
		Return(user, nil)

	timeline, err := suite.usecase.Execute(userID)

	suite.NoError(err)
	suite.Len(timeline, 0)
}

func TestObtainUserTimelineTestSuite(t *testing.T) {
	suite.Run(t, new(ObtainUserTimelineTestSuite))
}
