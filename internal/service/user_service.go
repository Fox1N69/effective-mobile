package service

import (
	"test-task/internal/models"
	"test-task/internal/repo"
	"time"
)

type UserService interface {
	Users() ([]*models.User, error)
	UsersWithFiltersAndPagination(params models.UserFilters, pagination models.Pagination) ([]*models.User, error)
	CreateUser(user *models.User) (uint, error)
	UpdateUser(id uint, user *models.User) error
	DeleteUser(id uint) error
}

type userService struct {
	repository repo.UserRepo
}

func NewUserService(userRepo repo.UserRepo) UserService {
	return &userService{repository: userRepo}
}

func (s *userService) Users() ([]*models.User, error) {
	return s.repository.Users()
}

func (s *userService) UsersWithFiltersAndPagination(params models.UserFilters, pagination models.Pagination) ([]*models.User, error) {
	return s.repository.UsersWithFiltersAndPagination(params, pagination)
}

func (s *userService) CreateUser(user *models.User) (uint, error) {
	return s.repository.Create(user)
}

func (s *userService) UpdateUser(id uint, user *models.User) error {
	user.UpdatedAt = time.Now()
	return s.repository.Update(id, user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.repository.Delete(id)
}
