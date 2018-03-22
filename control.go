package main

import "math"

var (
	isCatching bool
	offset     complex128
)

func cursorControl() {

	mindist := 9999.9
	limit := math.Pow(48+48, 2)

	for _, c := range charas {
		cursor := complex(cursorX, cursorY)
		diff := c.pos - cursor
		dist := math.Pow(real(diff), 2) + math.Pow(imag(diff), 2)
		if dist < limit && dist < mindist {
			mindist = dist
		}
	}

}
