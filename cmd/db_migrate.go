package cmd

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/anupkumarpanwar/Go-Gin-Postgres-Boilerplate/constants"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
)

func DbMigrate() {
	db, err := sql.Open(constants.DbDriver, fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", constants.DbUsername, constants.DbPassword, constants.DbHost, constants.DbPort, constants.DbName, constants.DbSslmode))
	if err != nil {
		log.Fatalf("could not connect to the MySQL database... %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("could not ping DB... %v", err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("could not start sql migration... %v", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		constants.DbName, driver)
	if err != nil {
		log.Fatalf("migration failed... %v", err)
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("error occurred while syncing the database.. %v", err)
	}
	log.Println("database migrated")
}
