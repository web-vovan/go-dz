package main

import (
	"fmt"
	"go-course/bibin/bins"
	"go-course/bibin/config"
	"go-course/bibin/file"
	"go-course/bibin/storage"
)

const BIN_FILE = "bins.json"

func main() {
	if !file.IsExistsFile(BIN_FILE) {
		err := file.CreateFile(BIN_FILE)
		if err != nil {
			fmt.Println(err)
		}
	}

	_, err := config.NewConfig()

	if err != nil {
		fmt.Println(err)
		return
	}

	storage := storage.NewFileStorage()

	bin := bins.NewBin("uuid1", false, "one bin")

	err = storage.SaveBin(bin)
	if err != nil {
		fmt.Println(err)
	}

	binList, err := storage.ReadBins()
	if err != nil {
		fmt.Println(binList)
	}

	fmt.Println(binList)
}
