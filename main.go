package main

import (
	"errors"
	_ "image/jpeg"
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
	score      int
	root       *RootScene
)

func isOutOfScreen(pos complex128, margin float64) bool {
	x, y := real(pos), imag(pos)
	if x < -margin || screenW+margin < x {
		return true
	}
	if y < -margin || screenH+margin < y {
		return true
	}
	return false
}

func main() {

	rand.Seed(time.Now().UnixNano())
	var err error

	// Create Scene
	root = &RootScene{}
	root.Load()

	// Run
	s := td.DisplayScale(screenW, screenH)
	err = et.Run(update, screenW, screenH, s, "百合っぷるメイキング")
	if err != nil && err != ErrSuccess {
		log.Fatal(err)
	}
}

func update(screen *et.Image) error {
	err := root.Update(screen)
	if err != nil {
		return err
	}
	return nil
}
