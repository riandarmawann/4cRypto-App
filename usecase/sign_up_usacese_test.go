package usecase

import (
	repomock "4crypto/mock/repo_mock"
	usecasemock "4crypto/mock/usecase_mock"
	"4crypto/model/entity"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// MockUserRepository is a mock implementation of UserRepository interface
type SignUpMock struct {
	mock.Mock
}

var (
	signupmock = entity.User{
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

type SignUpUseCaseTestSuite struct {
	suite.Suite
	sm   *usecasemock.SignUpUseCaseMock
	usrm *repomock.UserRepoMock
	su   SignUpUseCase
}

func (suite *SignUpUseCaseTestSuite) SetupTest() {
	suite.usrm = new(repomock.UserRepoMock)
	suite.sm = new(usecasemock.SignUpUseCaseMock)
	suite.su = NewSignUpUseCase(suite.usrm)
}

func (suite *SignUpUseCaseTestSuite) TestSignUp_Success() {
	suite.usrm.On("GetByUsername", signupmock.Username).Return(entity.User{}, errors.New("not found"))
	//suite.usrm.On("GetByEmail", signupmock.Email).Return(entity.User{}, errors.New("not found"))
	suite.usrm.On("Create", signupmock).Return(nil)
	err := suite.su.SignUp(signupmock)
	assert.NoError(suite.T(), err)
}

func (suite *SignUpUseCaseTestSuite) TestSignUp_Failed() {
	suite.usrm.On("GetByUsername", signupmock.Username).Return(entity.User{}, errors.New("not found"))
	//suite.usrm.On("GetByEmail", signupmock.Email).Return(entity.User{}, errors.New("not found"))
	suite.usrm.On("Create", signupmock).Return(errors.New("failed to create data customer"))
	err := suite.su.SignUp(signupmock)
	assert.Error(suite.T(), err)
}

func (suite *SignUpUseCaseTestSuite) TestSignUp_UsernameExists() {
	suite.usrm.On("GetByUsername", signupmock.Username).Return(signupmock, nil)
	err := suite.su.SignUp(signupmock)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "username already exists", err.Error())
}

// func (suite *SignUpUseCaseTestSuite) TestSignUp_EmailExists() {
// 	suite.usrm.On("GetByUsername", signupmock.Username).Return(entity.User{}, errors.New("not found"))
// 	suite.usrm.On("GetByEmail", signupmock.Email).Return(signupmock, nil)
// 	err := suite.su.SignUp(signupmock)
// 	assert.Error(suite.T(), err)
// 	assert.Equal(suite.T(), "email already exists", err.Error())
// }

func TestSignUpUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(SignUpUseCaseTestSuite))
}
