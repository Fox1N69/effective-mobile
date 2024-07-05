package infra

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type SQLClient struct {
	DB *sql.DB
}

func NewSQLClient() *SQLClient {
	return &SQLClient{}
}

func (s *SQLClient) Connect(user, password, host, port, dbname string) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logrus.Fatalf("[infra][Connect][sql.Open] %v", err)
	}

	if err := db.Ping(); err != nil {
		logrus.Fatalf("[infra][Connect][db.Ping] %v", err)
	}

	s.DB = db
}

func (s *SQLClient) Close() {
	if s.DB != nil {
		s.DB.Close()
		logrus.Println("Connection to PostgreSQL closed")
	}
}

func (s *SQLClient) SqlMigrate() {
	const op = "infra.SQLClient.SqlMigrate"
	driver, err := postgres.WithInstance(s.DB, &postgres.Config{})
	if err != nil {
		logrus.Fatal(op, err)
	}
	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		logrus.Fatal(op, err)
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		logrus.Fatal(op, err)
	}
	logrus.Println("SQL migrations completed")
}
