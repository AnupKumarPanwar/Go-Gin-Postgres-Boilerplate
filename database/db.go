package database

import (
	"fmt"
	"log"

	"github.com/anupkumarpanwar/Go-Gin-Postgres-Boilerplate/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Init() {
	databaseConfig := config.GetConfig().Database
	host := databaseConfig.Host
	port := databaseConfig.Port
	username := databaseConfig.Username
	dbName := databaseConfig.DBName
	sslMode := databaseConfig.SSLMode
	password := databaseConfig.Password
	var err error
	db, err = gorm.Open(databaseConfig.Driver, fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", username, password, host, port, dbName, sslMode))
	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
