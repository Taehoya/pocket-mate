package config

import (
	"fmt"
	"os"
)

type Config struct {
	Host     string
	Port     string
	DbUser   string
	DbPasswd string
	DbName   string
	DbAddr   string
}

func MakeConfig() *Config {
	return &Config{
		Host:     os.Getenv("APP_URL"),
		Port:     os.Getenv("PORT"),
		DbUser:   os.Getenv("DB_USERNAME"),
		DbPasswd: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
		DbAddr:   fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
	}
}
