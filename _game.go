package game

import (
	"fmt"
	"log"
	"sync"

	td "github.com/ei1chi/tendon"
	et "github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"golang.org/x/image/font"
)

type gameState = int

const (
	playing gameState = iota
	ending
)

type Game struct {
	parent    *RootScene
	state     *td.Stm
	atlas     *td.Atlas
	bgImage   *et.Image
	mplus24   font.Face
	gaugeText *td.Text

	charas []*Chara
}

func (g *Game) Load() {
	s.StartNextLoading(&GameOverScene{}) // 同時に読み込み始める

	// 自分のリソース
	var err error
	s.state = &td.Stm{}
	s.atlas, err = td.NewAtlas("resources/atlas")
	if err != nil {
		log.Fatal(err)
	}

	s.loadSprites([]string{})
	bgPath := "resources/yuri_bg.jpg"
	s.bgImage, _, err = ebitenutil.NewImageFromFile(bgPath, et.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	mplusFont, err := td.NewFont("resources/mplus-subset.ttf")
	if err != nil {
		log.Fatal(err)
	}
	s.mplus24 = td.NewFontFace(mplusFont, 24)

	// next (GameOver) のロード待ち
	for {
		if next, _ := s.NextScene(); next != nil {
			break
		}
	}

	// 多数待ち（停止あり）
	wg := sync.WaitGroup{}

	s.gameOver = &GameOverScene{}
	wg.Add(1)
	go func() {
		s.gameOver.Load()
		wg.Done()
	}()

	wg.Wait()

	// 多数待ち（無停止）
	s.task = make(chan struct{}, 1)
	s.task <- struct{}{}

	go func() {
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go func() {
			s.gameOver.Load()
			wg.Done()
		}()
		wg.Wait()
		<-s.task
	}()

	select {
	case s.task <- struct{}{}:
		return
	default:
	}

	// 単数待ち（無停止）
	s.task = make(chan struct{}, 1)
	s.task <- struct{}{}

	go func() {
		s.next.Load()
		<-s.task
	}()

	select {
	case s.task <- struct{}{}:
		return
	default:
		// 他のこと
		return s.Parent.GameOver, nil
	}
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
}
