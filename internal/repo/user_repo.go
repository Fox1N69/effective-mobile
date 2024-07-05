package repo

import (
	"database/sql"
	"fmt"
	"test-task/internal/models"
)

type UserRepo interface {
	Users() ([]*models.User, error)
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Users() ([]*models.User, error) {
	const op = "repo.userRepo.Users"

	rows, err := r.db.Query(`SELECT * FROM users`)
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
