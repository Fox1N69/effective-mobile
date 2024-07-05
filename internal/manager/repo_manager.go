package manager

import (
	"sync"

	"test-task/infra"
	"test-task/internal/repo"
)

type RepoManager interface {
	UserRepo() repo.UserRepo
}

type repoManager struct {
	infra infra.Infra
}

// NewRepoManager creates a new instance of RepoManager using the provided infrastructure.
func NewRepoManager(infra infra.Infra) RepoManager {
	return &repoManager{infra: infra}
}

var (
	userRepoOnce sync.Once
	userRepo     repo.UserRepo
)

// UserRepo returns an instance of the UserRepo implementation through the RepoManager.
func (rm *repoManager) UserRepo() repo.UserRepo {
	userRepoOnce.Do(func() {
		userRepo = repo.NewUserRepo(rm.infra.GormDB())
	})
	return userRepo
}
