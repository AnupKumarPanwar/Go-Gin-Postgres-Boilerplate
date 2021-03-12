package cmd

import (
	"github.com/anupkumarpanwar/Go-Gin-Postgres-Boilerplate/config"
	"github.com/anupkumarpanwar/Go-Gin-Postgres-Boilerplate/database"
	"github.com/anupkumarpanwar/Go-Gin-Postgres-Boilerplate/router"
)

func Serve() {
	config.Init()
	database.Init()
	router.Init()
}
