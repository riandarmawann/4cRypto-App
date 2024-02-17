package controllermock

import (
	"4crypto/model/dto"
	"4crypto/model/entity"

	"github.com/stretchr/testify/mock"
)

type AuthControllerMock struct {
	mock.Mock
}

func (a *AuthControllerMock) Register(payload entity.User) (entity.User, error) {
	args := a.Called(payload)
	return args.Get(0).(entity.User), args.Error(1)
}

func (a *AuthControllerMock) Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error) {
	args := a.Called(payload)
	return args.Get(0).(dto.AuthResponseDto), args.Error(1)
}
