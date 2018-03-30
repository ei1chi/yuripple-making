package main

import (
	"fmt"
	"image/color"

	td "github.com/ei1chi/tendon"
	et "github.com/hajimehoshi/ebiten"
	"golang.org/x/image/font"
)

type Menu struct {
	cursor        Cursor
	selection     difficulties
	lastSelection difficulties
	state         td.Stm
	mplus40       font.Face

	arrow, textEasy, textNormal, textHard *td.Text
}

type difficulties = int

const (
	easyLevel difficulties = iota
	normalLevel
	hardLevel
)

func (m *Menu) Load() {
	m.mplus40 = td.NewFontFace(root.mplus, 40)
	m.arrow = td.NewText(m.mplus40, color.Black, "→")
	m.selection = easyLevel
	m.cursor.pos = m.CursorPos()
	m.textEasy = td.NewText(m.mplus40, color.Black, "EASY")
	m.textNormal = td.NewText(m.mplus40, color.Black, "NORMAL")
	m.textHard = td.NewText(m.mplus40, color.Black, "HARD")
}

func (m *Menu) Update(sc *et.Image) error {
	x := screenW / 2
	y := 320.0
	m.textEasy.Draw(sc, x, y, td.AlignCenter)
	y = 400
	m.textNormal.Draw(sc, x, y, td.AlignCenter)
	y = 480
	m.textHard.Draw(sc, x, y, td.AlignCenter)

	// move cursor
	down := et.IsKeyPressed(et.KeyDown)
	if down {
		switch m.selection {
		case easyLevel:
			m.selection = normalLevel
			//case normalLevel:
			//m.selection = hardLevel
		}
	}

	up := et.IsKeyPressed(et.KeyUp)
	if up {
		switch m.selection {
		case normalLevel:
			m.selection = easyLevel
		case hardLevel:
			m.selection = normalLevel
		}
	}

	if m.selection != m.lastSelection {
		m.cursor.setPos(m.CursorPos())
		m.lastSelection = m.selection
		fmt.Println("changed")
	}
	m.cursor.move()

	m.arrow.Draw(sc, real(m.cursor.pos), imag(m.cursor.pos), td.AlignCenter)

	return nil
}

func (m *Menu) CursorPos() complex128 {
	switch m.selection {
	case easyLevel:
		return complex(100, 280)
	case normalLevel:
		return complex(100, 360)
	case hardLevel:
		return complex(100, 440)
	}
	return 0 + 1i
}

type Cursor struct {
	pos, last, next complex128
	isMoving        bool
	count           int
}

const cursorCountMax = 10

func (c *Cursor) setPos(p complex128) {
	if c.isMoving {
		return
	}
	c.last = c.pos
	c.next = p
	c.count = cursorCountMax
	c.isMoving = true
}

func (c *Cursor) move() {
	if !c.isMoving {
		return
	}
	t := 1.0 - float64(c.count)/float64(cursorCountMax)
	ratio := td.TweenRatio(0.2, 0.8, t)
	c.pos = c.last + (c.next-c.last)*complex(ratio, 0)
	fmt.Printf("ratio = %+v\n", ratio)

	c.count -= 1
	if c.count == 0 {
		c.pos = c.next
		c.isMoving = false
	}
}
