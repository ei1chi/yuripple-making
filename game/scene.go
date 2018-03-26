package game

import (
	"fmt"

	et "github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"golang.org/x/image/font"

	"github.com/ei1chi/yuripple-making/ending"
)

type GameScene struct {
	atlas     *td.Atlas
	bgImage   *et.Image
	mplus24   font.Face
	gaugeText *td.Text
}

func CreateScene() Scene {
	return &GameScene{}
}

func (s *GameScene) update(screen *et.Image) (Scene, error) {

	s.processCharas()
	s.sweepAll()
	s.collisionAll()
	s.draw(screen)

	// 終了判定
	quit := et.IsKeyPressed(et.KeyQ)
	if quit {
		return ending.CreateScene(), nil
	}

	// FPS
	str := "FPS: %f\n"
	ebitenutil.DebugPrint(screen, fmt.Sprintf(str, et.CurrentFPS()))

	return nil, nil
}

func (s *GameScene) sweepAll() {
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
