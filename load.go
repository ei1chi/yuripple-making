package main

import (
	"fmt"
	"log"

	td "github.com/ei1chi/tendon"
)

var sprites = map[string]*td.Sprite{}

func loadSprites(pngs []string) {
	var err error
	for _, name := range pngs {
		path := fmt.Sprintf("resources/%s.png", name)
		sprites[name], err = td.NewSprite(path)
		if err != nil {
			log.Fatal(err)
		}
	}
}
