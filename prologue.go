package main

import (
	"image/color"
	"math"

	td "github.com/ei1chi/tendon"
	et "github.com/hajimehoshi/ebiten"
)

type Prologue struct {
	title, lead *td.TextBox
	state       td.Stm
}

func (p *Prologue) Load() {
	rect := td.Rect{0, 0, screenW, screenH}.WithMargin(0, 100, 0, 420)
	p.title = td.NewTextBox(rect, root.mp45, 5, "タイトルロゴ")

	rect = rect.SnapOutside(8, screenW, 40)
	p.lead = td.NewTextBox(rect, root.mp20, 5, "tap to start")
}

func (p *Prologue) Update(sc *et.Image) error {

	p.title.Draw(sc, 0, 0, color.Black)

	p.state.Update()
	if p.state.Elapsed() > 40 && td.IsPressed {
		return ErrSuccess
	}
	a := math.Cos(float64(p.state.Elapsed())/60)*128 + 128
	p.lead.Draw(sc, 0, 0, color.RGBA{0, 0, 0, uint8(a)})

	return nil
}
