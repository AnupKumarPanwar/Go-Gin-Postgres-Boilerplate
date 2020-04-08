package config

import "rightswipe/constants"

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
