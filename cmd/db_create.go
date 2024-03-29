package cmd

import (
	"Go-Gin-Postgres-Boilerplate/constants"
	"bytes"
	"log"
	"os/exec"
)

func DbCreate() {
	cmd := exec.Command("createdb", "-p", constants.DbPort, "-h", constants.DbHost, "-U", constants.DbUsername, "-e", constants.DbName)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Fatalf("unable to create database : %v", err)
	}
	log.Println("database created successfully")
}
