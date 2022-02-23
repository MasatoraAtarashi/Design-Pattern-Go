package main

import "fmt"

// NOTE: もっと良い実現方法がある気がする？

type Printer interface {
	Open()
	Print()
	Close()
}

type Displayer interface {
	Display()
}

type display struct {
	printer Printer
}

func (d display) Display() {
	d.printer.Open()
	for i := 0; i < 5; i++ {
		d.printer.Print()
	}
	d.printer.Close()
}

type RuneDisplay struct {
	display
	rune []rune
}

func NewRuneDisplay(r []rune) Displayer {
	runeDisplay := RuneDisplay{
		display: display{},
		rune:    r,
	}
	runeDisplay.display.printer = runeDisplay
	return runeDisplay
}

func (c RuneDisplay) Open() {
	fmt.Print("<<")
}

func (c RuneDisplay) Print() {
	fmt.Printf("%s", string(c.rune))
}

func (c RuneDisplay) Close() {
	fmt.Println(">>")
}

type StringDisplay struct {
	display
	value string
}

func NewStringDisplay(value string) Displayer {
	StringDisplay := StringDisplay{
		display: display{},
		value:   value,
	}
	StringDisplay.display.printer = StringDisplay
	return StringDisplay
}

func (s StringDisplay) Open() {
	s.printLine()
}

func (s StringDisplay) Print() {
	fmt.Printf("|%s|\n", s.value)
}

func (s StringDisplay) Close() {
	s.printLine()
}

func (s StringDisplay) printLine() {
	fmt.Print("+")
	for i := 0; i < len(s.value); i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}
