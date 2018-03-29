package main

import (
	"fmt"
	"image/color"

	td "github.com/ei1chi/tendon"
	et "github.com/hajimehoshi/ebiten"
)

func drawGame(screen *et.Image) {

	op := &et.DrawImageOptions{}

	screen.DrawImage(bgImage, op)

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

	if catched != nil {
		sp := sprites[table[catched.sexual].i]
		op := sp.Center()
		op.GeoM.Translate(real(catched.pos), imag(catched.pos))
		sp.Draw(screen, op)
	}

	for _, c := range charas {
		sp := sprites[table[c.sexual].i]
		op := sp.Center()
		op.GeoM.Translate(real(c.pos), imag(c.pos))
		sp.Draw(screen, op)
	}

	st := fmt.Sprintf("尊み: %d", score)
	gaugeText = td.NewText(mplus24, color.RGBA{0, 0, 0, 255}, st)
	gaugeText.Draw(screen, 0, 24, td.AlignLeft)
}
