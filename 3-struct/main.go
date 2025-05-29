package main

import (
	"fmt"
	"go-course/bibin/args"
	"go-course/bibin/config"
	"go-course/bibin/storage"
)

func main() {
	// Аргументы из командной строки
	args := args.NewArgs()

	// Загружаем конфиг
	config, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Инициализируем хранилище
	storage := storage.NewFileStorage(config)

	// Получаем результат
	result, err := Run(args, config, storage)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
