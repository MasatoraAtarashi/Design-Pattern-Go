package main

import "fmt"

type Banner struct {
	value string
}

func (b Banner) ShowWithParen() {
	fmt.Printf("(%s)\n", b.value)
}

func (b Banner) ShowWithAster() {
	fmt.Printf("*%s*\n", b.value)
}
