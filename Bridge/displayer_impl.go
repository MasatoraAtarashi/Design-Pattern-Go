package main

import "fmt"

type DisplayerImpl interface {
	RawOpen()
	RawPrint()
	RawClose()
}

type stringDisplayImpl struct {
	value string
	width int
}

func NewStringDisplayImpl(value string) DisplayerImpl {
	return stringDisplayImpl{
		value: value,
		width: len(value),
	}
}

func (s stringDisplayImpl) RawOpen() {
	s.printLine()
}

func (s stringDisplayImpl) RawPrint() {
	fmt.Printf("|%s|\n", s.value)
}

func (s stringDisplayImpl) RawClose() {
	s.printLine()
}

func (s stringDisplayImpl) printLine() {
	fmt.Printf("+")
	for i := 0; i < s.width; i++ {
		fmt.Printf("-")
	}
	fmt.Println("+")
}
