package args

import "flag"

type Args struct {
	Create bool
	Update bool
	Delete bool
	Get    bool
	List   bool
	File   string
	Name   string
	Id     string
}

func NewArgs() *Args {
	create := flag.Bool("create", false, "")
	update := flag.Bool("update", false, "")
	delete := flag.Bool("delete", false, "")
	get := flag.Bool("get", false, "")
	list := flag.Bool("list", false, "")
	file := flag.String("file", "", "")
	name := flag.String("name", "", "")
	id := flag.String("id", "", "")

	flag.Parse()

	return &Args{
		Create: *create,
		Update: *update,
		Delete: *delete,
		Get:    *get,
		List:   *list,
		File:   *file,
		Name:   *name,
		Id:     *id,
	}
}