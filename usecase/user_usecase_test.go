package usecase

import (
	"errors"
	"fmt"
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

func (suite *UserUseCaseTestSuite) SetupTest() {
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
	// Mock GetByUsername to return a user with the correct Password
	suite.urm.On("GetByUsername", mockuser.Password).Return(mockuser, nil)

	// Call the method under test with an incorrect password
	_, actualErr := suite.uc.FindByUsernamePassword(mockuser.Password, "invalid_password")

	// Check that an error is returned
	assert.Error(suite.T(), actualErr, "invalid username or password")

	// Check the error message
	assert.Equal(suite.T(), "invalid username or password", actualErr.Error(), "invalid username or password")
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
	suite.urm.On("DeleteUser", mockuser.Id).Return(errors.New("failed to delete user"))
	actualErr := suite.uc.DeleteById(mockuser.Id)
	assert.Error(suite.T(), actualErr)
	assert.Equal(suite.T(), "failed to delete user with ID 1: failed to delete user", actualErr.Error())
}

func (suite *UserUseCaseTestSuite) TestUpdateUser_Success() {
	// Persiapkan pengguna yang ada di repository
	existingUser := entity.User{
		Id:        "1",
		Username:  "existing_username",
		Password:  "existing_password",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	// Persiapkan pengguna baru untuk diperbarui
	newUser := entity.User{
		Username: "existing_username",
		Password: "existing_password",
	}
	// Mock repository
	mockRepo := &repomock.UserRepoMock{}
	// Expectation: GetById akan mengembalikan pengguna yang ada
	mockRepo.On("GetById", existingUser.Id).Return(existingUser, nil)
	// Expectation: UpdateUser akan berhasil
	mockRepo.On("UpdateUser", existingUser.Id, mock.Anything).Return(nil)

	// Inisialisasi use case dengan mock repository
	useCase := NewUserUseCase(mockRepo)

	// Panggil fungsi UpdateUser
	err := useCase.UpdateUser(existingUser.Id, newUser)

	// Periksa apakah tidak ada error yang dikembalikan
	assert.Nil(suite.T(), err)

	// Assert bahwa metode GetById dan UpdateUser pada repository telah dipanggil dengan argumen yang sesuai
	mockRepo.AssertCalled(suite.T(), "GetById", existingUser.Id)
	mockRepo.AssertCalled(suite.T(), "UpdateUser", existingUser.Id, mock.Anything)

	// Periksa apakah pengguna telah diperbarui dengan benar
	assert.Equal(suite.T(), newUser.Username, existingUser.Username)
	assert.Equal(suite.T(), newUser.Password, existingUser.Password)
}

func (suite *UserUseCaseTestSuite) TestUpdateUser_UserFail() {
	// Persiapkan pengguna yang tidak ditemukan di repository
	id := "999"
	// Persiapkan pengguna baru untuk diperbarui
	newUser := entity.User{
		Username: "existing_username",
		Password: "existing_password",
	}
	// Mock repository
	mockRepo := &repomock.UserRepoMock{}
	// Expectation: GetById akan mengembalikan kesalahan "user not found"
	mockRepo.On("GetById", id).Return(entity.User{}, errors.New("user not found"))

	// Inisialisasi use case dengan mock repository
	useCase := NewUserUseCase(mockRepo)

	// Panggil fungsi UpdateUser
	err := useCase.UpdateUser(id, newUser)

	// Periksa apakah error yang diharapkan dikembalikan
	expectedErrMsg := fmt.Sprintf("user with id %s not found", id)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedErrMsg, err.Error())

	// Assert bahwa metode GetById pada repository telah dipanggil dengan argumen yang sesuai
	mockRepo.AssertCalled(suite.T(), "GetById", id)

	// Pastikan bahwa UpdateUser tidak dipanggil
	mockRepo.AssertNotCalled(suite.T(), "UpdateUser", mock.Anything, mock.Anything)
}

func (suite *UserUseCaseTestSuite) TestUpdateUser_RepositoryError() {
	// Persiapkan pengguna yang ada di repository
	existingUser := entity.User{
		Id:        "1",
		Username:  "existing_username",
		Password:  "existing_password",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	// Persiapkan pengguna baru untuk diperbarui
	newUser := entity.User{
		Username: "existing_username",
		Password: "existing_password",
	}
	// Mock repository
	mockRepo := &repomock.UserRepoMock{}
	// Expectation: GetById akan mengembalikan pengguna yang ada
	mockRepo.On("GetById", existingUser.Id).Return(existingUser, nil)
	// Expectation: UpdateUser akan mengalami kegagalan
	expectedErr := errors.New("repository error")
	mockRepo.On("UpdateUser", existingUser.Id, mock.Anything).Return(expectedErr)

	// Inisialisasi use case dengan mock repository
	useCase := NewUserUseCase(mockRepo)

	// Panggil fungsi UpdateUser
	err := useCase.UpdateUser(existingUser.Id, newUser)

	// Periksa apakah error yang diharapkan dikembalikan
	expectedErrMsg := fmt.Sprintf("failed to update user with ID %s: %v", existingUser.Id, expectedErr)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedErrMsg, err.Error())

	// Assert bahwa metode GetById dan UpdateUser pada repository telah dipanggil dengan argumen yang sesuai
	mockRepo.AssertCalled(suite.T(), "GetById", existingUser.Id)
	mockRepo.AssertCalled(suite.T(), "UpdateUser", existingUser.Id, mock.Anything)
}

func (suite *UserUseCaseTestSuite) TestCreate_Success() {
	suite.urm.On("Create", mockuser).Return(nil)
	actualErr := suite.uc.Create(mockuser)
	assert.Nil(suite.T(), actualErr)
}

func (suite *UserUseCaseTestSuite) TestCreate_UserFail() {
	suite.urm.On("Create", mockuser).Return(errors.New("failed to create data customer"))
	actualErr := suite.uc.Create(mockuser)
	assert.Error(suite.T(), actualErr)
	assert.Equal(suite.T(), "failed to create data customer", actualErr.Error())
}
func TestUserUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUseCaseTestSuite))
}
