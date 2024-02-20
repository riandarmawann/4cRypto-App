package repository

import (
	"4crypto/model/entity"
	"database/sql"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var (
	mockUser = entity.User{
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

type UserRepoTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    UserRepository
}

func (suite *UserRepoTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	suite.mockDB = db
	suite.mockSql = mock
	suite.repo = NewUserRepository(suite.mockDB)
}

func (suite *UserRepoTestSuite) TestGetId_Success() {
	suite.mockSql.ExpectQuery("SELECT (.+) FROM users").WithArgs(mockUser.Id).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "username", "password", "role", "created_at", "updated_at"}).
		AddRow(
			mockUser.Id,
			mockUser.Name,
			mockUser.Email,
			mockUser.Username,
			mockUser.Password,
			mockUser.Role,
			mockUser.CreatedAt,
			mockUser.UpdatedAt,
		))
	actualUser, actualErr := suite.repo.GetById(mockUser.Id)
	assert.Nil(suite.T(), actualErr)
	assert.Equal(suite.T(), mockUser, actualUser)
}
func (suite *UserRepoTestSuite) TestGetId_UserFail() {
	suite.mockSql.ExpectQuery("SELECT (.+) FROM users").WithArgs("XX").WillReturnError(errors.New("error"))
	_, actualErr := suite.repo.GetById("XX")
	assert.Error(suite.T(), actualErr)
	assert.Equal(suite.T(), "error", actualErr.Error())
}
func (suite *UserRepoTestSuite) TestGetUsername_Success() {
	suite.mockSql.ExpectQuery("SELECT (.+) FROM users").WithArgs(mockUser.Username).WillReturnRows(sqlmock.NewRows([]string{"id", "email", "username", "password", "role"}).
		AddRow(
			mockUser.Id,
			mockUser.Email,
			mockUser.Username,
			mockUser.Password,
			mockUser.Role,
		))
	actualUser, actualErr := suite.repo.GetByUsername(mockUser.Username)
	fmt.Println(actualUser, mockUser)
	fmt.Println(actualErr)
	assert.NoError(suite.T(), actualErr)
	assert.Equal(suite.T(), mockUser.Email, actualUser.Email)
}

func (suite *UserRepoTestSuite) TestGetUsername_UserFail() {
	suite.mockSql.ExpectQuery("SELECT (.+) FROM users").WithArgs("XX").WillReturnError(errors.New("error"))
	_, actualErr := suite.repo.GetByUsername("XX")
	assert.Error(suite.T(), actualErr)
}

func (suite *UserRepoTestSuite) TestDeleteUser_Success() {
	suite.mockSql.ExpectExec("DELETE FROM users").WithArgs(mockUser.Id).WillReturnResult(sqlmock.NewResult(1, 1))
	actualErr := suite.repo.DeleteUser(mockUser.Id)
	assert.NoError(suite.T(), actualErr)
}

func (suite *UserRepoTestSuite) TestDeleteUser_UserFail() {
	suite.mockSql.ExpectExec("DELETE FROM users").WithArgs("XX").WillReturnError(errors.New("error"))
	actualErr := suite.repo.DeleteUser("XX")
	assert.Error(suite.T(), actualErr)
}

func (suite *UserRepoTestSuite) TestUpdateUser_Success() {
	suite.mockSql.ExpectExec("UPDATE users").
		WithArgs(mockUser.Name, mockUser.Email, mockUser.Username, mockUser.Password, mockUser.Role, mockUser.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	actualErr := suite.repo.UpdateUser(mockUser.Id, mockUser)
	assert.NoError(suite.T(), actualErr)
}

func (suite *UserRepoTestSuite) TestUpdateUser_UserFail() {
	suite.mockSql.ExpectExec("UPDATE users").
		WithArgs(mockUser.Name, mockUser.Email, mockUser.Username, mockUser.Password, mockUser.Role, mockUser.Id).
		WillReturnError(errors.New("error"))

	actualErr := suite.repo.UpdateUser(mockUser.Id, mockUser)
	assert.Error(suite.T(), actualErr)
}

func (suite *UserRepoTestSuite) TestCreate_Success() {
	suite.mockSql.ExpectExec("INSERT INTO users").
		WithArgs(mockUser.Name, mockUser.Email, mockUser.Username, mockUser.Password, mockUser.Role).
		WillReturnResult(sqlmock.NewResult(1, 1))

	actualErr := suite.repo.Create(mockUser)
	assert.NoError(suite.T(), actualErr)
}

func (suite *UserRepoTestSuite) TestCreate_UserFail() {
	suite.mockSql.ExpectExec("INSERT INTO users").
		WithArgs(mockUser.Id, mockUser.Name, mockUser.Email, mockUser.Username, mockUser.Password, mockUser.Role).
		WillReturnError(errors.New("error"))

	actualErr := suite.repo.Create(mockUser)
	assert.Error(suite.T(), actualErr)
}

func TestUserRepoTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
}
