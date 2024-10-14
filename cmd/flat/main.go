package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/Flixnore/rubiks-cube-go/pkg/cube"
)

func main() {
	c := cube.NewSolved()
	g := NewGame(c)
	ebiten.SetWindowSize(600, 600)
	ebiten.SetWindowTitle("Rubik's Cube")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
