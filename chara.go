package main

import (
	"math"
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

	// update CATCHED chara
	if g.catched != nil {
		g.catched.state.Update()
		g.catched.vec = complex(0, 0)
		g.catched.pos = td.CursorPos + g.offset
	}

	if g.state.Get() == gamePlaying && td.IsJustPressed {
		// catch one & update all
		min := math.Pow(32, 2)
		idx := -1
		for i, c := range g.charas {
			c.update()
			if c.state.Get() == judgeNil {
				dist := td.AbsSq(c.pos - td.CursorPos)
				if dist < min {
					min = dist
					idx = i
				}
			}
		}

		if idx != -1 {
			g.catched = g.charas[idx]
			g.charas[idx] = nil
			g.offset = g.catched.pos - td.CursorPos
		}
	} else {
		// update all & release one
		for _, c := range g.charas {
			c.update()
		}

		if g.catched != nil && !td.IsPressed {
			g.charas = append(g.charas, g.catched)
			g.catched = nil
		}
	}
}

const grav = complex(0, 0.03)

type charaState int

const (
	charaMoving charaState = iota
	charaCatched
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
