package main

import (
	"image/color"

	td "github.com/ei1chi/tendon"
	et "github.com/hajimehoshi/ebiten"
)

type Menu struct {
	state  td.Stm
	accept bool

	rs struct {
		menu []td.Rect
		help td.Rect
	}
	ts struct {
		easy, normal, hard, help *td.Text
	}
}

func (m *Menu) Load() {

	halfr := td.Rect{0, 0, screenW, screenH}.HSplit(320)[1]
	m.rs.menu = halfr.WithMargin(100, 0, 100, 0).HSplit(80, 80, 80)
	m.ts.easy = td.NewText(root.mplus, 40, 5, "EASY")
	m.ts.normal = td.NewText(root.mplus, 40, 5, "NORMAL")
	m.ts.hard = td.NewText(root.mplus, 40, 5, "HARD")

	m.rs.help = halfr.SnapOutside(2, screenW, 60)
	m.ts.help = td.NewText(root.mplus, 20, 5, "tap to select level")
}

func (m *Menu) Update(sc *et.Image) error {

	if td.IsJustPressed {
		m.accept = true
	}

	f := func(t *td.Text, r td.Rect) (b bool) {
		var c color.Color
		c = color.Gray{128}

		if r.Contains(td.CursorPos) {
			c = color.Black
			if m.accept && td.IsJustReleased {
				b = true
			}
		}
		t.DrawR(sc, r, c)
		return b
	}

	var b bool
	var err error
	b = f(m.ts.easy, m.rs.menu[0])
	if b {
		root.diff = easyLevel
		err = ErrSuccess
	}
	b = f(m.ts.normal, m.rs.menu[1])
	if b {
		root.diff = normalLevel
		err = ErrSuccess
	}
	b = f(m.ts.hard, m.rs.menu[2])
	if b {
		root.diff = hardLevel
		err = ErrSuccess
	}

	m.ts.help.DrawR(sc, m.rs.help, color.Black)

	return err
}
