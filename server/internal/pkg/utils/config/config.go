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
	DbHost   string
	DbPort   string
	DbAddr   string
}

func New() *Config {
	return &Config{
		Host:     os.Getenv("APP_URL"),
		Port:     os.Getenv("PORT"),
		DbUser:   os.Getenv("DB_USERNAME"),
		DbPasswd: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
		DbHost:   os.Getenv("DB_HOST"),
		DbPort:   os.Getenv("DB_PORT"),
	}
}

func (conf *Config) GetDbAddr() string {
	return fmt.Sprintf("%s:%s", conf.Host, conf.DbPort)
}
