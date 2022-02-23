package main

func main() {
	d1 := NewRuneDisplay([]rune("H"))
	d2 := NewStringDisplay("Hello, World.")
	d3 := NewStringDisplay("こんにちは。")
	d1.Display()
	d2.Display()
	d3.Display()
}
