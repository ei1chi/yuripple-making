package main

import (
	td "github.com/ei1chi/tendon"
	et "github.com/hajimehoshi/ebiten"
)

func drawGame(screen *et.Image) {

	var sp *td.Sprite
	op := &et.DrawImageOptions{}

	screen.DrawImage(bgImage, op)
	gaugeText.Draw(screen, 0, 24, td.AlignLeft)

	drawCharas(screen)

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
		if catched != nil {
			if catched.sexual == t.s {
				op := sp.Center()
				op.GeoM.Translate(real(catched.pos), imag(catched.pos))
				screen.DrawImage(sp.Image, op)
			}
		}

		for _, c := range charas {
			if c.sexual == t.s {
				op := sp.Center()
				op.GeoM.Translate(real(c.pos), imag(c.pos))
				screen.DrawImage(sp.Image, op)
			}
		}
	}
}

func drawCharas(screen *et.Image) {

	for _, c := range charas {
		drawChara(screen, c)
	}
}
