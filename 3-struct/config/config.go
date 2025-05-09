package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic("Не удалось найти .env файла")
	}

	key := os.Getenv("API_KEY")
	if key == "" {
		panic("Не установлен API_KEY .env файле")
	}

	return &Config{
		Key: key,
	}
}
