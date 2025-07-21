package usecase_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	"github.com/nicobellanich/migroblogging-platform/internal/domain"
	mocks "github.com/nicobellanich/migroblogging-platform/internal/mocks/repository"
	"github.com/nicobellanich/migroblogging-platform/internal/usecase"
)

type FollowUseCaseTestSuite struct {
	suite.Suite
	ctrl              *gomock.Controller
	mockFollowersRepo *mocks.MockIFollowersRepository
	mockUsersRepo     *mocks.MockIUsersRepository

	usecase *usecase.Follow
}

func (suite *FollowUseCaseTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.mockUsersRepo = mocks.NewMockIUsersRepository(suite.ctrl)
	suite.usecase = usecase.NewFollow(suite.mockUsersRepo)
}

func (suite *FollowUseCaseTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *FollowUseCaseTestSuite) TestExecute_Success() {
	userID := "nicolas"
	newFollow := "maria"

	user := &domain.User{Name: userID}
	userToFollow := &domain.User{Name: newFollow}

	suite.mockUsersRepo.
		EXPECT().
		Get(userID).
		Return(user, nil)

	suite.mockUsersRepo.
		EXPECT().
		Get(newFollow).
		Return(userToFollow, nil)

	err := suite.usecase.Execute(userID, newFollow)
	suite.NoError(err)
	suite.Len(user.Following, 1)
	suite.Equal(newFollow, user.Following[0].Name)
}

func (suite *FollowUseCaseTestSuite) TestExecute_ErrorFromRepo() {
	userID := "nicolas"
	newFollow := "maria"
	expectedErr := domain.ErrInvalidArgument

	suite.mockUsersRepo.
		EXPECT().
		Get(userID).
		Return(nil, expectedErr)

	err := suite.usecase.Execute(userID, newFollow)
	suite.Equal(expectedErr, err)
}

func TestFollowUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(FollowUseCaseTestSuite))
}
