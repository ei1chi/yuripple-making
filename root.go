package main

import (
	"fmt"
	"log"

	td "github.com/ei1chi/tendon"
	"github.com/golang/freetype/truetype"
	et "github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type RootScene struct {
	pro  *Prologue
	game *Game
	//epi   *Epilogue
	state td.Stm

	mplus *truetype.Font
	bg    *et.Image
}

type RootState = int

const (
	prologue RootState = iota
	game
	epilogue
)

func (r *RootScene) Load() {

	var err error
	r.mplus, err = td.NewFont("resources/mplus-1p-regular.ttf")
	if err != nil {
		log.Fatal(err)
	}
	r.bg, _, err = ebitenutil.NewImageFromFile("resources/yuri_bg.jpg", et.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	r.pro = &Prologue{}
	r.pro.Load() // sync (blocking)
	r.game = &Game{}
	r.game.Load() // sync (blocking)

	r.state.Transition(prologue)
}

func (r *RootScene) Update(sc *et.Image) error {

	td.UpdateInput()
	r.state.Update()

	// Draw
	msg := fmt.Sprintf("FPS: %f", et.CurrentFPS())
	ebitenutil.DebugPrint(sc, msg)
	sc.DrawImage(r.bg, nil)

	var err error
	switch r.state.Get() {
	case prologue:
		err = r.updatePrologue(sc)
	case game:
		err = r.updateGame(sc)
	case epilogue:
		err = r.updateEpilogue(sc)
	}

	return err
}

func (r *RootScene) updatePrologue(sc *et.Image) error {
	err := r.pro.Update(sc)
	if err == ErrSuccess {
		// To Game
		r.state.Transition(game)
		return nil
	}
	if err != nil {
		return err
	}
	return nil
}

func (r *RootScene) updateGame(sc *et.Image) error {
	err := r.game.Update(sc)
	if err != nil {
		return err
	}
	return nil
}

func (r *RootScene) updateEpilogue(sc *et.Image) error {
	err := r.game.Update(sc)
	if err != nil {
		return err
	}
	//err = r.epi.Update(sc)
	if err != nil {
		return err
	}
	return nil
}
