package main

import (
	"test-task/infra"
	"test-task/internal/api"
)

func main() {
	// Connect to config and set mode
	i := infra.New("config/config.json")
	i.SetMode()

	// Connect and migrate database
	i.SQLClient()
	i.RunSQLMigrations()

	// Start Api server
	api.NewServer(i).Run()
}
