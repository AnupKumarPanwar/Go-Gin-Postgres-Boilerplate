package cmd

import (
	"github.com/anupkumarpanwar/Go-Gin-Postgres-Boilerplate/config"
	"github.com/anupkumarpanwar/Go-Gin-Postgres-Boilerplate/db"
	"github.com/anupkumarpanwar/Go-Gin-Postgres-Boilerplate/router"
)

func Serve() {
	config.Init()
	db.Init()
	router.Init()
}
