package main

import (
	"test-task/infra"
	"test-task/internal/api"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Println("Starting application...")

	// Инициализация инфраструктуры
	i := infra.New("config/config.json")
	i.SetMode()

	// Выполнение миграций
	i.SQLClient()
	i.RunSQLMigrations()

	// Запуск API сервера
	api.NewServer(i, i.RedisClient()).Run()
}
