package cmd

import (
	"Go-Gin-Postgres-Boilerplate/config"
	"Go-Gin-Postgres-Boilerplate/database"
	"Go-Gin-Postgres-Boilerplate/router"
)

func Serve() {
	config.Init()
	database.Init()
	router.Init()
}
