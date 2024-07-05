package service

import (
	"test-task/internal/repo"
)

type UserService interface {
}

type userService struct {
	repository repo.UserRepo
}

func NewUserService(userRepo repo.UserRepo) UserService {
	return &userService{repository: userRepo}
}

