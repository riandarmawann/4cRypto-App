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
