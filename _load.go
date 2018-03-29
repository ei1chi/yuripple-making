package main

import (
	"log"

	td "github.com/ei1chi/tendon"
)

var sprites = map[string]*td.Sprite{}

func loadSprites(pngs []string) {
	var err error
	for _, p := range pngs {
		sprites[p], err = atlas.NewSprite(p + ".png")
		if err != nil {
			log.Fatal(err)
		}
	}
}
