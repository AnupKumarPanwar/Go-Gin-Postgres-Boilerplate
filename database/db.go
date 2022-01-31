package database

import (
	"Go-Gin-Postgres-Boilerplate/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, username, password, dbName, port, sslMode)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
