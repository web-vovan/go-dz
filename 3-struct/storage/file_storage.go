package storage

import (
	"encoding/json"
	"go-course/bibin/bins"
	"go-course/bibin/file"
)

type FileStorage struct {}

func NewFileStorage() Storage {
    return &FileStorage{}
}

func (f *FileStorage) SaveBin(bin bins.Bin) error {
	binList, err := f.ReadBins()
	if err != nil {
		return err
	}

	binList.AddBin(bin)
	content, err := json.Marshal(binList)
	if err != nil {
		return err
	}

	err = file.WriteFile(content, "bins.json")
	if err != nil {
		return err
	}

	return nil
}

func (f *FileStorage) ReadBins() (bins.BinList, error) {
	var binList bins.BinList

	data, err := file.ReadFile("bins.json")
	if err != nil {
		return binList, err
	}

	json.Unmarshal(data, &binList)

	return binList, nil
}