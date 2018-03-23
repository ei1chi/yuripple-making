package main

import (
	"math/rand"
)

type Sexual int

const (
	nonke Sexual = iota
	tachi
	neco
	ribaTachi
	ribaNeco

	grav = complex(0, 0.01)
)

type Chara struct {
	isDead   bool
	sexual   Sexual
	pos, vec complex128
}

func (c *Chara) update() {
	c.vec += grav
	c.pos += c.vec
	x, y := real(c.pos), imag(c.pos)
	if x < -64 || 480+64 < x || y < -64 || 640+64 < y {
		c.isDead = true
	}
}

var charas []*Chara

var count int

func processCharas() {

	// spawn
	base := 0.6
	rng := 0.95 - base
	half := around / 2
	spd := complex(3, 0)

	interval := 50
	count += 1
	if count%interval == 0 {
		dir := base + rng*rand.Float64()
		c := &Chara{}
		c.sexual = Sexual(rand.Intn(int(ribaNeco + 1)))

		if count%(interval*2) == 0 {
			c.pos = complex(0, 640)
			c.vec = powi(around-dir) * spd
			charas = append(charas, c)
		} else {
			c.pos = complex(480, 640)
			c.vec = powi(half+dir) * spd
			charas = append(charas, c)
		}
	}

	// update
	for _, c := range charas {
		c.update()
	}
}

func catch(idx int) {
	catched = charas[idx]
	charas[idx] = nil
}
