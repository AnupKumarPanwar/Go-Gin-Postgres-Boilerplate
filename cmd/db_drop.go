package cmd

import (
	"bytes"
	"log"
	"os/exec"
	"rightswipe/constants"
)

func DbDrop() {
	cmd := exec.Command("dropdb", "-h", constants.DbHost, "-U", constants.DbUsername, constants.DbName)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Fatalf("unable to drop database : %v", err)
	}
	log.Println("database dropped successfully")
}
