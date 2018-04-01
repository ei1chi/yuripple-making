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
	menu *Menu
	game *Game
	//epi   *Epilogue
	state td.Stm

	mplus *truetype.Font
	rect  td.Rect
	bg    *et.Image
	diff  difficulty
}

type RootState = int

const (
	prologue RootState = iota
	menu
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
	r.menu = &Menu{}
	r.menu.Load() // sync (blocking)
	r.game = &Game{}
	r.game.Load() // sync (blocking)

	r.rect = td.Rect{0, 0, screenW, screenH}
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
	case menu:
		err = r.updateMenu(sc)
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
		// To Menu
		r.state.Transition(menu)
		return nil
	}
	if err != nil {
		return err
	}
	return nil
}

func (r *RootScene) updateMenu(sc *et.Image) error {
	err := r.menu.Update(sc)
	if err == ErrSuccess {
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
