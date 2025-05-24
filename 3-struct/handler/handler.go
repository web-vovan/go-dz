package handler

import (
	"errors"
	"fmt"

	"go-course/bibin/api"
	"go-course/bibin/args"
	"go-course/bibin/bins"
	"go-course/bibin/file"
	"go-course/bibin/storage"
)

const SUCCESS_CREATE_MSG = "bin успешно создан"
const SUCCESS_UPDATE_MSG = "bin успешно обновлен"
const SUCCESS_DELETE_MSG = "bin успешно удален"

const ERROR_FILE_MSG = "не указан обязательный параметр file"
const ERROR_ID_MSG = "не указан обязательный параметр id"
const ERROR_NAME_MSG = "не указан обязательный параметр name"

func CreateBin(args *args.Args, binApi *api.BinApi, storage storage.Storage) (string, error) {
	if args.File == "" {
		return "", errors.New(ERROR_FILE_MSG)
	}

	if args.Name == "" {
		return "", errors.New(ERROR_NAME_MSG)
	}

	fileData, err := file.ReadFile(args.File)
	if err != nil {
		return "", err
	}

	if r := file.IsValidJson(&fileData); !r {
		return "", fmt.Errorf("файл %s содержит невалидный json", args.File)
	}

	response, err := binApi.Create(&fileData)
	if err != nil {
		return "", err
	}

	bin := bins.NewBin(
		response.Metadata.Id,
		response.Metadata.Private,
		args.Name,
	)

	err = storage.SaveBin(bin)
	if err != nil {
		return "", err
	}

	return SUCCESS_CREATE_MSG, err
}

func GetBin(args *args.Args, binApi *api.BinApi) (string, error) {
	if args.Id == "" {
		return "", errors.New(ERROR_ID_MSG)
	}
	return binApi.Get(args.Id)
}

func UpdateBin(args *args.Args, binApi *api.BinApi) (string, error) {
	if args.File == "" {
		return "", errors.New(ERROR_FILE_MSG)
	}

	if args.Id == "" {
		return "", errors.New(ERROR_ID_MSG)
	}

	fileData, err := file.ReadFile(args.File)
	if err != nil {
		return "", err
	}

	if r := file.IsValidJson(&fileData); !r {
		return "", fmt.Errorf("файл %s содержит невалидный json", args.File)
	}

	err = binApi.Update(args.Id, &fileData)
	if err != nil {
		return "", err
	}

	return SUCCESS_UPDATE_MSG, nil
}

func ListBin(storage storage.Storage) (string, error) {
	var result string

	binList, err := storage.ReadBins()
	if err != nil {
		return "", err
	}

	for _, v := range binList.Bins {
		result += v.Id + ": " + v.Name + "\n"
	}

	return result, nil
}

func DeleteBin(args *args.Args, binApi *api.BinApi, storage storage.Storage) (string, error) {
	if args.Id == "" {
		return "", errors.New(ERROR_ID_MSG)
	}

	err := binApi.Delete(args.Id)
	if err != nil {
		return "", err
	}

	err = storage.DeleteBinById(args.Id)
	if err != nil {
		return "", err
	}

	return SUCCESS_DELETE_MSG, nil
}

func CallHandler(args *args.Args, binApi *api.BinApi, storage storage.Storage) (string, error) {
	switch {
	case args.Create:
		return CreateBin(args, binApi, storage)
	case args.Update:
		return UpdateBin(args, binApi)
	case args.Get:
		return GetBin(args, binApi)
	case args.List:
		return ListBin(storage)
	case args.Delete:
		return DeleteBin(args, binApi, storage)
	default:
		return "", errors.New("указанная операция не поддерживается")
	}
}
