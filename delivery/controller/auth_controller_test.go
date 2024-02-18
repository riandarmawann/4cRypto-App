package controller

import (
	// "4crypto/config"
	servicemock "4crypto/mock/service_mock"
	usecasemock "4crypto/mock/usecase_mock"
	"4crypto/model/dto"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var (
	mockAuthRequestDto = dto.AuthRequestDto{
		Username: "1",
		Password: "1",
	}
	mockAuthResponseDto = dto.AuthResponseDto{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
	}

	// mockTokenJWT = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
)

type AuthControllerTestSuite struct {
	suite.Suite
	//acm *controllermock.AuthControllerMock
	jtm *servicemock.JwtTokenMock
	// ac  AuthController
	aum *usecasemock.AuthUseCaseMock
	rg  *gin.RouterGroup
}

func (suite *AuthControllerTestSuite) SetupTest() {
	//suite.acm = new(controllermock.AuthControllerMock)
	suite.jtm = new(servicemock.JwtTokenMock)
	suite.aum = new(usecasemock.AuthUseCaseMock)
	rg := gin.Default()
	suite.rg = rg.Group("/api/v1")
	// suite.ac = *NewAuthController(suite.aum, suite.rg, suite.jtm)
}

// func (suite *AuthControllerTestSuite) TestRoute_Success() {
// authGroup := suite.rg.Group(config.AuthGroup)
//assert.Equal(suite.T(), authGroup.Group())
// authGroup.POST(config.AuthLogin, suite.ac.loginHandler)
// authGroup.GET(config.AuthRefreshToken, suite.ac.refreshTokenHandler)
// }

func (suite *AuthControllerTestSuite) TestLoginHandler_Success() {

	suite.aum.On("Login", mockAuthRequestDto).Return(mockAuthResponseDto, nil)
	// Record -> menangkap response
	record := httptest.NewRecorder()
	// Simulasi pengiriman request menggunakan payload dalam bentuk JSON
	authRequestDtoJSON, err := json.Marshal(mockAuthRequestDto)
	assert.NoError(suite.T(), err)
	// Simulasi hit ke path API -> /api/v1/bills
	request, err := http.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(authRequestDtoJSON))
	// request.Header.Set("Authorization", "Bearer "+mockTokenJWT)
	assert.NoError(suite.T(), err)
	// kita panggil Controller
	authController := NewAuthController(suite.aum, suite.rg, suite.jtm)
	authController.Route()
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = request
	authController.loginHandler(ctx)
	assert.Equal(suite.T(), http.StatusOK, record.Code)

}

func (suite *AuthControllerTestSuite) TestLoginHandler_FailBind() {
	// suite.aum.On("Login", mockAuthRequestDto).Return(mockAuthResponseDto, nil)
	// Record -> menangkap response
	record := httptest.NewRecorder()
	// Simulasi pengiriman request menggunakan payload dalam bentuk JSON

	request, err := http.NewRequest(http.MethodPost, "/api/v1/auth/login", nil)
	// request.Header.Set("Authorization", "Bearer "+mockTokenJWT)
	assert.NoError(suite.T(), err)
	// kita panggil Controller
	authController := NewAuthController(suite.aum, suite.rg, suite.jtm)
	authController.Route()
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = request
	authController.loginHandler(ctx)
	assert.Equal(suite.T(), http.StatusBadRequest, record.Code)
}

func (suite *AuthControllerTestSuite) TestLoginHandler_Fail() {
	suite.aum.On("Login", mockAuthRequestDto).Return(dto.AuthResponseDto{}, errors.New("error"))
	// Record -> menangkap response
	record := httptest.NewRecorder()
	// Simulasi pengiriman request menggunakan payload dalam bentuk JSON

	authRequestDtoJSON, err := json.Marshal(mockAuthRequestDto)
	assert.NoError(suite.T(), err)

	request, err := http.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(authRequestDtoJSON))
	// request.Header.Set("Authorization", "Bearer "+mockTokenJWT)
	assert.NoError(suite.T(), err)
	// kita panggil Controller
	authController := NewAuthController(suite.aum, suite.rg, suite.jtm)
	authController.Route()
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = request
	authController.loginHandler(ctx)
	assert.Equal(suite.T(), http.StatusInternalServerError, ctx.Writer.Status())
}

func (suite *AuthControllerTestSuite) TestRefreshTokenHandler_Success() {
	suite.jtm.On("RefreshToken", "").Return(dto.AuthResponseDto{}, nil)

	record := httptest.NewRecorder()

	request, err := http.NewRequest(http.MethodPost, "/api/v1/auth/refresh-token", nil)
	// request.Header.Set("Authorization", "Bearer "+mockTokenJWT)
	assert.NoError(suite.T(), err)
	// kita panggil Controller
	authController := NewAuthController(suite.aum, suite.rg, suite.jtm)
	authController.Route()
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = request
	authController.refreshTokenHandler(ctx)
	assert.Equal(suite.T(), http.StatusCreated, record.Code)
}

func (suite *AuthControllerTestSuite) TestRefreshTokenHandler_Failed() {
	suite.jtm.On("RefreshToken", "").Return(dto.AuthResponseDto{}, errors.New("error"))

	record := httptest.NewRecorder()

	request, err := http.NewRequest(http.MethodPost, "/api/v1/auth/refresh-token", nil)
	// request.Header.Set("Authorization", "Bearer "+mockTokenJWT)
	assert.NoError(suite.T(), err)
	// kita panggil Controller
	authController := NewAuthController(suite.aum, suite.rg, suite.jtm)
	authController.Route()
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = request
	authController.refreshTokenHandler(ctx)
	assert.Equal(suite.T(), http.StatusUnauthorized, record.Code)
}

func TestAuthControllerTestSuite(t *testing.T) {
	suite.Run(t, new(AuthControllerTestSuite))
}
