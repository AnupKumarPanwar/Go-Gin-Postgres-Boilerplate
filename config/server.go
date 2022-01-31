package config

import "Go-Gin-Postgres-Boilerplate/constants"

type ServerConfig struct {
	Host string
	Port string
}

func GetServerConfig() ServerConfig {
	return ServerConfig{
		Host: constants.ServerHost,
		Port: constants.ServerPort,
	}
}
