package storage

import (
	"go-course/bibin/bins"
)

type Storage interface {
	SaveBin(bin bins.Bin) error
	ReadBins() (bins.BinList, error)
}
