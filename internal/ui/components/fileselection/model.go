package fileselection

import "os"

type Model struct {
	Files      []os.DirEntry
	FileCursor int
}

func New(available []os.DirEntry) Model {
	return Model{
		Files:      available,
		FileCursor: 0,
	}
}
