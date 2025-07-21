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
	suite.mockFollowersRepo = mocks.NewMockIFollowersRepository(suite.ctrl)
	suite.usecase = usecase.NewFollow(suite.mockUsersRepo)
}

func (suite *FollowUseCaseTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *FollowUseCaseTestSuite) TestExecute_Success() {
	userID := "nicolas"
	newFollow := "maria"

	suite.mockFollowersRepo.
		EXPECT().
		Save(userID, newFollow).
		Return(nil)

	err := suite.usecase.Execute(userID, newFollow)
	suite.NoError(err)
}

func (suite *FollowUseCaseTestSuite) TestExecute_ErrorFromRepo() {
	userID := "nicolas"
	newFollow := "maria"
	expectedErr := domain.ErrInvalidArgument

	suite.mockFollowersRepo.
		EXPECT().
		Save(userID, newFollow).
		Return(expectedErr)

	err := suite.usecase.Execute(userID, newFollow)
	suite.Equal(expectedErr, err)
}

func TestFollowUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(FollowUseCaseTestSuite))
}
