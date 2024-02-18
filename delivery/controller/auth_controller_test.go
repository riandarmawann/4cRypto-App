package controller

import (
	"4crypto/config"
	servicemock "4crypto/mock/service_mock"
	usecasemock "4crypto/mock/usecase_mock"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type AuthControllerTestSuite struct {
	suite.Suite
	//acm *controllermock.AuthControllerMock
	jtm *servicemock.JwtTokenMock
	ac  AuthController
	aum *usecasemock.AuthUseCaseMock
	rg  *gin.RouterGroup
}

func (suite *AuthControllerTestSuite) SetupTest() {
	//suite.acm = new(controllermock.AuthControllerMock)
	suite.jtm = new(servicemock.JwtTokenMock)
	suite.aum = new(usecasemock.AuthUseCaseMock)
	suite.rg = &gin.New().RouterGroup
	suite.ac = *NewAuthController(suite.aum, suite.rg, suite.jtm)
}

func (suite *AuthControllerTestSuite) TestRoute_Success() {
	authGroup := suite.rg.Group(config.AuthGroup)
	//assert.Equal(suite.T(), authGroup.Group())
	authGroup.POST(config.AuthLogin, suite.ac.loginHandler)
	authGroup.GET(config.AuthRefreshToken, suite.ac.refreshTokenHandler)
}

func (suite *AuthControllerTestSuite) LoginHandler_Success() {
}

func (suite *AuthControllerTestSuite) LoginHandler_Fail() {

}

func (suite *AuthControllerTestSuite) RefreshTokenHandler_Success() {

}

func (suite *AuthControllerTestSuite) RefreshTokenHandler_Fail() {

}
func TestAuthControllerTestSuite(t *testing.T) {
	suite.Run(t, new(AuthControllerTestSuite))
}
