package main

import (
	"errors"
	_ "image/jpeg"
	"log"
	"math"
	"math/cmplx"
	"math/rand"
	"time"

	td "github.com/ei1chi/tendon"
	et "github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"

	"github.com/ei1chi/yuripple-making/game"
)

var (
	ErrSuccess = errors.New("successfully finished")
	score      int
)

const (
	around       = 4.0 // 4 phases
	screenWidth  = 480.0
	screenHeight = 640.0
)

func powi(angle float64) complex128 {
	return cmplx.Pow(1i, complex(angle, 0))
}

func absSq(c complex128) float64 {
	return math.Pow(real(c), 2) + math.Pow(imag(c), 2)
}

func isOutOfScreen(pos complex128, margin float64) bool {
	x, y := real(pos), imag(pos)
	if x < -margin || screenWidth+margin < x {
		return true
	}
	if y < -margin || screenHeight+margin < y {
		return true
	}
	return false
}

func main() {

	rand.Seed(time.Now().UnixNano())
	var err error

	// 画像読み込み
	atlas, err = td.NewAtlas("resources/atlas")
	if err != nil {
		log.Fatal(err)
	}

	loadSprites([]string{
		"nonke",
		"neco",
		"riba_neco",
		"tachi",
		"riba_tachi",
		"heart",
	})
	bgPath := "resources/yuri_bg.jpg"
	bgImage, _, err = ebitenutil.NewImageFromFile(bgPath, et.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	// フォント読み込み
	mplusFont, err := td.NewFont("resources/mplus-subset.ttf")
	if err != nil {
		log.Fatal(err)
	}
	mplus24 = td.NewFontFace(mplusFont, 24)

	// Run
	w, h, err := td.GetDeviceSize()
	s := 1.0
	if err != nil {
		s, sh := w/screenWidth, h/screenHeight
		if s < sh {
			s = sh
		}
	}

	// Create Scene
	curScene = game.CreateScene()
	err = et.Run(update, screenWidth, screenHeight, s, "百合っぷるメイキング")
	if err != nil && err != ErrSuccess {
		log.Fatal(err)
	}
}

type Scene interface {
	load()
	update(*et.Image) (Scene, error)
}

var (
	curScene, nextScene Scene
	task                = make(chan bool, 1)
)

func update(screen *et.Image) error {

	td.UpdateInput()
	next, err = curScene.update(screen, loading)

	if err != nil {
		if err == ErrSuccess {
			if nextScene == nil {
				return ErrSuccess // 終了
			}
			curScene = nextScene
		}
		return err
	}

	// 次のシーンを開始したい
	if next != nil {
		nextScene = next
		select {
		case task <- true: // タスクが空いている
			go func() {
				nextScene.load()
				<-task
			}()
		default:
		}
	}

}
