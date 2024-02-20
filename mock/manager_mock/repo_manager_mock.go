package managermock

import (
	"4crypto/repository"

	"github.com/stretchr/testify/mock"
)

type RepoManagerMock struct {
	mock.Mock
}

func (r *RepoManagerMock) NewUserRepo() repository.UserRepository {
	args := r.Called()
	return args.Get(0).(repository.UserRepository)
}
