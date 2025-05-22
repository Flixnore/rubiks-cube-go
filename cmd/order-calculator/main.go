package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Flixnore/rubiks-cube-go/pkg/cube"
)

func main() {
	for {
		c := cube.NewSolved().State()()
		fmt.Print("Enter move(s): ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		repetitions := 0
    moves := cube.ParseMoves(input)

		for {
			repetitions++;
			for _, m := range moves {
				c.Rotate(m)
			}

			if c.Solved() {
				break
			}
		}

		fmt.Printf("Repetitions until solved: %v\n", repetitions)
	}
}

