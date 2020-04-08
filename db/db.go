package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	"rightswipe/config"
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
