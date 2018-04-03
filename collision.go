package main

import (
	"math"

	td "github.com/ei1chi/tendon"
)

func (g *Game) collisionAll() {

	// キャッチキャラと全キャラで当たり判定
	var hit *Chara = nil
	if g.catched != nil {
		min := math.Pow(24, 2) // 若干狭い
		for _, c := range g.charas {
			if c.state.Get() == judgeNil {
				dist := td.AbsSq(g.catched.pos - c.pos)
				if dist < min {
					min = dist
					hit = c
				}
			}
		}
	}

	if hit != nil {
		result := affinity(hit.sexual, g.catched.sexual)
		g.score += (int(result) - 1) * 100

		state := charaCoupled
		switch result {
		case judgeMiss:
			state = charaMissed
		}
		hit.state.Transition(int(state))
		g.catched.state.Transition(int(state))
		center := (hit.pos-g.catched.pos)/2 + g.catched.pos

		a := center - complex(32, 0)
		b := center + complex(32, 0)
		if real(hit.pos) < real(g.catched.pos) {
			hit.pos = a
			g.catched.pos = b
		} else {
			hit.pos = b
			g.catched.pos = a
		}
		hit.vec = complex(0, 0)
		g.catched.vec = complex(0, 0)
		g.charas = append(g.charas, g.catched)
		g.catched = nil

	}

}

func affinity(a, b Sexual) Judgement {

	return judgeNice
}
