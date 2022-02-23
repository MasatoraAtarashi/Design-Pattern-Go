package main

import (
	"math/rand"
	"time"
)

type Strategy interface {
	NextHand() Hand
	Study(bool)
}

type winningStrategy struct {
	won      bool
	prevHand Hand
}

func NewWinningStrategy() Strategy {
	return winningStrategy{}
}

func (w winningStrategy) NextHand() Hand {
	if !w.won {
		w.prevHand = RandomHand()
	}
	return w.prevHand
}

func (w winningStrategy) Study(win bool) {
	w.won = win
}

type probStrategy struct {
	won         bool
	currentHand Hand
	prevHand    Hand
	history     [3][3]int
}

func NewProbStrategy() Strategy {
	return probStrategy{
		currentHand: GUU,
		prevHand:    GUU,
		history: [3][3]int{
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
		},
	}
}

func (p probStrategy) NextHand() Hand {
	rand.Seed(time.Now().UnixNano())
	bet := rand.Intn(p.getSum())
	var hand Hand
	if bet < p.history[p.currentHand][0] {
		hand = GUU
	} else if bet < (p.history[p.currentHand][0] + p.history[p.currentHand][1]) {
		hand = CHO
	} else {
		hand = PAA
	}
	p.prevHand = p.currentHand
	p.currentHand = hand
	return hand
}

func (p probStrategy) getSum() int {
	var sum int
	for i := 0; i < 3; i++ {
		sum += p.history[p.currentHand][i]
	}
	return sum
}

func (p probStrategy) Study(win bool) {
	if win {
		p.history[p.prevHand][p.currentHand]++
	} else {
		p.history[p.prevHand][(p.currentHand+1)%3]++
		p.history[p.prevHand][(p.currentHand+2)%3]++
	}
}

type randomStrategy struct {
	won      bool
	prevHand Hand
}

func NewRandomStrategy() Strategy {
	return randomStrategy{}
}

func (w randomStrategy) NextHand() Hand {
	return RandomHand()
}

func (w randomStrategy) Study(win bool) {
	w.won = win
}
