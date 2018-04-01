package main

type difficulty = int // should be compatible with int for layouting and moving

const (
	easyLevel difficulty = iota
	normalLevel
	hardLevel
)

var levels = map[difficulty]Level{}

type Level struct {
	interval, limit            int
	nonke, neco, tachi, rn, rt float64
}

func (l *Level) total() float64 {
	return l.nonke + l.neco + l.tachi + l.rn + l.rt
}

func initLevels() {
	levels[easyLevel] = Level{
		interval: 70,
		limit:    60 * 10,
		nonke:    0,
		neco:     1,
		tachi:    1,
		rn:       1,
		rt:       1,
	}
}
