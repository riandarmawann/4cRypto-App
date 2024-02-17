package manager

import "4crypto/usecase"

type UseCaseManager interface {
	NewUserUseCase() usecase.UserUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{repo: repo}
}

func (u *useCaseManager) NewUserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repo.NewUserRepo())
}
