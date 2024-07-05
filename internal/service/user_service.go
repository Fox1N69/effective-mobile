package service

import (
	"errors"
	"strings"
	"test-task/infra/clients"
	"test-task/internal/models"
	"test-task/internal/repo"
	"time"
)

type UserService interface {
	Users() ([]*models.User, error)
	UsersWithFiltersAndPagination(params models.UserFilters, pagination models.Pagination) ([]*models.User, error)
	CreateUser(passportNumber string) (*models.User, error)
	UpdateUser(id uint, user *models.User) error
	DeleteUser(id uint) error
}

type userService struct {
	repository    repo.UserRepo
	userApiClient *clients.UserAPIClient
}

func NewUserService(userRepo repo.UserRepo, userApiClinet *clients.UserAPIClient) UserService {
	return &userService{
		repository:    userRepo,
		userApiClient: userApiClinet,
	}
}

func (s *userService) Users() ([]*models.User, error) {
	return s.repository.Users()
}

func (s *userService) UsersWithFiltersAndPagination(params models.UserFilters, pagination models.Pagination) ([]*models.User, error) {
	return s.repository.UsersWithFiltersAndPagination(params, pagination)
}

func (s *userService) CreateUser(passportNumber string) (*models.User, error) {
	// Разделите паспорт на серию и номер
	parts := strings.Split(passportNumber, " ")
	if len(parts) != 2 {
		return nil, errors.New("invalid passport number format")
	}
	passportSerie := parts[0]
	passportNumber = parts[1]

	// Получите информацию из внешнего API
	peopleInfo, err := s.userApiClient.FetchUserInfo(passportSerie, passportNumber)
	if err != nil {
		return nil, err
	}

	// Создайте нового пользователя с полученной информацией
	user := &models.User{
		PassportNumber: passportNumber,
		Surname:        peopleInfo.Surname,
		Name:           peopleInfo.Name,
		Patronymic:     peopleInfo.Patronymic,
		Address:        peopleInfo.Address,
	}

	createdUser, err := s.repository.Create(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (s *userService) UpdateUser(id uint, user *models.User) error {
	user.UpdatedAt = time.Now()
	return s.repository.Update(id, user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.repository.Delete(id)
}
