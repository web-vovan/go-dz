package file

import (
	"encoding/json"
	"fmt"
	"go-course/bibin/config"
	"os"
)

// Чтение файл
func ReadFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать файл %s", filename)
	}

	return data, nil
}

// Запись в файл
func WriteFile(content []byte, filename string) error {
	if !IsExistsFile(filename) {
		return fmt.Errorf("нельзя записать в файл %s, который не существует", filename)
	}

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("не удалось открыть файл %s", filename)
	}

	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		return fmt.Errorf("не удалось записать в файл %s", filename)
	}

	return nil
}

// Проверка существования файла
func IsExistsFile(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}

	return true
}

// Проверка валидности json
func IsValidJson(data *[]byte) bool {
	var jsonData interface{}
	if err := json.Unmarshal(*data, &jsonData); err != nil {
		return false
	}
	return true
}

// Создание файла
func CreateFile(filename string) error {
	file, err := os.Create(filename)
	file.Close()
	if err != nil {
		return fmt.Errorf("не удалось создать файл %s", filename)
	}

	return nil
}

// Удаление файла
func DeleteFile(filename string) error {
	if !IsExistsFile(filename) {
		return nil
	}

	err := os.Remove(filename)
	if err != nil {
		return err
	}

	return nil
}

// Создание bin файла
func CreateBinFile(config *config.Config) error {
	filepath := config.BinFile

	if !IsExistsFile(filepath) {
		err := CreateFile(filepath)
		if err != nil {
			return fmt.Errorf("не удалось создать bin файл %s", filepath)
		}
	}

	return nil
}

// Удаление файла с бинами
func DeleteBinFile(config *config.Config) error {
	filepath := config.BinFile

	err := DeleteFile(filepath)
	if err != nil {
		return err
	}

	return nil
}
