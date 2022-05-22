package main

import (
	"facade/welcome"
)

func main() {
	pageMaker := welcome.NewPageMaker()
	pageMaker.Make("hyuki@example.com", "./welcome/maildata.json")
}
