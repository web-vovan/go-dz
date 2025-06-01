package storage

import (
	"encoding/json"
	"go-course/bibin/bins"
	"go-course/bibin/config"
	"go-course/bibin/file"
)

type FileStorage struct {
	filePath string
}

func NewFileStorage(config *config.Config) Storage {
	return &FileStorage{
		filePath: config.BinFile,
	}
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

	err = file.WriteFile(content, f.filePath)
	if err != nil {
		return err
	}

	return nil
}

func (f *FileStorage) ReadBins() (bins.BinList, error) {
	var binList bins.BinList

	data, err := file.ReadFile(f.filePath)
	if err != nil {
		return binList, err
	}

	json.Unmarshal(data, &binList)

	return binList, nil
}

func (f *FileStorage) DeleteBinById(id string) error {
	binList, err := f.ReadBins()
	if err != nil {
		return err
	}

	for i, v := range binList.Bins {
		if v.Id == id {
			binList.Bins = append(binList.Bins[:i], binList.Bins[i+1:]...)
		}
	}

	content, err := json.Marshal(binList)
	if err != nil {
		return err
	}

	err = file.WriteFile(content, f.filePath)
	if err != nil {
		return err
	}

	return nil
}