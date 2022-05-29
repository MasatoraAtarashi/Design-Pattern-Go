package main

type Entry interface {
	GetName() string
	GetSize() int
	Add(entry Entry) error
	PrintList(prefix string)
	String() string
}
