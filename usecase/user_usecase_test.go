package usecase

import (
	"errors"
	"testing"
	"time"

	repomock "4crypto/mock/repo_mock"
	usecasemock "4crypto/mock/usecase_mock"
	"4crypto/model/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// MockUserRepository is a mock implementation of UserRepository interface
type MockUserRepository struct {
	mock.Mock
}

var (
	mockuser = entity.User{
		Id:        "1",
		Name:      "Redo",
		Email:     "redo@example.com",
		Username:  "Redo",
		Password:  "redo1234",
		Role:      "admin",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
)

type UserUseCaseTestSuite struct {
	suite.Suite
	ucm *usecasemock.UserUseCaseMock
	urm *repomock.UserRepoMock
	uc  UserUseCase
}

func (s *UserUseCaseTestSuite) SetupTest() {
	suite.ucm = new(usecasemock.UserUseCaseMock)
	suite.urm = new(repomock.UserRepoMock)
	suite.uc = NewUserUseCase(suite.urm)
}

func (suite *UserUseCaseTestSuite) TestFindById_Success() {
	suite.urm.On("GetById", mockuser.Id).Return(mockuser, nil)
	actualUser, actualErr := suite.uc.FindById(mockuser.Id)
	assert.Nil(suite.T(), actualErr)
	assert.Equal(suite.T(), mockuser, actualUser)
}

func (suite *UserUseCaseTestSuite) TestFindById_UserFail() {
	suite.urm.On("GetById", mockuser.Id).Return(entity.User{}, errors.New("user not found"))
	_, actualErr := suite.uc.FindById(mockuser.Id)
	assert.Error(suite.T(), actualErr)
	assert.Equal(suite.T(), "user with id 1 not found", actualErr.Error())
}

func (suite *UserUseCaseTestSuite) TestFindByUsernamePassword_Success() {
	suite.urm.On("GetByUsername", mockuser.Username).Return(mockuser, nil)
	_, actualErr := suite.uc.FindByUsernamePassword(mockuser.Username, mockuser.Password)
	assert.Nil(suite.T(), actualErr)
}

func (suite *UserUseCaseTestSuite) TestFindByUsernamePassword_InvalidUsername() {
	suite.urm.On("GetByUsername", mockuser.Username).Return(entity.User{}, errors.New("user not found"))
	_, actualErr := suite.uc.FindByUsernamePassword(mockuser.Username, mockuser.Password)
	assert.Error(suite.T(), actualErr)
	assert.Equal(suite.T(), "invalid username or password", actualErr.Error())
}

func (suite *UserUseCaseTestSuite) TestFindUsernamePassword_InvalidPassword() {
	// Prepare the mock to return a user with a different password
	mockUserWithDifferentPassword := mockuser
	mockUserWithDifferentPassword.Password = "differentpassword"
	suite.urm.On("GetByUsername", mockuser.Username).Return(mockUserWithDifferentPassword, nil)

	// Call the FindByUsernamePassword function
	_, actualErr := suite.uc.FindByUsernamePassword(mockuser.Username, mockuser.Password)

	// Assert that error is not nil and contains the expected error message
	assert.Error(suite.T(), actualErr)
	assert.Equal(suite.T(), "invalid username or password", actualErr.Error())
}

func (suite *UserUseCaseTestSuite) TestFindUsernamePassword_Success_PasswordEmpty() {
	mockuser.Password = "redo1234"
	suite.urm.On("GetByUsername", mockuser.Username).Return(mockuser, nil)
	user, actualErr := suite.uc.FindByUsernamePassword(mockuser.Username, mockuser.Password)
	assert.Nil(suite.T(), actualErr)
	assert.Equal(suite.T(), "", user.Password)
}

func (suite *UserUseCaseTestSuite) TestDeleteById_Success() {
	suite.urm.On("DeleteUser", mockuser.Id).Return(nil)
	actualErr := suite.uc.DeleteById(mockuser.Id)
	assert.Nil(suite.T(), actualErr)
}

func (suite *UserUseCaseTestSuite) TestDeleteById_UserFail() {
	suite.urm.On("DeleteUser", mockuser.Id).Return(errors.New("user not found"))
	actualErr := suite.uc.DeleteById(mockuser.Id)
	assert.Error(suite.T(), actualErr)
	assert.Equal(suite.T(), "user with id 1 not found", actualErr.Error())
}

func (suite *UserUseCaseTestSuite) TestUpdateUser_Success() {

}

func (suite *UserUseCaseTestSuite) TestUpdateUser_UserFail() {
}

func (suite *UserUseCaseTestSuite) TestCreate_Success() {
	suite.urm.On("Create", mockuser).Return(nil)
	actualErr := suite.uc.Create(mockuser)
	assert.Nil(suite.T(), actualErr)
}

func (suite *UserUseCaseTestSuite) TestCreate_UserFail() {

}
func TestUserUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUseCaseTestSuite))
}
