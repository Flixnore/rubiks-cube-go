package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/Flixnore/rubiks-go/pkg/cube"
)

func main() {
	c := cube.NewSolved()
	g := NewGame(c)
	ebiten.SetWindowSize(600, 600)
	ebiten.SetWindowTitle("Rubik's Cube")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
	for {
		if c.Solved() {
			fmt.Println("Solved Cube!\n")
		}
		fmt.Printf("%s", c)
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter move(s): ")
		input, _ := reader.ReadString('\n')
		input = input[:len(input)-1]

		fmt.Print("\033[H\033[2J")
		if len(input) == 0 {
			continue
		}
		moves := strings.Split(input, " ")

		for _, m := range moves {
			iters := 1
			if strings.Contains(m, "2") {
				iters = 2
			}
			if strings.Contains(m, "'") {
				iters = 3
			}
			if strings.Contains(m, "undo") {
        c.Undo()
        continue
			}
			if strings.Contains(m, "redo") {
        c.Redo()
        continue
			}
			c.Rotate(cube.NewMove(cube.Face(m[0]), iters))
		}
	}
}
