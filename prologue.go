package main

import (
	"image/color"
	"math"

	td "github.com/ei1chi/tendon"
	et "github.com/hajimehoshi/ebiten"
)

type Prologue struct {
	title, lead *td.Text
	state       td.Stm
}

func (p *Prologue) Load() {
	p.title = td.NewText(root.mplus, 45, 5, "百合っぷるメイキング")
	p.lead = td.NewText(root.mplus, 20, 5, "tap to start")
}

func (p *Prologue) Update(sc *et.Image) error {

	rect := td.Rect{0, 0, screenW, screenH}.WithMargin(0, 100, 0, 420)
	p.title.DrawR(sc, rect, color.Black)

	p.state.Update()
	if p.state.Elapsed() > 40 && td.IsPressed {
		return ErrSuccess
	}
	a := math.Cos(float64(p.state.Elapsed())/60)*128 + 128
	rect = rect.SnapOutside(8, screenW, 40)
	p.lead.DrawR(sc, rect, color.RGBA{0, 0, 0, uint8(a)})

	return nil
}
