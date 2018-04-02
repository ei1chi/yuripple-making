package main

import (
	"math/rand"

	td "github.com/ei1chi/tendon"
)

func (g *Game) processCharas() {

	// spawn
	lower, higher := 0.6, 0.95
	rng := higher - lower
	spd := complex(5, 0)
	level := root.level

	if g.state.Elapsed()%level.interval == 0 {
		dir := lower + rng*rand.Float64()
		c := &Chara{}
		c.sexual = level.lot()

		if g.state.Elapsed()%(level.interval*2) == 0 {
			c.pos = complex(0, screenH)
			c.vec = td.Powi(4.0-dir) * spd
		} else {
			c.pos = complex(screenW, screenH)
			c.vec = td.Powi(2.0+dir) * spd
		}
		g.charas = append(g.charas, c)
	}

	// update
	for _, c := range g.charas {
		c.update()
	}
}

const grav = complex(0, 0.03)

type charaState int

const (
	charaMoving charaState = iota
	charaCoupled
	charaMissed
)

type Judgement = int

const (
	judgeNil Judgement = iota
	judgeMiss
	judgeOK
	judgeNice
	judgeGreat
)

type Chara struct {
	isDead   bool
	sexual   Sexual
	pos, vec complex128
	state    td.Stm
}

func (c *Chara) update() {

	c.state.Update()

	switch charaState(c.state.Get()) {
	case charaMoving:
		c.vec += grav
		c.pos += c.vec
		if isOutOfArea(c.pos, 64) {
			c.isDead = true
		}
	}
}
