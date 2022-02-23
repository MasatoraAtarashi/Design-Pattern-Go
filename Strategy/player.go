package main

import "fmt"

type Player struct {
	name     string
	strategy Strategy
	gameCnt  int
	winCnt   int
	loseCnt  int
}

func NewPlayer(name string, strategy Strategy) Player {
	return Player{
		name:     name,
		strategy: strategy,
	}
}

func (p Player) String() string {
	return fmt.Sprintf("[%s: %dgames, %dwin, %dlose]", p.name, p.gameCnt, p.winCnt, p.loseCnt)
}

func (p Player) NextHand() Hand {
	return p.strategy.NextHand()
}

func (p *Player) Win() {
	p.strategy.Study(true)
	p.gameCnt++
	p.winCnt++
}

func (p *Player) Lose() {
	p.strategy.Study(false)
	p.gameCnt++
	p.loseCnt++
}

func (p *Player) Even() {
	p.strategy.Study(false)
	p.gameCnt++
}
