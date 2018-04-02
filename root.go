package main

import (
	"fmt"
	_ "image/jpeg"
	"log"

	td "github.com/ei1chi/tendon"
	"github.com/golang/freetype/truetype"
	et "github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"golang.org/x/image/font"
)

// global resources
var (
	mplusFont *truetype.Font
	mplus     map[int]font.Face
	display   td.Rect
)

type RootScene struct {
	pro   *Prologue
	menu  *Menu
	game  *Game
	state td.Stm
	bg    *et.Image
	level Level
}

type RootState = int

const (
	prologue RootState = iota
	menu
	game
)

func (r *RootScene) Load() {

	initLevels()
	display = td.Rect{0, 0, screenW, screenH}

	var err error
	mplusFont, err = td.NewFont("resources/mplus-1p-regular.ttf")
	if err != nil {
		log.Fatal(err)
	}
	mplus = map[int]font.Face{}
	for _, size := range []int{45, 40, 20} {
		mplus[size] = td.NewFontFace(mplusFont, float64(size))
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

	r.state.Transition(prologue)
}

func (r *RootScene) Update(sc *et.Image) error {

	td.UpdateInput()
	r.state.Update()

	// Draw
	sc.DrawImage(r.bg, nil)
	msg := fmt.Sprintf("FPS: %f", et.CurrentFPS())
	ebitenutil.DebugPrint(sc, msg)

	var err error
	switch r.state.Get() {
	case prologue:
		err = r.updatePrologue(sc)
	case menu:
		err = r.updateMenu(sc)
	case game:
		err = r.updateGame(sc)
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
	return err
}

func (r *RootScene) updateMenu(sc *et.Image) error {
	err := r.menu.Update(sc)
	if err == ErrSuccess {
		r.state.Transition(game)
		return nil
	}
	return err
}

func (r *RootScene) updateGame(sc *et.Image) error {
	err := r.game.Update(sc)
	if err == ErrSuccess {
		r.menu = &Menu{}
		r.menu.Load()
		r.game = &Game{}
		r.game.Load()
		r.state.Transition(menu) // メニューに戻る
		return nil
	}
	return err
}
