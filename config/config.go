package config

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

var config Config

func Init() {
	config = Config{
		Server:   GetServerConfig(),
		Database: GetDataBaseConfig(),
	}
}

func GetConfig() Config {
	return config
}
