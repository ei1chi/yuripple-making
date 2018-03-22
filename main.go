package main

import (
	"errors"
	"fmt"
	"image/color"
	_ "image/jpeg"
	"log"
	"math/cmplx"
	"math/rand"
	"time"

	"github.com/ei1chi/tendon"
	et "github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"golang.org/x/image/font"
)

var (
	ErrSuccess = errors.New("successfully finished")
	bgImage    *et.Image
	mplus24    font.Face
	gaugeText  *tendon.Text
)

const (
	around = 4.0 // 4 phases
)

func powi(angle float64) complex128 {
	return cmplx.Pow(1i, complex(angle, 0))
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// 画像読み込み
	loadSprites([]string{
		"nonke",
		"neco",
		"riba_neco",
		"tachi",
		"riba_tachi",
		"heart",
	})
	var err error
	bgPath := "resources/yuri_bg.jpg"
	bgImage, _, err = ebitenutil.NewImageFromFile(bgPath, et.FilterDefault)
	if err != nil {
		panic(err)
	}

	// フォント読み込み
	mplusFont, err := tendon.NewFont("resources/mplus-subset.ttf")
	if err != nil {
		panic(err)
	}
	mplus24 = tendon.NewFontFace(mplusFont, 24)

	gaugeText = tendon.NewText(mplus24, color.RGBA{0, 0, 0, 255}, "尊みゲージ")

	// Run
	s := getScale()
	err = et.Run(update, 480, 640, s, "百合っぷるメイキング")
	if err != nil && err != ErrSuccess {
		log.Fatal(err)
	}
}

func update(screen *et.Image) error {

	updateInput()
	processCharas()
	//collisionAll()
	drawAll(screen)
	sweepAll()

	// 終了判定
	quit := et.IsKeyPressed(et.KeyQ)
	if (480-32) < cursorX && cursorY < 32 {
		if pressed {
			quit = true
		}
	}
	if quit {
		return ErrSuccess
	}

	// FPS
	str := "FPS: %f\n"
	str += "charas : %d\n"
	ebitenutil.DebugPrint(screen, fmt.Sprintf(str, et.CurrentFPS(), len(charas)))

	return nil
}

func sweepAll() {
	next := charas[:0]
	for _, c := range charas {
		if !c.isDead {
			next = append(next, c)
		}
	}
	charas = next
}
