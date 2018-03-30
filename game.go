package main

import (
	"log"

	td "github.com/ei1chi/tendon"
	et "github.com/hajimehoshi/ebiten"
	"golang.org/x/image/font"
)

type gameState = int

const (
	gamePlaying gameState = iota
	gameFinished
)

type Game struct {
	state   td.Stm
	atlas   *td.Atlas
	mplus24 font.Face
	//charas  []*Chara
}

func (g *Game) Load() {

	var err error

	g.atlas, err = td.NewAtlas("resources/atlas")
	if err != nil {
		log.Fatal(err)
	}

	//func(s []string) {
	//	for _, path := range s {
	//		path += ".png"
	//		g.sprites[s], err = g.atlas.NewSprite(path)
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//	}
	//}([]string{
	//	"neco",
	//	"nonke",
	//})
}

func (g *Game) Update(sc *et.Image) error {
	return nil
}
