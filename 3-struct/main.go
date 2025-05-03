package main

import (
	"fmt"
	"go-course/bibin/bins"
	"go-course/bibin/file"
	"go-course/bibin/storage"
)

func main() {
	err := file.CreateBinFile("bins.json")

	if err != nil {
		fmt.Println(err)
	}

	bin := bins.NewBin("uuid1", false, "one bin")

	err = storage.SaveBin(bin)

	if err != nil {
		fmt.Println(err)
	}
}
