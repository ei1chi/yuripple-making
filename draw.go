package main

import (
	"fmt"
	"image"
	"image/color"

	et "github.com/hajimehoshi/ebiten"
)

var table = []struct {
	n string
	s Sexual
}{
	{"nonke", nonke},
	{"tachi", tachi},
	{"neco", neco},
	{"riba_tachi", ribaTachi},
	{"riba_neco", ribaNeco},
}

func (g *Game) drawAll(sc *et.Image) {

	// draw characters
	for _, c := range g.charas {
		if c.state.Get() == int(charaMoving) {
			sp := g.sprites[table[c.sexual].n]
			op := sp.Center()
			op.GeoM.Translate(real(c.pos), imag(c.pos))
			sp.Draw(sc, op)
		}
	}
	// draw fading out characters
	for _, c := range g.charas {
		if c.state.Get() != int(charaMoving) {
			sp := g.sprites[table[c.sexual].n]
			op := sp.Center()
			op.GeoM.Translate(real(c.pos), imag(c.pos))
			op.ColorM.Translate(0, 0, 0, -float64(c.state.Elapsed())/100)
			sp.Draw(sc, op)
		}
	}
	if g.catched != nil {
		sp := g.sprites[table[g.catched.sexual].n]
		op := sp.Center()
		op.GeoM.Translate(real(g.catched.pos), imag(g.catched.pos))
		sp.Draw(sc, op)
	}

	// draw ui
	if g.state.Get() == gamePlaying {
		g.t.mode.T.SetText(fmt.Sprintf("MODE: %s", root.level.name))
		g.t.mode.Draw(sc, 0, 0, color.Black)
		g.t.score.T.SetText(fmt.Sprintf("score: %d", g.score))
		g.t.score.Draw(sc, 0, 0, color.Black)
		g.t.time.Draw(sc, 0, 0, color.Black)

		x, y := g.t.time.R.AnchorPos(3)
		op := &et.DrawImageOptions{}
		op.GeoM.Translate(x, y)
		w, h := g.fullGauge.Size()
		w -= 100 * w * g.state.Elapsed() / root.level.limit / 100
		rc := image.Rect(0, 0, w, h)
		op.SourceRect = &rc
		sc.DrawImage(g.fullGauge, op)
	}
}
