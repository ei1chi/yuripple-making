package main

import (
	"fmt"
	"image/color"
	"log"

	td "github.com/ei1chi/tendon"
	et "github.com/hajimehoshi/ebiten"
)

type gameState = int

const (
	gamePlaying gameState = iota
	gameFinished
)

type Game struct {
	state td.Stm
	atlas *td.Atlas
	//charas  []*Chara

	rs struct {
		mode, score td.Rect
		time        []td.Rect
	}
	ts struct {
		mode, time, score *td.Text
	}
}

func (g *Game) Load() {

	var err error

	g.atlas, err = td.NewAtlas("resources/atlas")
	if err != nil {
		log.Fatal(err)
	}

	uir := td.Rect{5, 5, screenW - 5, screenH}.HSplit(30, 40, 40)
	g.rs.mode = uir[0]
	g.rs.score = uir[1].WithMargin(0, 10, 0, 10)
	g.rs.time = uir[2].WithMargin(0, 15, 0, 15).VSplit(80)

	g.ts.mode = td.NewText(root.mplus, 15, 5, "MODE: EASY")
	g.ts.score = td.NewText(root.mplus, 20, 8, "score")
	g.ts.time = td.NewText(root.mplus, 20, 8, "time")

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
	g.state.Update()
	g.ts.mode.DrawR(sc, g.rs.mode, color.Black)
	g.ts.score.SetText(fmt.Sprintf("score: %d", g.state.Elapsed()))
	g.ts.score.DrawR(sc, g.rs.score, color.Black)
	g.ts.time.DrawR(sc, g.rs.time[0], color.Black)

	w := g.rs.time[1].Width() * (1.0 - float64(g.state.Elapsed())/1000)
	h := g.rs.time[1].Height()
	gauge, _ := et.NewImage(int(w), int(h), et.FilterDefault)
	gauge.Fill(color.Black)
	op := &et.DrawImageOptions{}
	op.GeoM.Translate(g.rs.time[1].Left, g.rs.time[1].Top)
	sc.DrawImage(gauge, op)
	return nil
}
