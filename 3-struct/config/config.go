package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ApiKey  string
	BinFile string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errors.New("не удалось найти .env файла")
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return nil, errors.New("не установлен API_KEY .env файле")
	}

	binFile := os.Getenv("BIN_FILE")
	if binFile == "" {
		return nil, errors.New("не установлен BIN_FILE .env файле")
	}

	return &Config{
		ApiKey:  apiKey,
		BinFile: binFile,
	}, nil
}
