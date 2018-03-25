package main

import (
	"errors"
	"fmt"
	"image/color"
	_ "image/jpeg"
	"log"
	"math"
	"math/cmplx"
	"math/rand"
	"time"

	td "github.com/ei1chi/tendon"
	et "github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"golang.org/x/image/font"
)

var (
	ErrSuccess = errors.New("successfully finished")
	bgImage    *et.Image
	mplus24    font.Face
	gaugeText  *td.Text
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
	mplusFont, err := td.NewFont("resources/mplus-subset.ttf")
	if err != nil {
		panic(err)
	}
	mplus24 = td.NewFontFace(mplusFont, 24)

	gaugeText = td.NewText(mplus24, color.RGBA{0, 0, 0, 255}, "尊みゲージ")

	// Run
	w, h, err := td.GetDeviceSize()
	s := 1.0
	if err != nil {
		s, sh := w/screenWidth, h/screenHeight
		if s < sh {
			s = sh
		}
	}

	err = et.Run(update, screenWidth, screenHeight, s, "百合っぷるメイキング")
	if err != nil && err != ErrSuccess {
		log.Fatal(err)
	}
}

func update(screen *et.Image) error {

	td.UpdateInput()
	processCharas()
	sweepAll()
	collisionAll()
	drawGame(screen)

	// 終了判定
	quit := et.IsKeyPressed(et.KeyQ)
	if quit {
		return ErrSuccess
	}

	// FPS
	str := "FPS: %f\n"
	str += "charas : %d\n"
	str += "score : %d\n"
	ebitenutil.DebugPrint(screen, fmt.Sprintf(str, et.CurrentFPS(), len(charas), score))

	return nil
}

func sweepAll() {
	next := charas[:0]
	for _, c := range charas {
		if c != nil {
			if !c.isDead {
				next = append(next, c)
			}
		}
	}
	charas = next
}
