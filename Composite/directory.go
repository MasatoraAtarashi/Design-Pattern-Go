package main

import (
	"fmt"
	"strconv"
)

type directory struct {
	name    string
	entries []Entry
}

func (d *directory) GetName() string {
	return d.name
}

func (d *directory) GetSize() int {
	var size int
	for _, entry := range d.entries {
		size += entry.GetSize()
	}
	return size
}

func (d *directory) Add(entry Entry) error {
	d.entries = append(d.entries, entry)
	return nil
}

func (d *directory) PrintList(prefix string) {
	// TODO: イテレータパターンを採用する
	fmt.Println(prefix + "/" + d.String())
	for _, entry := range d.entries {
		entry.PrintList(prefix + "/" + d.name)
	}
}

func (d *directory) String() string {
	return d.GetName() + " (" + strconv.Itoa(d.GetSize()) + ")"
}

func NewDirectory(name string) Entry {
	return &directory{
		name: name,
	}
}
