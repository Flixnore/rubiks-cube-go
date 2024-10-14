package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Flixnore/rubiks-cube-go/pkg/cube"
)

func main() {
	c := cube.NewSolved()
	for {
		if c.Solved() {
			fmt.Println("Solved Cube!\n")
		}
		fmt.Printf("%s", c)
		fmt.Print("Enter move(s): ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		fmt.Print("\033[H\033[2J")

    if input == "undo\n" {
      c.Undo()
      continue
    }

    if input == "redo\n" {
      c.Redo()
      continue
    }

    moves := cube.ParseMoves(input)
		for _, m := range moves {
			c.Rotate(m)
		}
	}
}

