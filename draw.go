package main

import (
	"fmt"
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
		sp := g.sprites[table[c.sexual].n]
		op := sp.Center()
		op.GeoM.Translate(real(c.pos), imag(c.pos))
		sp.Draw(sc, op)
	}

	// draw ui
	if g.state.Get() == gamePlaying {
		g.t.mode.T.SetText(fmt.Sprintf("MODE: %s", root.level.name))
		g.t.mode.Draw(sc, 0, 0, color.Black)
		g.t.score.T.SetText(fmt.Sprintf("score: %d", g.state.Elapsed()))
		g.t.score.Draw(sc, 0, 0, color.Black)
		g.t.time.Draw(sc, 0, 0, color.Black)

		r := float64(g.state.Elapsed()) / float64(root.level.limit)
		w := g.r.gauge.Width()*(1.0-r) + 1
		h := g.r.gauge.Height()
		gauge, _ := et.NewImage(int(w), int(h), et.FilterDefault)
		gauge.Fill(color.Black)
		op := &et.DrawImageOptions{}
		x, y := g.t.time.R.AnchorPos(6)
		op.GeoM.Translate(x, y-h/2)
		sc.DrawImage(gauge, op)
	}
}
