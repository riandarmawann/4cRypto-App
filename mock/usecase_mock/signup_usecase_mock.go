package usecasemock

import (
	"4crypto/model/entity"

	"github.com/stretchr/testify/mock"
)

type SignUpUseCaseMock struct {
	mock.Mock
}

func (su *SignUpUseCaseMock) SignUp(newUser entity.User) error {
	args := su.Called(newUser)
	return args.Error(0)
}
