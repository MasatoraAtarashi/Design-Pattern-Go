package main

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	errAddEntryToFile = errors.New("adding entry to file is prohibited")
)

type file struct {
	name string
	size int
}

func (f file) GetName() string {
	return f.name
}

func (f file) GetSize() int {
	return f.size
}

func (f file) Add(entry Entry) error {
	// FileにEntryを追加することはできない
	return errAddEntryToFile
}

func (f file) PrintList(prefix string) {
	fmt.Println(prefix + "/" + f.String())
}

func (f file) String() string {
	return f.GetName() + " (" + strconv.Itoa(f.GetSize()) + ")"
}

func NewFile(name string, size int) Entry {
	return file{
		name: name,
		size: size,
	}
}
