package main

import (
	"image/color"

	td "github.com/ei1chi/tendon"
	et "github.com/hajimehoshi/ebiten"
	"golang.org/x/image/font"
)

type Prologue struct {
	mplus48 font.Face
	state   td.Stm
}

func (p *Prologue) Load() {
	p.mplus48 = td.NewFontFace(root.mplus, 48)
}

func (p *Prologue) Update(sc *et.Image) error {
	t := td.NewText(p.mplus48, color.Black, "タイトルロゴ")
	t.Draw(sc, screenW/2, screenH/2, td.AlignCenter)
	p.state.Update()
	if p.state.Elapsed() > 40 && td.IsPressed {
		return ErrSuccess
	}
	return nil
}
