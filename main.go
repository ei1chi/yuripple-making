package main

import (
	"errors"
	"log"
	"math/rand"
	"time"

	td "github.com/ei1chi/tendon"
	et "github.com/hajimehoshi/ebiten"
)

const (
	screenW = 480.0
	screenH = 640.0
)

var (
	ErrSuccess = errors.New("successfully finished")
	root       *RootScene
)

func isOutOfArea(pos complex128, margin float64) bool {
	x, y := real(pos), imag(pos)
	return x < -margin || y < -margin || margin+screenW < x || margin+screenH < y
}

func main() {

	rand.Seed(time.Now().UnixNano())

	// Create Scene
	root = &RootScene{}
	root.Load()

	// Run
	s := td.DisplayScale(screenW, screenH)

	err := et.Run(update, screenW, screenH, s, "百合っぷるメイキング")
	if err != nil && err != ErrSuccess {
		log.Fatal(err)
	}
}

func update(screen *et.Image) error {
	if et.IsRunningSlowly() {
		return nil
	}
	return root.Update(screen)
}
