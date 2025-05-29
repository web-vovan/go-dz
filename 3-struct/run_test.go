package main

import (
	// "encoding/json"
	"fmt"
	"go-course/bibin/args"
	"go-course/bibin/config"
	"go-course/bibin/file"
	"go-course/bibin/handler"
	"go-course/bibin/storage"
	"os"
	"testing"

    "github.com/stretchr/testify/assert"
)

const TEST_JSON_FILE = "test.json"
const TEST_JSON_FILE_DATA = "{\"sample\": \"Hello World\"}"

// Создание тестового json файла
func createTestJsonFile() error {
	err := os.WriteFile(TEST_JSON_FILE, []byte(TEST_JSON_FILE_DATA), 0644)
	if err != nil {
		return fmt.Errorf("не удалось создать тестовый json файл: %v", err)
	}
	return nil
}

func TestRun_CreateBin_Success(t *testing.T) {
	// Инициализируем конфиг
	config, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Инициализируем хранилище
	storage := storage.NewFileStorage(config)

	// Удаляем файл с бинами
	err = file.DeleteBinFile(config)
	if err != nil {
		t.Error(err)
	}

	// Создаем тестовый json файл
	err = createTestJsonFile()
	if err != nil {
		t.Error(err)
	}

	binName := "new bin"

	args := &args.Args{
		Create: true,
		File:   TEST_JSON_FILE,
		Name:   binName,
	}

	output, err := Run(args, config, storage)
	if err != nil {
		t.Error(err)
	}

	if output != handler.SUCCESS_CREATE_MSG {
		t.Error("ошибка при создании bin файла")
	}

	bins, err := storage.ReadBins()
	if err != nil {
		t.Errorf("ошибка при чтении файла с бинами: %v", err)
	}

	if len(bins.Bins) != 1 {
		t.Errorf("в файле с бинами ожидался 1 бин, а по факту %d бинов", len(bins.Bins))
	}

	if bins.Bins[0].Name != binName {
		t.Errorf("неверное имя бина в файле, expected: %s, actual: %s", binName, bins.Bins[0].Name)
	}
}

func TestRun_CreateBin_Fail(t *testing.T) {
	// Инициализируем конфиг
	config, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Инициализируем хранилище
	storage := storage.NewFileStorage(config)

	// Создаем тестовый json файл
	err = createTestJsonFile()
	if err != nil {
		t.Error(err)
	}

	binName := "new bin"

	testCases := []struct {
		args     args.Args
		name     string
		expected string
	}{
		{
			args: args.Args{
				Create: true,
				Name:   binName,
			},
			name:     "не передан файл",
			expected: handler.ERROR_FILE_MSG,
		},
		{
			args: args.Args{
				Create: true,
				File:   TEST_JSON_FILE,
			},
			name:     "не передано имя",
			expected: handler.ERROR_NAME_MSG,
		},
	}

    for _, tc := range testCases {
        t.Run(tc.name, func (t *testing.T) {
            _, err := Run(&tc.args, config, storage)
            if err.Error() != tc.expected {
                t.Errorf("неверный текст ошибки, expected: %s, actual: %s", tc.expected, err.Error())
            }
        })
    }
}

func TestRun_UpdateBin_Success(t *testing.T) {
	// Инициализируем конфиг
	config, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Инициализируем хранилище
	storage := storage.NewFileStorage(config)

	// Удаляем файл с бинами
	err = file.DeleteBinFile(config)
	if err != nil {
		t.Error(err)
	}

	// Создаем тестовый json файл
	err = createTestJsonFile()
	if err != nil {
		t.Error(err)
	}

	binName := "new bin"

	argsCreate := &args.Args{
		Create: true,
		File:   TEST_JSON_FILE,
		Name:   binName,
	}

    // Создаем бин
	outputCreate, err := Run(argsCreate, config, storage)
	if err != nil {
		t.Error(err)
	}

	if outputCreate != handler.SUCCESS_CREATE_MSG {
		t.Error("ошибка при создании bin файла")
	}

	bins, err := storage.ReadBins()

	if err != nil {
		t.Errorf("ошибка при чтении файла с бинами: %v", err)
	}

	if len(bins.Bins) != 1 {
		t.Errorf("в файле с бинами ожидался 1 бин, а по факту %d бинов", len(bins.Bins))
	}

	if bins.Bins[0].Name != binName {
		t.Errorf("неверное имя бина в файле, expected: %s, actual: %s", binName, bins.Bins[0].Name)
	}

    argsUpdate := &args.Args{
		Update: true,
		File:   TEST_JSON_FILE,
		Id:   bins.Bins[0].Id,
	}

    // Обновляем бин
	outputUpdate, err := Run(argsUpdate, config, storage)
	if err != nil {
		t.Error(err)
	}

    if outputUpdate != handler.SUCCESS_UPDATE_MSG {
		t.Error("ошибка при создании bin файла")
	}
}

func TestRun_UpdateBin_Fail(t *testing.T) {
	// Инициализируем конфиг
	config, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Инициализируем хранилище
	storage := storage.NewFileStorage(config)

    testCases := []struct {
		args     args.Args
		name     string
		expected string
	}{
		{
			args: args.Args{
				Update: true,
				Id:   "123",
			},
			name:     "не передан файл",
			expected: handler.ERROR_FILE_MSG,
		},
		{
			args: args.Args{
				Update: true,
				File:   TEST_JSON_FILE,
			},
			name:     "не передан id",
			expected: handler.ERROR_ID_MSG,
		},
	}

    for _, tc := range testCases {
        t.Run(tc.name, func (t *testing.T) {
            _, err := Run(&tc.args, config, storage)
            if err.Error() != tc.expected {
                t.Errorf("неверный текст ошибки, expected: %s, actual: %s", tc.expected, err.Error())
            }
        })
    }
}

func TestRun_GetBin_Success(t *testing.T) {
	// Инициализируем конфиг
	config, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Инициализируем хранилище
	storage := storage.NewFileStorage(config)

	// Удаляем файл с бинами
	err = file.DeleteBinFile(config)
	if err != nil {
		t.Error(err)
	}

	// Создаем тестовый json файл
	err = createTestJsonFile()
	if err != nil {
		t.Error(err)
	}

	binName := "new bin"

	argsCreate := &args.Args{
		Create: true,
		File:   TEST_JSON_FILE,
		Name:   binName,
	}

    // Создаем бин
	outputCreate, err := Run(argsCreate, config, storage)
	if err != nil {
		t.Error(err)
	}

	if outputCreate != handler.SUCCESS_CREATE_MSG {
		t.Error("ошибка при создании bin файла")
	}

	bins, err := storage.ReadBins()

	if err != nil {
		t.Errorf("ошибка при чтении файла с бинами: %v", err)
	}

	if len(bins.Bins) != 1 {
		t.Errorf("в файле с бинами ожидался 1 бин, а по факту %d бинов", len(bins.Bins))
	}

	if bins.Bins[0].Name != binName {
		t.Errorf("неверное имя бина в файле, expected: %s, actual: %s", binName, bins.Bins[0].Name)
	}

    argsGet := &args.Args{
		Get: true,
		Id:   bins.Bins[0].Id,
	}

    // Получаем бин
	outputGet, err := Run(argsGet, config, storage)
	if err != nil {
		t.Error(err)
	}

    assert.JSONEq(t, outputGet, TEST_JSON_FILE_DATA)
}

func TestRun_GetBin_Fail(t *testing.T) {
	// Инициализируем конфиг
	config, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Инициализируем хранилище
	storage := storage.NewFileStorage(config)

	testCases := []struct {
		args     args.Args
		name     string
		expected string
	}{
		{
			args: args.Args{
				Get: true,
			},
			name:     "не передан id",
			expected: handler.ERROR_ID_MSG,
		},
	}

    for _, tc := range testCases {
        t.Run(tc.name, func (t *testing.T) {
            _, err := Run(&tc.args, config, storage)
            if err.Error() != tc.expected {
                t.Errorf("неверный текст ошибки, expected: %s, actual: %s", tc.expected, err.Error())
            }
        })
    }
}

func TestRun_DeleteBin_Success(t *testing.T) {
	// Инициализируем конфиг
	config, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Инициализируем хранилище
	storage := storage.NewFileStorage(config)

	// Удаляем файл с бинами
	err = file.DeleteBinFile(config)
	if err != nil {
		t.Error(err)
	}

	// Создаем тестовый json файл
	err = createTestJsonFile()
	if err != nil {
		t.Error(err)
	}

	binName := "new bin"

	argsCreate := &args.Args{
		Create: true,
		File:   TEST_JSON_FILE,
		Name:   binName,
	}

    // Создаем бин
	outputCreate, err := Run(argsCreate, config, storage)
	if err != nil {
		t.Error(err)
	}

	if outputCreate != handler.SUCCESS_CREATE_MSG {
		t.Error("ошибка при создании bin файла")
	}

	bins, err := storage.ReadBins()
	if err != nil {
		t.Errorf("ошибка при чтении файла с бинами: %v", err)
	}

	if len(bins.Bins) != 1 {
		t.Errorf("в файле с бинами ожидался 1 бин, а по факту %d бинов", len(bins.Bins))
	}

	if bins.Bins[0].Name != binName {
		t.Errorf("неверное имя бина в файле, expected: %s, actual: %s", binName, bins.Bins[0].Name)
	}

    argsDelete := &args.Args{
		Delete: true,
		Id:   bins.Bins[0].Id,
	}

    // Удаляем бин
	outputDelete, err := Run(argsDelete, config, storage)
	if err != nil {
		t.Error(err)
	}

    if outputDelete != handler.SUCCESS_DELETE_MSG {
        t.Errorf("ошибка при удалении бина, expected: %s, actual: %s", handler.SUCCESS_DELETE_MSG, outputDelete)
    }
    
    bins, err = storage.ReadBins()
    if err != nil {
        t.Errorf("ошибка при чтении файла с бинами: %v", err)
    }

    if len(bins.Bins) != 0 {
		t.Errorf("в файле с бинами ожидалось 0 бинов, а по факту %d бинов", len(bins.Bins))
	}
}

func TestRun_DeleteBin_Fail(t *testing.T) {
	// Инициализируем конфиг
	config, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Инициализируем хранилище
	storage := storage.NewFileStorage(config)

	testCases := []struct {
		args     args.Args
		name     string
		expected string
	}{
		{
			args: args.Args{
				Delete: true,
			},
			name:     "не передан id",
			expected: handler.ERROR_ID_MSG,
		},
	}

    for _, tc := range testCases {
        t.Run(tc.name, func (t *testing.T) {
            _, err := Run(&tc.args, config, storage)
            if err.Error() != tc.expected {
                t.Errorf("неверный текст ошибки, expected: %s, actual: %s", tc.expected, err.Error())
            }
        })
    }
}

func TestRun_ListBin_Success(t *testing.T) {
	// Инициализируем конфиг
	config, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Инициализируем хранилище
	storage := storage.NewFileStorage(config)

	// Удаляем файл с бинами
	err = file.DeleteBinFile(config)
	if err != nil {
		t.Error(err)
	}

	// Создаем тестовый json файл
	err = createTestJsonFile()
	if err != nil {
		t.Error(err)
	}

	binName := "new bin"

	argsCreate := &args.Args{
		Create: true,
		File:   TEST_JSON_FILE,
		Name:   binName,
	}

    // Создаем бин
	outputCreate, err := Run(argsCreate, config, storage)
	if err != nil {
		t.Error(err)
	}

	if outputCreate != handler.SUCCESS_CREATE_MSG {
		t.Error("ошибка при создании bin файла")
	}

	bins, err := storage.ReadBins()
	if err != nil {
		t.Errorf("ошибка при чтении файла с бинами: %v", err)
	}

	if len(bins.Bins) != 1 {
		t.Errorf("в файле с бинами ожидался 1 бин, а по факту %d бинов", len(bins.Bins))
	}

	if bins.Bins[0].Name != binName {
		t.Errorf("неверное имя бина в файле, expected: %s, actual: %s", binName, bins.Bins[0].Name)
	}
}
