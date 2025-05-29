package storage

import (
	"go-course/bibin/bins"
)

type Storage interface {
	SaveBin(bins.Bin) error
	ReadBins() (bins.BinList, error)
	DeleteBinById(string) error
}
