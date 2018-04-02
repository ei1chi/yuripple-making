package main

import (
	"fmt"
	"image/color"

	td "github.com/ei1chi/tendon"
	et "github.com/hajimehoshi/ebiten"
)

type Menu struct {
	state  td.Stm
	accept bool

	t struct {
		easy, normal, hard, help *td.TextBox
	}
}

func (m *Menu) Load() {

	halfr := td.Rect{0, 0, screenW, screenH / 2}.Shift(0, screenH/2)
	menu := halfr.HSplit(80, 80, 80)
	m.t.easy = td.NewTextBox(menu[0], root.mp40, 5, "EASY").Fit()
	m.t.normal = td.NewTextBox(menu[1], root.mp40, 5, "NORMAL").Fit()
	m.t.hard = td.NewTextBox(menu[2], root.mp40, 5, "HARD").Fit()

	helpr := halfr.SnapOutside(2, screenW, 60)
	m.t.help = td.NewTextBox(helpr, root.mp20, 5, "tap to select level")
}

func (m *Menu) Update(sc *et.Image) error {

	fmt.Println("menu update")
	if td.IsJustPressed {
		m.accept = true
	}

	f := func(t *td.TextBox) (b bool) {
		var c color.Color
		c = color.Gray{128}

		if t.R.Contains(td.CursorPos) {
			c = color.Black
			if m.accept && td.IsJustReleased {
				b = true
			}
		}
		t.Draw(sc, 0, 0, c)
		return
	}

	var err error
	var b bool
	fmt.Println("draw")
	b = f(m.t.easy)
	if b {
		root.level = levels[easyLevel]
		err = ErrSuccess
	}
	b = f(m.t.normal)
	if b {
		root.level = levels[normalLevel]
		err = ErrSuccess
	}
	b = f(m.t.hard)
	if b {
		root.level = levels[hardLevel]
		err = ErrSuccess
	}

	m.t.help.Draw(sc, 0, 0, color.Black)

	return err
}
