package usecasemock

import (
	"4crypto/model/entity"

	"github.com/stretchr/testify/mock"
)

type UserUseCaseMock struct {
	mock.Mock
}

func (u *UserUseCaseMock) FindById(id string) (entity.User, error) {
	args := u.Called(id)
	return args.Get(0).(entity.User), args.Error(1)

}
func (u *UserUseCaseMock) FindByUsernamePassword(username string, password string) (entity.User, error) {
	args := u.Called(username, password)
	return args.Get(0).(entity.User), args.Error(1)
}

func (u *UserUseCaseMock) DeleteById(id string) error {
	args := u.Called(id)
	return args.Error(0)
}

func (u *UserUseCaseMock) UpdateUser(id string, newUser entity.User) error {
	args := u.Called(id, newUser)
	return args.Error(0)
}

func (u *UserUseCaseMock) Create(user entity.User) error {
	args := u.Called(user)
	return args.Error(0)
}
