package service

import (
	"test-task/internal/models"
	"test-task/internal/repo"
)

type UserService interface {
	Users() ([]*models.User, error)
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
