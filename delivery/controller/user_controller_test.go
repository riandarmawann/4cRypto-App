package controller

import (
	"4crypto/config"
	usecasemock "4crypto/mock/usecase_mock"
	"4crypto/model/dto/res"
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
	uc := NewUserController(suite.uc, suite.rg)
	uc.Route()
	suite.rg.POST(config.CreateUser, uc.Create)
	suite.rg.GET(config.UserGetByID, uc.FindById)
	suite.rg.DELETE(config.DeleteUserByID, uc.DeleteUserByID)
	suite.rg.PUT(config.UpdateUserByID, uc.UpdateUserByID)
}
func (suite *UserControllerTestSuite) TestDeleteUser_Success() {
	// Persiapkan kondisi awal dengan mengatur ekspektasi bahwa pemanggilan use case DeleteUserByID akan berhasil tanpa error
	suite.uc.On("DeleteUserByID", []string{"1"}).Return(nil)

	// Membuat request HTTP tes
	req, _ := http.NewRequest("DELETE", "/api/v1/user/:id", nil)
	w := httptest.NewRecorder()

	// Membuat context Gin untuk mensimulasikan request
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})

	// Memanggil handler DeleteUserByID
	uc := NewUserController(suite.uc, suite.rg)
	uc.DeleteUserByID(c)

	// Memeriksa status kode HTTP
	assert.Equal(suite.T(), http.StatusNoContent, w.Code)

	// Periksa bahwa pemanggilan use case DeleteUserByID telah terjadi dengan benar
	suite.uc.AssertCalled(suite.T(), "DeleteUserByID", "1")
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

func (suite *UserControllerTestSuite) TestFindById_Success() {
	// Membuat objek pengguna yang diharapkan
	expectedUser := &mockuser

	// Mengatur pengembalian nilai yang diharapkan dari FindById
	suite.uc.On("FindById", mock.Anything).Return(expectedUser, nil)

	// Membuat request HTTP tes
	req, _ := http.NewRequest("GET", "/api/v1/user/1", nil)
	w := httptest.NewRecorder()

	// Membuat context Gin untuk mensimulasikan request
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})

	// Memanggil handler FindById
	uc := NewUserController(suite.uc, suite.rg)
	uc.FindById(c)

	// Memeriksa status kode HTTP
	assert.Equal(suite.T(), http.StatusOK, w.Code)

	// Memeriksa response body
	var resBody res.CommonResponse
	err := json.Unmarshal(w.Body.Bytes(), &resBody)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), http.StatusOK, resBody.Code)
	assert.Equal(suite.T(), "Success", resBody.Status)
	assert.Equal(suite.T(), "Retrieved data successfully", resBody.Message)

	// Melakukan konversi dari interface{} ke *entity.User
	actualUser, ok := resBody.Data.(*entity.User)
	assert.True(suite.T(), ok)

	// Memeriksa bahwa data pengguna yang diharapkan sesuai dengan yang diterima
	assert.Equal(suite.T(), expectedUser, actualUser)
}

func (suite *UserControllerTestSuite) TestFindById_Failed() {
	// Persiapkan ekspektasi bahwa pemanggilan use case FindById akan mengembalikan error
	expectedError := errors.New("user not found")
	suite.uc.On("FindById", mock.Anything).Return(nil, expectedError)

	// Membuat request HTTP tes
	req, _ := http.NewRequest("GET", "/api/v1/user/1", nil)
	w := httptest.NewRecorder()

	// Membuat context Gin untuk mensimulasikan request
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})

	// Memanggil handler FindById
	uc := NewUserController(suite.uc, suite.rg)
	uc.FindById(c)

	// Memeriksa status kode HTTP
	assert.Equal(suite.T(), http.StatusNotFound, w.Code)

	// Memeriksa response body
	var resBody map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resBody)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), http.StatusNotFound, int(resBody["code"].(float64)))
	assert.Equal(suite.T(), "Error", resBody["status"])
	assert.Equal(suite.T(), "user not found", resBody["error"])
}

func (suite *UserControllerTestSuite) TestCreate_Success() {
	// Persiapkan data yang akan dikirim dalam JSON request
	newUserData := entity.User{
		Username: "new_user",
		Password: "password123",
		Email:    "newuser@example.com",
	}
	requestBody, _ := json.Marshal(newUserData)

	// Persiapkan kondisi awal dengan mengatur ekspektasi bahwa pemanggilan use case CreateUser akan berhasil tanpa error
	suite.uc.On("CreateUser", newUserData).Return(nil)

	// Buat HTTP request context dengan JSON request body
	req, _ := http.NewRequest("POST", "/api/v1/user/create", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Membuat context Gin untuk mensimulasikan request
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	// Panggil method untuk diuji
	controller := NewUserController(suite.uc, suite.rg)
	controller.Create(ctx)

	// Periksa bahwa pemanggilan use case CreateUser telah terjadi dengan benar
	suite.uc.AssertCalled(suite.T(), "CreateUser", newUserData)
	assert.Equal(suite.T(), http.StatusCreated, w.Code)
}

func (suite *UserControllerTestSuite) TestCreate_Failed() {

}
func TestUserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}
