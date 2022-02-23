package main

import (
	"fmt"
)

func main() {
	p1 := NewPlayer("Taro", NewWinningStrategy())
	p2 := NewPlayer("Hana", NewProbStrategy())
	for i := 0; i < 10000; i++ {
		nextHand1 := p1.NextHand()
		nextHand2 := p2.NextHand()
		if nextHand1.isStrongerThan(nextHand2) {
			fmt.Printf("winner: %s\n", p1.name)
			p1.Win()
			p2.Lose()
		} else if nextHand1.isWeakerThan(nextHand2) {
			fmt.Printf("winner: %s\n", p2.name)
			p1.Lose()
			p2.Win()
		} else {
			fmt.Println("even")
			p1.Even()
			p2.Even()
		}
	}
	fmt.Println("Total result:")
	fmt.Println(p1.String())
	fmt.Println(p2.String())
}
