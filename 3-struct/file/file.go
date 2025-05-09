package file

import (
	"fmt"
	"os"
)

func ReadFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать файл %s", filename)
	}

	return data, nil
}

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

func IsExistsFile(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}

	return true
}

func CreateFile(filename string) error {
	file, err := os.Create(filename)
	file.Close()
	if err != nil {
		return fmt.Errorf("не удалось создать файл %s", filename)
	}

	return nil
}
