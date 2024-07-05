package manager

import (
	"sync"

	"test-task/infra"
	"test-task/internal/service"
)

type ServiceManager interface {
}

type serviceManager struct {
	infra infra.Infra
	repo  RepoManager
}

// NewServiceManager ...
func NewServiceManager(infra infra.Infra) ServiceManager {
	return &serviceManager{
		infra: infra,
		repo:  NewRepoManager(infra),
	}
}

var (
	userServiceOnce sync.Once
	userService     service.UserService
)

// UserService returns an instance of the UserService implementation through the ServiceManager.
func (sm *serviceManager) UserService() service.UserService {
	userServiceOnce.Do(func() {
		userRepo := sm.repo.UserRepo()
		userService = service.NewUserService(userRepo)
	})
	return userService
}
