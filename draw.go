package main

import (
	td "github.com/ei1chi/tendon"
	et "github.com/hajimehoshi/ebiten"
)

func drawAll(screen *et.Image) {

	var sp *Sprite
	op := &et.DrawImageOptions{}

	screen.DrawImage(bgImage, op)
	gaugeText.Draw(screen, 0, 24, td.AlignLeft)

	table := []struct {
		i string
		s Sexual
	}{
		{"nonke", nonke},
		{"tachi", tachi},
		{"neco", neco},
		{"riba_tachi", ribaTachi},
		{"riba_neco", ribaNeco},
	}

	for _, t := range table {
		sp = sprites[t.i]
		for _, c := range charas {
			if c.sexual == t.s {
				op = sp.center()
				op.GeoM.Translate(real(c.pos), imag(c.pos))
				screen.DrawImage(sp.image, op)
			}
		}
	}
}
