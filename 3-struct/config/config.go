package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errors.New("не удалось найти .env файл")
	}

	key := os.Getenv("API_KEY")
	if key == "" {
		return nil, errors.New("не установлен API_KEY .env файле")
	}

	return &Config{
		Key: key,
	}, nil
}
