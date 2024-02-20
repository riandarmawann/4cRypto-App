package usecase

import (
	"4crypto/model/entity"
	"time"

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

type signuptestsuite struct {
	suite.Suite
}

func (suite *signuptestsuite) SetupTest() {

}

func (suite *signuptestsuite) TestSignUp_Success() {

}
