package cube

import (
	"fmt"
	"time"
)

type Cube interface {
	Rotate(Move)
	Solved() bool
	State() [54]Color
	LastMoveTime() time.Time
	History() []Move
	Undo()
	Redo()
}

type cube struct {
	colors       [54]Color
	history      []Move
  undoHistory  []Move
	lastMoveTime time.Time
}

func NewSolved() Cube {
	return &cube{
		colors: solvedColors}
}

func (c *cube) Rotate(m Move) {
  c.rotate(m)
	c.history = append(c.history, m)
  c.undoHistory = []Move{}
}

func (c cube) Solved() bool {
	return c.colors == solvedColors
}

func (c *cube) State() [54]Color {
	return c.colors
}

func (c *cube) LastMoveTime() time.Time {
	return c.lastMoveTime
}

func (c *cube) History() []Move {
	return c.history
}

func (c *cube) Undo() {
	if len(c.history) == 0 {
		return
	}

	undoMove := c.history[len(c.history)-1]
  c.undoHistory = append(c.undoHistory, undoMove)
	undoMove.iters = 4 - undoMove.iters
	c.rotate(undoMove)
	c.history = c.history[:len(c.history)-1]
}

func (c *cube) Redo() {
  if len(c.undoHistory) == 0 {
    return 
  }

  redo := c.undoHistory[len(c.undoHistory)-1]
  c.rotate(redo)
  c.undoHistory = c.undoHistory[:len(c.undoHistory)-1]
  c.history = append(c.history, redo)
}

// The cube is represented as a slice of 54 colors.

//            0  1  2
//            3  4  5
//            6  7  8
// 9  10 11 | 18 19 20 | 27 28 29 | 36 37 38
// 12 13 14 | 21 22 23 | 30 31 32 | 39 40 41
// 15 16 17 | 24 25 26 | 33 34 35 | 42 43 44
//            45 46 47
//            48 49 50
//            51 52 53

// This cube is rotated twice around the y-axis
// (for easier understanding of B moves)
//            8  7  6
//            5  4  3
//            2  1  0
// 27 28 29 | 36 37 38 | 9  10 11 | 18 19 20 |
// 30 31 32 | 39 40 41 | 12 13 14 | 21 22 23 |
// 33 34 35 | 42 43 44 | 15 16 17 | 24 25 26 |
//            53 52 51
//            50 49 48
//            47 46 45

func (c *cube) String() string {
	return fmt.Sprintf(
		"        %s %s %s\n"+
			"        %s %s %s\n"+
			"        %s %s %s\n"+
			"%s %s %s | %s %s %s | %s %s %s | %s %s %s\n"+
			"%s %s %s | %s %s %s | %s %s %s | %s %s %s\n"+
			"%s %s %s | %s %s %s | %s %s %s | %s %s %s\n"+
			"        %s %s %s\n"+
			"        %s %s %s\n"+
			"        %s %s %s\n"+
			"History: %s\n" +
			"Undo History: %s\n",
		c.colors[0], c.colors[1], c.colors[2],
		c.colors[3], c.colors[4], c.colors[5],
		c.colors[6], c.colors[7], c.colors[8],
		c.colors[9], c.colors[10], c.colors[11], c.colors[18], c.colors[19], c.colors[20], c.colors[27], c.colors[28], c.colors[29], c.colors[36], c.colors[37], c.colors[38],
		c.colors[12], c.colors[13], c.colors[14], c.colors[21], c.colors[22], c.colors[23], c.colors[30], c.colors[31], c.colors[32], c.colors[39], c.colors[40], c.colors[41],
		c.colors[15], c.colors[16], c.colors[17], c.colors[24], c.colors[25], c.colors[26], c.colors[33], c.colors[34], c.colors[35], c.colors[42], c.colors[43], c.colors[44],
		c.colors[45], c.colors[46], c.colors[47],
		c.colors[48], c.colors[49], c.colors[50],
		c.colors[51], c.colors[52], c.colors[53],
		c.history,
		c.undoHistory,
	)
}
