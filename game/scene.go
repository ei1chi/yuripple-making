package game

import (
	"fmt"

	td "github.com/ei1chi/tendon"
	et "github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"golang.org/x/image/font"

	"github.com/ei1chi/yuripple-making/ending"
)

type Scene struct {
	td.SceneBase
	state     *td.Stm
	atlas     *td.Atlas
	bgImage   *et.Image
	mplus24   font.Face
	gaugeText *td.Text

	charas []*Chara
}

type sceneState = int

const (
	playing sceneState = iota
	dying
)

func CreateScene() Scene {
	return &Scene{}
}

// load is invoked asynchronously
func (s *Scene) load() {
	s.atlas
}

func (s *Scene) update(screen *et.Image) (Scene, error) {

	switch s.state.Get() {
	case playing:
	case ending:
		if s.HasSlaveLoaded() && s.state.Elapsed() > 30 {
			return s.Slave, nil
		}
	}

	s.processCharas()
	s.sweepAll()
	s.collisionAll()
	s.drawAll(screen)

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

func (s *Scene) sweepAll() {
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

func (s *Scene) toEnding() {
	s.state.Transition(ending)
	s.Slave = &epilogue.Scene{}
	s.StartSlaveLoading() // async
}
