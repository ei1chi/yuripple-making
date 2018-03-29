package main

import (
	"math"
	"math/rand"

	td "github.com/ei1chi/tendon"
)

type Sexual = int

const (
	nonke Sexual = iota
	tachi
	neco
	ribaTachi
	ribaNeco

	grav = complex(0, 0.01)
)

type charaState = int

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
	state    *td.Stm
}

func (c *Chara) update() {
	c.state.Update()
	switch c.state.Get() {
	case charaMoving:
		c.vec += grav
		c.pos += c.vec
		if isOutOfScreen(c.pos, 64) {
			c.isDead = true
		}
	case charaCoupled:
		c.pos += complex(0, -1)
		if c.state.Elapsed() > 100 {
			c.isDead = true
		}
	case charaMissed:
		c.pos += complex(0, 1)
		if c.state.Elapsed() > 100 {
			c.isDead = true
		}
	}
}

var (
	charas  []*Chara
	catched *Chara
	offset  complex128
	count   int
)

func processCharas() {

	// spawn
	base := 0.6
	rng := 0.95 - base
	half := around / 2
	spd := complex(3, 0)

	interval := 40
	count += 1
	if count%interval == 0 {
		dir := base + rng*rand.Float64()
		c := &Chara{}
		c.sexual = rand.Intn(ribaNeco + 1)
		c.state = &td.Stm{}

		if count%(interval*2) == 0 {
			c.pos = complex(0, screenHeight)
			c.vec = powi(around-dir) * spd
			charas = append(charas, c)
		} else {
			c.pos = complex(screenWidth, screenHeight)
			c.vec = powi(half+dir) * spd
			charas = append(charas, c)
		}
	}

	// update CATCHED chara
	if catched != nil {
		catched.state.Update()
		catched.vec = complex(0, 0)
		catched.pos = td.CursorPos + offset
	}

	if td.IsJustPressed {
		// catch & update
		mindist := math.Pow(32, 2)
		idx := -1
		for i, c := range charas {
			c.update()
			if c.state.Get() == judgeNil && !c.isDead {
				dist := absSq(c.pos - td.CursorPos)
				if dist < mindist {
					mindist = dist
					idx = i
				}
			}
		}

		if idx != -1 {
			catched = charas[idx]
			charas[idx] = nil
			offset = catched.pos - td.CursorPos
		}
	} else {
		// update & release
		for _, c := range charas {
			c.update()
		}
		if !td.IsPressed && catched != nil {
			charas = append(charas, catched)
			catched = nil
		}
	}

}
