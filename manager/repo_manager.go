package manager

import "4crypto/repository"

type RepoManager interface {
	NewUserRepo() repository.UserRepository
	NewCryptoRepo() repository.CryptoRepository
}

type repoManager struct {
	infra InfraManager
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}

func (r *repoManager) NewUserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infra.Conn())
}

func (r *repoManager) NewCryptoRepo() repository.CryptoRepository {
	return repository.NewCryptoRepository(r.infra.Conn())
}
