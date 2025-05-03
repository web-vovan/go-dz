package file

import (
	"fmt"
	"os"
	"path/filepath"
)

func ReadFile(filename string) ([]byte, error) {
	ext := filepath.Ext(filename)

	if ext != ".json" {
		return nil, fmt.Errorf("файл %s имеет не правильное расширение %s, ожидается .json", filename, ext)
	}

	data, err := os.ReadFile(filename)

	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать файл %s", filename)
	}

	return data, nil
}

func WriteFile(content []byte, filename string) error {
	if !IsExistsFile(filename) {
		return fmt.Errorf("файл %s не существует", filename)
	}

	file, err := os.Create(filename)

	if err != nil {
		return fmt.Errorf("ошибка при записи в файл %s", filename)
	}

	defer file.Close()

	_, err = file.Write(content)

	if err != nil {
		return err
	}

	return nil
}

func IsExistsFile(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}

	return true
}

func CreateBinFile(filename string) error {
	fileExist := IsExistsFile(filename)

	if !fileExist {
		file, err := os.Create(filename)
		file.Close()

		if err != nil {
			return fmt.Errorf("не удалось создать файл %s", filename)
		}
	}

	return nil
}
