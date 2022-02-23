package main

import (
	"math/rand"
	"time"
)

type Hand int

const (
	GUU Hand = iota
	CHO
	PAA
)

var _HandValueMap = map[Hand]string{
	GUU: "グー",
	CHO: "チョキ",
	PAA: "パー",
}

func RandomHand() Hand {
	rand.Seed(time.Now().UnixNano())
	hn := rand.Intn(3)
	return Hand(hn)
}

func (h Hand) String() string {
	return _HandValueMap[h]
}

func (h Hand) isStrongerThan(hand Hand) bool {
	return h.fight(hand) == 1
}

func (h Hand) isWeakerThan(hand Hand) bool {
	return h.fight(hand) == -1
}

func (h Hand) fight(hand Hand) int {
	if h == hand {
		return 0
	} else if (h+1)%3 == hand {
		return 1
	} else {
		return -1
	}
}
