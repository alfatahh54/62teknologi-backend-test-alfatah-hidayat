package config

import "os"

type DbConfig struct {
	Host string
	User string
	Pass string
	Name string
	Port string
}

var DBConfig DbConfig

func init() {
	DBConfig.User = os.Getenv("DB_USER")
	DBConfig.Pass = os.Getenv("DB_PASS")
	DBConfig.Host = os.Getenv("DB_HOST")
	DBConfig.Port = os.Getenv("DB_PORT")
	DBConfig.Name = os.Getenv("DB_NAME")
}
