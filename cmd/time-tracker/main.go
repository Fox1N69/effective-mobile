package main

import (
	_ "test-task/cmd/time-tracker/docs"
	"test-task/infra"
	"test-task/internal/api"
)

// @title Tag Service API
// @varsion 1.0
// @description tag for service api

// @host localhost:4000
// @basePath /api
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
