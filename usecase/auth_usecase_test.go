package usecase

import (
	"errors"
	"testing"

	servicemock "4crypto/mock/service_mock"
	usecasemock "4crypto/mock/usecase_mock"
	"4crypto/model/dto"
	"4crypto/model/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var (
	mockAuthRequestDto = dto.AuthRequestDto{
		Username: "admin",
		Password: "password",
	}

	mockUser = entity.User{
		Id:       "1",
		Name:     "Admin",
		Username: "admin",
		Role:     "admin",
	}

	mockAuthResponseDto = dto.AuthResponseDto{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
	}
)

type AuthUseCaseTestSuite struct {
	suite.Suite
	uum *usecasemock.UserUseCaseMock
	jtm *servicemock.JwtTokenMock
	au  AuthUseCase
}

func (suite *AuthUseCaseTestSuite) SetupTest() {
	suite.uum = new(usecasemock.UserUseCaseMock)
	suite.jtm = new(servicemock.JwtTokenMock)
	suite.au = NewAuthUseCase(suite.uum, suite.jtm)
}

func (suite *AuthUseCaseTestSuite) TestRegister_Success() {
	suite.uum.On("FindByUsernamePassword", mockAuthRequestDto.Username, mockAuthRequestDto.Password).Return(entity.User{}, nil)
	actualUser, actualErr := suite.au.Register(mockUser)
	assert.Nil(suite.T(), actualErr)
	assert.Equal(suite.T(), mockUser, actualUser)
}
func (suite *AuthUseCaseTestSuite) TestLogin_Success() {
	suite.uum.On("FindByUsernamePassword", mockAuthRequestDto.Username, mockAuthRequestDto.Password).Return(mockUser, nil)
	suite.jtm.On("GenerateToken", mockUser).Return(mockAuthResponseDto, nil)
	actualToken, actualErr := suite.au.Login(mockAuthRequestDto)
	assert.Nil(suite.T(), actualErr)
	assert.Equal(suite.T(), mockAuthResponseDto, actualToken)
}

func (suite *AuthUseCaseTestSuite) TestLogin_Fail() {
	suite.uum.On("FindByUsernamePassword", mockAuthRequestDto.Username, mockAuthRequestDto.Password).Return(entity.User{}, errors.New("some error"))
	actualToken, actualErr := suite.au.Login(mockAuthRequestDto)
	assert.NotNil(suite.T(), actualErr)
	assert.Equal(suite.T(), "", actualToken.Token)
}

func (suite *AuthUseCaseTestSuite) TestLogin_Fail_GenerateToken() {
	suite.uum.On("FindByUsernamePassword", mockAuthRequestDto.Username, mockAuthRequestDto.Password).
		Return(mockUser, nil)
	suite.jtm.On("GenerateToken", mockUser).
		Return(dto.AuthResponseDto{}, errors.New("token generation failed"))

	actualToken, actualErr := suite.au.Login(mockAuthRequestDto)

	assert.NotNil(suite.T(), actualErr)
	assert.Equal(suite.T(), "", actualToken.Token)
	assert.Equal(suite.T(), "token generation failed", actualErr.Error())
}
func TestAuthUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(AuthUseCaseTestSuite))
}
