package repo

import (
	"context"
	"database/sql"
	"fmt"
	"test-task/internal/models"
	"time"
)

type UserRepo interface {
	Users() ([]*models.User, error)
	UsersWithFiltersAndPagination(params models.UserFilters, pagination models.Pagination) ([]*models.User, error)
	CreateUser(user *models.User) (uint, error)
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Users() ([]*models.User, error) {
	const op = "repo.userRepo.Users"

	rows, err := r.db.Query(`SELECT * FROM users 1 = 1`)
	if err != nil {
		return nil, fmt.Errorf("error fetching all users: %s %w", op, err)
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Surname, &user.Patronymic, &user.PassportNumber, &user.Address)
		if err != nil {
			return nil, fmt.Errorf("error scanning user row: %s %w", op, err)
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error scaning user row: %s %w", op, err)
	}

	return users, nil
}

func (r *userRepo) UsersWithFiltersAndPagination(params models.UserFilters, pagination models.Pagination) ([]*models.User, error) {
	query := `
		SELECT id, passport_number, surname, name, patronymic, address, created_at, updated_at
		FROM users
		WHERE ($1 = '' OR passport_number = $1)
		  AND ($2 = '' OR surname = $2)
		  AND ($3 = '' OR name = $3)
		ORDER BY id
		LIMIT $4 OFFSET $5
	`
	offset := (pagination.Page - 1) * pagination.PageSize

	rows, err := r.db.QueryContext(context.Background(), query, params.PassportNumber, params.Surname, params.Name, pagination.PageSize, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users: %v", err)
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.ID, &user.PassportNumber, &user.Surname, &user.Name, &user.Patronymic, &user.Address, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user row: %v", err)
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *userRepo) GetUserByID(id int) (*models.User, error) {
	query := `
		SELECT id, passport_number, surname, name, patronymic, address, created_at, updated_at
		FROM users
		WHERE id = $1
	`
	user := &models.User{}
	err := r.db.QueryRowContext(context.Background(), query, id).Scan(&user.ID, &user.PassportNumber, &user.Surname, &user.Name, &user.Patronymic, &user.Address, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user: %v", err)
	}

	return user, nil
}

func (r *userRepo) CreateUser(user *models.User) (uint, error) {
	query := `
		INSERT INTO users (passport_number, surname, name, patronymic, address, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`
	now := time.Now()
	err := r.db.QueryRowContext(context.Background(), query, user.PassportNumber, user.Surname, user.Name, user.Patronymic, user.Address, now, now).Scan(&user.ID)
	if err != nil {
		return 0, fmt.Errorf("failed to insert user: %v", err)
	}

	return user.ID, nil
}
