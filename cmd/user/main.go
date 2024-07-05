package main

import (
	"test-task/infra"
	"test-task/internal/api"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Println("Starting application...")
	i := infra.New("config/config.json")
	i.SetMode()
	i.SQLClient()
	i.RunSQLMigrations()

	api.NewServer(i, i.RedisClient()).Run()
}
