package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DB_DRIVER             string
	DB_CONNECTION_STRING  string
	SERVER_PORT           string
	TOKEN_SYMMETRIC_KEY   string
	ACCESS_TOKEN_DURATION string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	config := &Config{
		DB_DRIVER:             os.Getenv("DB_DRIVER"),
		DB_CONNECTION_STRING:  os.Getenv("DB_CONNECTION_STRING"),
		SERVER_PORT:           os.Getenv("SERVER_PORT"),
		TOKEN_SYMMETRIC_KEY:   os.Getenv("TOKEN_SYMMETRIC_KEY"),
		ACCESS_TOKEN_DURATION: os.Getenv("ACCESS_TOKEN_DURATION"),
	}

	return config, nil
}
