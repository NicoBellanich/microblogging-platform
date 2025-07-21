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

	user := &domain.User{Name: userID, Following: []*domain.User{{Name: followingName1}, {Name: followingName2}}}

	msgJuan, _ := domain.NewMessage("Hola soy Juan", "juan")
	juan := &domain.User{Name: followingName2}
	juan.Publications.AddMessage(msgJuan)

	msgMaria, _ := domain.NewMessage("Hola soy Maria", "maria")
	maria := &domain.User{Name: followingName1}
	maria.Publications.AddMessage(msgMaria)

	suite.mockUsersRepo.
		EXPECT().
		Get(userID).
		Return(user, nil)

	suite.mockUsersRepo.
		EXPECT().
		Get(followingName1).
		Return(maria, nil)

	suite.mockUsersRepo.
		EXPECT().
		Get(followingName2).
		Return(juan, nil)

	timeline, err := suite.usecase.Execute(userID)

	suite.NoError(err)
	suite.Len(timeline, 2)
	suite.Contains(timeline[0], "Maria")
	suite.Contains(timeline[1], "Juan")
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

	user := &domain.User{Name: userID, Following: []*domain.User{{Name: followingName}}}

	suite.mockUsersRepo.
		EXPECT().
		Get(userID).
		Return(user, nil)

	suite.mockUsersRepo.
		EXPECT().
		Get(followingName).
		Return(nil, domain.ErrNoMessagesForUser)

	timeline, err := suite.usecase.Execute(userID)

	suite.Error(err)
	suite.Nil(timeline)
	suite.Equal(domain.ErrNoMessagesForUser, err)
}

func TestObtainUserTimelineTestSuite(t *testing.T) {
	suite.Run(t, new(ObtainUserTimelineTestSuite))
}
