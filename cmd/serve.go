package cmd

import (
	"rightswipe/config"
	"rightswipe/db"
	"rightswipe/router"
)

func Serve() {
	config.Init()
	db.Init()
	router.Init()
}
