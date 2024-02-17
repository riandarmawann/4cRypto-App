package manager

import "4crypto/repository"

type RepoManager interface {
	NewUserRepo() repository.UserRepository
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
