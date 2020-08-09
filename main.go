package main

import (
	"log"
	"os"

	"github.com/anupkumarpanwar/Go-Gin-Postgres-Boilerplate/cmd"
	"github.com/anupkumarpanwar/Go-Gin-Postgres-Boilerplate/constants"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if len(os.Args) != 2 {
		cmd.Help()
		os.Exit(0)
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
