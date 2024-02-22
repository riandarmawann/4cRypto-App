package manager

import "4crypto/usecase"

type UseCaseManager interface {
	NewUserUseCase() usecase.UserUseCase
	NewCryptoUseCase() usecase.CryptoUseCase
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

func (u *useCaseManager) NewCryptoUseCase() usecase.CryptoUseCase {
	return usecase.NewCryptoUseCase(u.repo.NewCryptoRepo())
}
