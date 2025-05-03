package bins

import "time"

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

type BinList struct {
	Bins []Bin `json:"bins"`
}

func NewBin(id string, private bool, name string) Bin {
	return Bin{
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
}

func NewBinList(bins []Bin) BinList {
	return BinList{
		Bins: bins,
	}
}

func (binList *BinList) AddBin(bin Bin) {
	binList.Bins = append(binList.Bins, bin)
}
