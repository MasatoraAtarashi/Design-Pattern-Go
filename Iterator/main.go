package main

import "fmt"

type Book struct {
	Name string
}

func main() {
	bs := []Book{
		{"Around the World in 80 Days"},
		{"Bible"},
		{"Cinderella"},
		{"Daddy-Long-Legs"},
	}
	bookShelf := BookShelf{Books: bs}
	it := bookShelf.Iterator()
	for it.HasNext() {
		b := it.Next().(Book)
		fmt.Println(b.Name)
	}
}
