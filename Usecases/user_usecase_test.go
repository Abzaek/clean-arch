package usecases

import (
	"testing"

	"github.com/Abzaek/clean-arch/Usecases/mocks"
	"github.com/Abzaek/clean-arch/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type userUsecaseSuite struct {
	suite.Suite
	repo     *mocks.UserService
	uUsecase *UserUsecase
}

func (suite *userUsecaseSuite) SetupTest() {
	repo := mocks.UserService{}
	suite.repo = &repo
	suite.uUsecase = NewUserUseCase(&repo)
}

func (suite *userUsecaseSuite) TestSaveUser() {
	// Mock data
	user := &domain.User{
		ID:       "1",
		Role:     "admin",
		Password: "1234",
	}

	// Set up expected behavior
	suite.repo.On("Save", user).Return(nil)

	// Execute the use case
	err := suite.uUsecase.SaveUser(user)

	// Verify expectations
	assert.NoError(suite.T(), err)

	// Ensure that the expectations were met
	suite.repo.AssertExpectations(suite.T())
}

func (suite *userUsecaseSuite) TestDeleteUser() {
	// Set up expected behavior
	suite.repo.On("Delete", "1").Return(nil)

	// Execute the use case
	err := suite.uUsecase.DeleteUser("1")

	// Verify expectations
	assert.NoError(suite.T(), err)

	// Ensure that the expectations were met
	suite.repo.AssertExpectations(suite.T())
}

func (suite *userUsecaseSuite) TestUpdateUser() {
	// Mock data
	user := &domain.User{
		ID:       "1",
		Role:     "admin",
		Password: "1234",
	}

	// Set up expected behavior
	suite.repo.On("Update", user).Return(nil)

	// Execute the use case
	err := suite.uUsecase.UpdateUser(user)

	// Verify expectations
	assert.NoError(suite.T(), err)

	// Ensure that the expectations were met
	suite.repo.AssertExpectations(suite.T())
}

func (suite *userUsecaseSuite) TestFindUser() {
	// Mock data
	user := &domain.User{
		ID:       "1",
		Role:     "admin",
		Password: "1234",
	}

	// Set up expected behavior
	suite.repo.On("Find", "1").Return(user, nil)

	// Execute the use case
	result, err := suite.uUsecase.FindUser("1")

	// Verify expectations
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), user, result)

	// Ensure that the expectations were met
	suite.repo.AssertExpectations(suite.T())
}

func TestUserUsecaseSuite(t *testing.T) {
	suite.Run(t, new(userUsecaseSuite))
}
