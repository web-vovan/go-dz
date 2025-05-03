package storage

import (
	"encoding/json"
	"go-course/bibin/bins"
	"go-course/bibin/file"
)

func SaveBin(bin bins.Bin) error {
	binList, err := ReadBins()

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

func ReadBins() (bins.BinList, error) {
	data, err := file.ReadFile("bins.json")
	var binList bins.BinList

	if err != nil {
		return binList, err
	}

	json.Unmarshal(data, &binList)

	return binList, nil
}
