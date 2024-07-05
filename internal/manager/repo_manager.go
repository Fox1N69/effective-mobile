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

func NewRepoManager(infra infra.Infra) RepoManager {
	return &repoManager{infra: infra}
}

var (
	userRepoOnce sync.Once
	userRepo     repo.UserRepo
)

func (rm *repoManager) UserRepo() repo.UserRepo {
	userRepoOnce.Do(func() {
		userRepo = repo.NewUserRepo(rm.infra.GormDB())
	})
	return userRepo
}
