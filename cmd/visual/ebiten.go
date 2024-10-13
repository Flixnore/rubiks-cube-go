package main

import (
	"github.com/Flixnore/rubiks-go/pkg/cube"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"time"
)

const tileSize = 30
const tileTotalSize = 35.0
const moveTimeout = 150 * time.Millisecond

type Game struct {
	cube     cube.Cube
	lastMove time.Time
}

func NewGame(cube cube.Cube) *Game {
	game := &Game{
		cube: cube,
	}
	return game
}

func (g *Game) Update() error {
	keyToMove := map[ebiten.Key]cube.Face{
		ebiten.KeyI:         cube.U,
		ebiten.KeyJ:         cube.L,
		ebiten.KeyK:         cube.F,
		ebiten.KeyL:         cube.R,
		ebiten.KeySemicolon: cube.B,
		ebiten.KeyComma:     cube.D,
		ebiten.KeyZ:         cube.Z,
		ebiten.KeyX:         cube.X,
		ebiten.KeyC:         cube.Y,
	}
	for key, face := range keyToMove {
		if ebiten.IsKeyPressed(key) && time.Since(g.cube.LastMoveTime()) > moveTimeout {
			if ebiten.IsKeyPressed(ebiten.KeyShift) {
				g.cube.Rotate(cube.NewMove(face, 3))
			} else {
				g.cube.Rotate(cube.NewMove(face, 1))
			}
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyU) && time.Since(g.cube.LastMoveTime()) > moveTimeout {
		g.cube.Undo()
	}
	if ebiten.IsKeyPressed(ebiten.KeyR) && ebiten.IsKeyPressed(ebiten.KeyControl) && time.Since(g.cube.LastMoveTime()) > moveTimeout {
		g.cube.Redo()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	colors := g.cube.State()
	start := 0.0

	drawNine(screen, colors[0:9], start+tileTotalSize*3, start)
	drawNine(screen, colors[9:18], start, start+tileTotalSize*3)
	drawNine(screen, colors[18:27], start+tileTotalSize*3, start+tileTotalSize*3)
	drawNine(screen, colors[27:36], start+tileTotalSize*6, start+tileTotalSize*3)
	drawNine(screen, colors[36:45], start+tileTotalSize*9, start+tileTotalSize*3)
	drawNine(screen, colors[45:54], start+tileTotalSize*3, start+tileTotalSize*6)
}

func drawNine(screen *ebiten.Image, colors []cube.Color, posx float64, posy float64) {
	for i, sq := range colors {
		ebitenutil.DrawRect(screen, posx, posy, tileSize, tileSize, sq.RGBA())
		posx += tileTotalSize
		if (i+1)%3 == 0 {
			posy += tileTotalSize
			posx -= tileTotalSize * 3
		}
	}
}

func (g *Game) Layout(outsideWidth, insideWidth int) (int, int) {
	return 600, 600
}
