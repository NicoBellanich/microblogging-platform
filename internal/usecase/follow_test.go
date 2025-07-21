package usecase_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	mocks "github.com/nicobellanich/migroblogging-platform/internal/mocks/repository"
	"github.com/nicobellanich/migroblogging-platform/internal/usecase"
)

type FollowUseCaseTestSuite struct {
	suite.Suite
	ctrl              *gomock.Controller
	mockFollowersRepo *mocks.MockIFollowersRepository
	usecase           *usecase.Follow
}

func (suite *FollowUseCaseTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.mockFollowersRepo = mocks.NewMockIFollowersRepository(suite.ctrl)
	suite.usecase = usecase.NewFollow(suite.mockFollowersRepo)
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
	expectedErr := fmt.Errorf("could not save follow")

	suite.mockFollowersRepo.
		EXPECT().
		Save(userID, newFollow).
		Return(expectedErr)

	err := suite.usecase.Execute(userID, newFollow)
	suite.EqualError(err, expectedErr.Error())
}

func TestFollowUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(FollowUseCaseTestSuite))
}
