package main

import "math/rand"

type Sexual int

const (
	nonke Sexual = iota
	tachi
	neco
	ribaTachi
	ribaNeco
)

type gameLevel int

const (
	easyLevel gameLevel = iota
	normalLevel
	hardLevel
)

var levels = map[gameLevel]Level{}

type Level struct {
	interval, limit int
	name            string
	weight          map[Sexual]int
}

func (l *Level) lot() Sexual {
	total := 0
	for _, w := range l.weight {
		total += w
	}

	val := rand.Intn(total)
	total = 0
	for s, w := range l.weight {
		if total <= val && val < total+w {
			return s
		}
		total += w
	}
	return nonke
}

func initLevels() {
	levels[easyLevel] = Level{
		interval: 70,
		limit:    60 * 20,
		name:     "EASY",
		weight: map[Sexual]int{
			nonke:     0,
			tachi:     1,
			neco:      1,
			ribaTachi: 1,
			ribaNeco:  1,
		},
	}

	levels[normalLevel] = Level{
		interval: 70,
		limit:    60 * 30,
		name:     "NORMAL",
		weight: map[Sexual]int{
			nonke:     0,
			tachi:     1,
			neco:      1,
			ribaTachi: 1,
			ribaNeco:  1,
		},
	}

	levels[hardLevel] = Level{
		interval: 70,
		limit:    60 * 20,
		name:     "HARD",
		weight: map[Sexual]int{
			nonke:     0,
			tachi:     1,
			neco:      1,
			ribaTachi: 1,
			ribaNeco:  1,
		},
	}
}
