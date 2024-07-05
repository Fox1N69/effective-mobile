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
	UserByID(id uint) (*models.User, error)
	Create(user *models.User) (*models.User, error)
	Update(id uint, user *models.User) error
	Delete(id uint) error
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
		return nil, fmt.Errorf("%s %w", op, err)
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Surname, &user.Patronymic, &user.PassportNumber, &user.Address)
		if err != nil {
			return nil, fmt.Errorf("%s %w", op, err)
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s %w", op, err)
	}

	return users, nil
}

func (r *userRepo) UsersWithFiltersAndPagination(params models.UserFilters, pagination models.Pagination) ([]*models.User, error) {
	const op = "repo.user_repo.UsersWithFiltersAndPagination"

	query := `
		SELECT *
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
		return nil, fmt.Errorf("%s %w", op, err)
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.ID, &user.PassportNumber, &user.Surname, &user.Name, &user.Patronymic, &user.Address, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("%s %w", op, err)
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *userRepo) UserByID(id uint) (*models.User, error) {
	const op = "repo.user_repo.UserByID"

	query := `
		SELECT *
		FROM users
		WHERE id = $1
	`
	user := &models.User{}
	err := r.db.QueryRowContext(context.Background(), query, id).Scan(&user.ID, &user.PassportNumber, &user.Surname, &user.Name, &user.Patronymic, &user.Address, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("%s %w", op, err)
	}

	return user, nil
}

func (r *userRepo) Create(user *models.User) (*models.User, error) {
	const op = "repo.userRepo.Create"

	query := `
		INSERT INTO users (passport_number, surname, name, patronymic, address, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`

	now := time.Now()
	err := r.db.QueryRowContext(context.Background(), query, user.PassportNumber, user.Surname, user.Name, user.Patronymic, user.Address, now, now).Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("%s %w", op, err)
	}

	user.CreatedAt = now
	user.UpdatedAt = now
	return user, nil
}

func (r *userRepo) Update(id uint, user *models.User) error {
	const op = "repo.user_repo.Update"

	query := `
		UPDATE users
		SET passport_number = $1, surname = $2, name = $3, patronymic = $4, address = $5, updated_at = $6
		WHERE id = $7
	`
	now := time.Now()
	_, err := r.db.ExecContext(context.Background(), query, user.PassportNumber, user.Surname, user.Name, user.Patronymic, user.Address, now, id)
	if err != nil {
		return fmt.Errorf("%s %w", op, err)
	}

	return nil
}

func (r *userRepo) Delete(id uint) error {
	const op = "repo.user_repo.Delete"

	query := `
		DELETE FROM users
		WHERE id = $1
	`
	_, err := r.db.ExecContext(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("%s %w", op, err)
	}

	return nil
}
