package postgres

import (
	"database/sql"
	"fmt"
)

type Storage struct {
	db *sql.DB
}

func NewPostgres(storagePath string) (*Storage, error) {
	const op = "storage.postgres.NewPostgres"

	db, err := sql.Open("postgresql", "./user.db")
	if err != nil {
		return nil, fmt.Errorf("s%: w%", op, err)
	}

	return &Storage{db: db}, nil
}
