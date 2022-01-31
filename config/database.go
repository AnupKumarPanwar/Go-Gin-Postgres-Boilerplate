package config

import "Go-Gin-Postgres-Boilerplate/constants"

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	DBName   string
	Password string
	Driver   string
	SSLMode  string
}

func GetDataBaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Host:     constants.DbHost,
		Port:     constants.DbPort,
		Username: constants.DbUsername,
		DBName:   constants.DbName,
		Password: constants.DbPassword,
		Driver:   constants.DbDriver,
		SSLMode:  constants.DbSslmode,
	}
}
