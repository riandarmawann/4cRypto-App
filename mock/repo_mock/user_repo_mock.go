package repomock

import (
	"4crypto/model/entity"

	"github.com/stretchr/testify/mock"
)

type UserRepoMock struct {
	mock.Mock
}

func (u *UserRepoMock) GetById(id string) (entity.User, error) {
	args := u.Called(id)
	return args.Get(0).(entity.User), args.Error(1)
}

func (u *UserRepoMock) GetByUsername(username string) (entity.User, error) {
	args := u.Called(username)
	return args.Get(0).(entity.User), args.Error(1)
}

func (u *UserRepoMock) DeleteUser(id string) error {
	args := u.Called(id)
	return args.Error(0)
}

func (u *UserRepoMock) UpdateUser(id string, updatedUser entity.User) error {
	args := u.Called(id, updatedUser)
	return args.Error(0)
}

func (u *UserRepoMock) Create(user entity.User) error {
	args := u.Called(user)
	return args.Error(0)
}
