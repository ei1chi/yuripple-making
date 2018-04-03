package main

import td "github.com/ei1chi/tendon"

type Howto struct {
}

func (h *Howto) Load() {
	step0 := td.Rect{0, 0, screenW, 240}
	step1 := step0.SnapOutside(8, screenW, 240)
	step2 := step1.SnapOutside(8, screenW, 240)
}
