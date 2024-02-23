package managermock

import (
	"4crypto/usecase"

	"github.com/stretchr/testify/mock"
)

type UseCaseManagerMock struct {
	mock.Mock
}

func (m *UseCaseManagerMock) NewUserUseCase() usecase.UserUseCase {
	args := m.Called()
	return args.Get(0).(usecase.UserUseCase)
}
