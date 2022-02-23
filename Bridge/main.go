package main

func main() {
	d1 := NewDisplayer(NewStringDisplayImpl("Hello, Japan."))
	d2 := NewCountDisplayer(NewStringDisplayImpl("Hello, World."))
	d1.Display()
	d2.Display()
	d2.MultiDisplay(5)
}
