package main

import (
	"Go-Gin-Postgres-Boilerplate/cmd"
	"Go-Gin-Postgres-Boilerplate/constants"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if len(os.Args) != 2 {
		cmd.Serve()
	}
	command := os.Args[1]
	switch command {
	case constants.Serve:
		cmd.Serve()
		break
	case constants.DbCreate:
		cmd.DbCreate()
		break
	case constants.DbMigrate:
		cmd.DbMigrate()
		break
	case constants.DbDrop:
		cmd.DbDrop()
		break
	default:
		cmd.Help()
		os.Exit(0)
	}
}
