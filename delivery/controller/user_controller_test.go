package controller

import (
	usecasemock "4crypto/mock/usecase_mock"
	"4crypto/model/entity"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
	rg *gin.RouterGroup
}

func (suite *UserControllerTestSuite) SetupTest() {
	rg := gin.Default()
	suite.rg = rg.Group("/api/v1")
	suite.uc = new(usecasemock.UserUseCaseMock)
}


func (suite *UserControllerTestSuite) TestRoute() {
	 // Panggil method Route() untuk mengatur route
	 controller := NewUserController(suite.uc, suite.rg)
	 controller.Route()
 
	 // Dapatkan daftar rute yang telah ditambahkan ke grup
	 suite.rg := rg.()
 
	 // Periksa setiap rute
	 for _, route := range routes {
		 // Periksa bahwa route telah ditetapkan dengan benar untuk metode DELETE
		 if route.Method == "DELETE" && route.Path == "/api/v1/users/delete/:id" {
			 assert.NotNil(suite.T(), route.Handler)
			 assert.NotNil(suite.T(), route.HandlerFunc)
		 }
 
		 // Periksa bahwa route telah ditetapkan dengan benar untuk metode PUT
		 if route.Method == "PUT" && route.Path == "/api/v1/users/update/:id" {
			 assert.NotNil(suite.T(), route.Handler)
			 assert.NotNil(suite.T(), route.HandlerFunc)
		 }
	 }
}
func (suite *UserControllerTestSuite) TestDeleteUser_Success() {
	// Persiapkan kondisi awal dengan mengatur ekspektasi bahwa pemanggilan use case DeleteById akan berhasil tanpa error
	suite.uc.On("DeleteById", mockuser.Id).Return(nil)

	// Buat HTTP request context
	req, _ := http.NewRequest("DELETE", "/api/v1/user/"+mockuser.Id, nil)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	// Panggil method untuk diuji
	controller := NewUserController(suite.uc, suite.rg)
	controller.DeleteUserByID(ctx)

	// Periksa bahwa pemanggilan use case DeleteById telah terjadi dengan benar
	suite.uc.AssertCalled(suite.T(), "DeleteById", mockuser.Id)
	assert.Equal(suite.T(), http.StatusNoContent, ctx.Writer.Status())
}

func (suite *UserControllerTestSuite) TestDeleteUser_Failed() {

}

func (suite *UserControllerTestSuite) TestUpdateUser_Success() {
	// Persiapkan data yang akan dikirim dalam JSON request
	newUserData := entity.User{
		Username: "updated_username",
		Password: "updated_password",
	}
	requestBody, _ := json.Marshal(newUserData)

	// Persiapkan kondisi awal dengan mengatur ekspektasi bahwa pemanggilan use case UpdateUser akan berhasil tanpa error
	suite.uc.On("UpdateUser", mock.AnythingOfType("string"), newUserData).Return(nil)

	// Buat HTTP request context dengan JSON request body
	req, _ := http.NewRequest("PUT", "/api/v1/user/"+mockuser.Id, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	// Panggil method untuk diuji
	controller := NewUserController(suite.uc, suite.rg)
	controller.UpdateUserByID(ctx)

	// Periksa bahwa pemanggilan use case UpdateUser telah terjadi dengan benar
	suite.uc.AssertCalled(suite.T(), "UpdateUser", mock.AnythingOfType("string"), newUserData)
	assert.Equal(suite.T(), http.StatusNoContent, ctx.Writer.Status())
}

func (suite *UserControllerTestSuite) TestUpdateUser_Failed() {
	// Persiapkan data yang akan dikirim dalam JSON request
	newUserData := entity.User{
		Username: "updated_username",
		Password: "updated_password",
	}
	requestBody, _ := json.Marshal(newUserData)

	// Persiapkan kondisi awal dengan mengatur ekspektasi bahwa pemanggilan use case UpdateUser akan mengembalikan error
	expectedError := errors.New("update user failed")
	suite.uc.On("UpdateUser", mock.AnythingOfType("string"), newUserData).Return(expectedError)

	// Buat HTTP request context dengan JSON request body
	req, _ := http.NewRequest("PUT", "/api/v1/user/"+mockuser.Id, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = req

	// Panggil method untuk diuji
	controller := NewUserController(suite.uc, suite.rg)
	controller.UpdateUserByID(ctx)

	// Periksa bahwa pemanggilan use case UpdateUser telah terjadi dengan benar
	suite.uc.AssertCalled(suite.T(), "UpdateUser", mock.AnythingOfType("string"), newUserData)
	assert.Equal(suite.T(), http.StatusInternalServerError, ctx.Writer.Status())
}
func TestUserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}
