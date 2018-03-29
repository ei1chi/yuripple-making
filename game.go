package main

import et "github.com/hajimehoshi/ebiten"

type gameState = int

const (
	gamePlaying gameState = iota
	gameFinished
)

type Game struct {
}

func (g *Game) Load() {
}

func (g *Game) Update(sc *et.Image) error {
	return nil
}
