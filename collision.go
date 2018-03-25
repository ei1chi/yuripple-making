package main

import (
	"math"
)

func collisionAll() {

	// キャッチと全キャラで当たり判定
	// 当たったら指向に合わせて判定
	idx := -1
	if catched != nil {
		mindist := math.Pow(24, 2) // 若干狭い
		for i, c := range charas {
			if c.state.Get() == judgeNil {
				dist := absSq(catched.pos - c.pos)
				if dist < mindist {
					mindist = dist
					idx = i
				}
			}
		}
	}

	if idx != -1 {
		c := charas[idx]
		result := affinity(c.sexual, catched.sexual)
		score += (result - 1) * 100

		state := func() charaState {
			switch result {
			case judgeMiss:
				return charaMissed
			}
			return charaCoupled
		}()
		c.state.Transition(state)
		catched.state.Transition(state)
		center := (c.pos-catched.pos)/2 + catched.pos

		a := center - complex(32, 0)
		b := center + complex(32, 0)
		if real(c.pos) < real(catched.pos) {
			c.pos = a
			catched.pos = b
		} else {
			c.pos = b
			catched.pos = a
		}
		c.vec = complex(0, 0)
		catched.vec = complex(0, 0)

		charas = append(charas, catched)
		catched = nil
	}
}

func affinity(a, b Sexual) int {
	if b == nonke {
		return judgeMiss
	}

	switch a {
	case tachi:
		switch b {
		case tachi:
			return judgeOK
		case neco:
			return judgeGreat
		}
		return judgeNice

	case neco:
		switch b {
		case neco:
			return judgeOK
		}
		return judgeNice // riba

	case ribaTachi:
		switch b {
		case ribaNeco:
			return judgeNice
		}
		return judgeOK

	case ribaNeco:
		return judgeOK
	}

	return judgeMiss
}
