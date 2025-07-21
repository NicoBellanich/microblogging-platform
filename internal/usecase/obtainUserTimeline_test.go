package usecase_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	"github.com/nicobellanich/migroblogging-platform/internal/domain"
	mocks "github.com/nicobellanich/migroblogging-platform/internal/mocks/repository"
	"github.com/nicobellanich/migroblogging-platform/internal/usecase"
)

type ObtainUserTimelineTestSuite struct {
	suite.Suite
	ctrl              *gomock.Controller
	mockFollowersRepo *mocks.MockIFollowersRepository
	mockMessageRepo   *mocks.MockIMessageRepository
	usecase           *usecase.ObtainUserTimeline
}

func (suite *ObtainUserTimelineTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.mockFollowersRepo = mocks.NewMockIFollowersRepository(suite.ctrl)
	suite.mockMessageRepo = mocks.NewMockIMessageRepository(suite.ctrl)
	suite.usecase = usecase.NewObtainUserTimeline(suite.mockFollowersRepo, suite.mockMessageRepo)
}

func (suite *ObtainUserTimelineTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *ObtainUserTimelineTestSuite) TestExecute_Success() {
	userID := "nicolas"
	following := []string{"maria", "juan"}

	jm, _ := domain.NewMessage("Hola soy Juan", "juan")

	messagesJuan := []domain.Message{*jm}

	time.Sleep(time.Microsecond * 100)

	mm, _ := domain.NewMessage("Hola soy Maria", "maria")

	messagesMaria := []domain.Message{*mm}

	suite.mockFollowersRepo.
		EXPECT().
		LoadFollowersByUser(userID).
		Return(following, nil)

	suite.mockMessageRepo.
		EXPECT().
		LoadAllByUser("maria").
		Return(messagesMaria, nil)

	suite.mockMessageRepo.
		EXPECT().
		LoadAllByUser("juan").
		Return(messagesJuan, nil)

	timeline, err := suite.usecase.Execute(userID)

	suite.NoError(err)
	suite.Len(timeline, 2)
	suite.Equal("Hola soy Maria- said @maria", timeline[0])
	suite.Equal("Hola soy Juan- said @juan", timeline[1])
}

func (suite *ObtainUserTimelineTestSuite) TestExecute_FollowersRepoError() {
	userID := "nicolas"
	expectedErr := fmt.Errorf("failed to load followers")

	suite.mockFollowersRepo.
		EXPECT().
		LoadFollowersByUser(userID).
		Return(nil, expectedErr)

	timeline, err := suite.usecase.Execute(userID)

	suite.Error(err)
	suite.Nil(timeline)
	suite.EqualError(err, expectedErr.Error())
}

func (suite *ObtainUserTimelineTestSuite) TestExecute_MessageRepoError() {
	userID := "nicolas"
	following := []string{"maria"}
	expectedErr := fmt.Errorf("failed to load messages")

	suite.mockFollowersRepo.
		EXPECT().
		LoadFollowersByUser(userID).
		Return(following, nil)

	suite.mockMessageRepo.
		EXPECT().
		LoadAllByUser("maria").
		Return(nil, expectedErr)

	timeline, err := suite.usecase.Execute(userID)

	suite.Error(err)
	suite.Nil(timeline)
	suite.EqualError(err, expectedErr.Error())
}

func TestObtainUserTimelineTestSuite(t *testing.T) {
	suite.Run(t, new(ObtainUserTimelineTestSuite))
}
