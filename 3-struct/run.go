package main

import (
	"go-course/bibin/api"
	"go-course/bibin/args"
	"go-course/bibin/config"
	"go-course/bibin/file"
	"go-course/bibin/handler"
	"go-course/bibin/storage"
)

func Run(args *args.Args, config *config.Config, storage storage.Storage) (string, error) {
	// Создаем (если не существует) локальный bin файл
	err := file.CreateBinFile(config)
	if err != nil {
		return "", err
	}

	// Инициализируем API
	binApi := api.NewBinApi(config)

	// Получаем результат
	return handler.CallHandler(args, binApi, storage)
}