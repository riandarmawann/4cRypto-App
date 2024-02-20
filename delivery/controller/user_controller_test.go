package controller

import (
	usecasemock "4crypto/mock/usecase_mock"
	"4crypto/model/entity"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

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

type UserControllerTestSuite struct {
	suite.Suite
	uc *usecasemock.UserUseCaseMock
	rg *gin.Engine
}

func (suite *UserControllerTestSuite) SetupTest() {
	suite.rg = gin.Default()
	suite.uc = &usecasemock.UserUseCaseMock{}
}

func (suite *UserControllerTestSuite) TestGetUserByID_Success() {
	// Mock behavior of user use case method
	suite.uc.On("FindById", mockuser.Id).Return(mockuser, nil)

	// Perform GET request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/"+mockuser.Id, nil)
	suite.rg.ServeHTTP(w, req)

	// Assert HTTP status code
	assert.Equal(suite.T(), http.StatusOK, w.Code)

	// Parse response body
	var responseUser entity.User
	err := json.Unmarshal(w.Body.Bytes(), &responseUser)
	assert.NoError(suite.T(), err)

	// Assert response data
	assert.Equal(suite.T(), mockuser, responseUser)
}

// Implement other test cases similarly...

func TestUserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}
