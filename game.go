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
	gameResult
)

type Game struct {
	state     td.Stm
	atlas     *td.Atlas
	sprites   map[string]*td.Sprite
	charas    []*Chara
	catched   *Chara
	offset    complex128
	fullGauge *et.Image

	t struct {
		mode, score, time *td.TextBox
	}
	r struct {
		gauge td.Rect
	}
}

func (g *Game) Load() {

	var err error

	g.atlas, err = td.NewAtlas("resources/atlas")
	if err != nil {
		log.Fatal(err)
	}

	g.sprites = map[string]*td.Sprite{}
	func(s []string) {
		for _, name := range s {
			path := name + ".png"
			g.sprites[name], err = g.atlas.NewSprite(path)
			if err != nil {
				log.Fatal(err)
			}
		}
	}([]string{
		"neco",
		"nonke",
		"tachi",
		"riba_tachi",
		"riba_neco",
		"heart",
	})

	ui := td.Rect{5, 5, screenW - 5, screenH}.HSplit(30, 40, 40)
	g.t.mode = td.NewTextBox(ui[0], mplus[20], 5, "")
	g.t.score = td.NewTextBox(ui[1], mplus[20], 5, "score")

	time := ui[2].WithMargin(0, 10, 0, 10).VSplit(80)
	g.t.time = td.NewTextBox(time[0], mplus[20], 5, "time")
	g.r.gauge = time[1].WithMargin(0, 0, 30, 0)

	w, h := g.r.gauge.Width(), g.r.gauge.Height()
	g.fullGauge, _ = et.NewImage(int(w), int(h), et.FilterDefault)
	g.fullGauge.Fill(color.Black)
}

func (g *Game) Update(sc *et.Image) error {

	g.state.Update()
	g.processCharas()
	g.sweepAll()

	switch g.state.Get() {
	case gamePlaying:
		//g.collisionAll()
	}

	g.drawAll(sc)

	// 終了判定
	if g.state.Get() == gamePlaying {
		if g.state.Elapsed() >= levels[easyLevel].limit {
			g.state.Transition(gameResult)
		}
	}

	if g.state.Get() == gameResult {
		if g.state.Elapsed() > 40 && td.IsJustPressed {
			return ErrSuccess
		}
	}

	return nil
}

func (g *Game) sweepAll() {
	next := g.charas[:0]
	for _, c := range g.charas {
		if c != nil {
			if !c.isDead {
				next = append(next, c)
			} else {
				fmt.Println("sweepd")
			}
		}
	}
	g.charas = next
}
