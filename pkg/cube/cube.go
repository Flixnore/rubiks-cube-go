package cube

import (
	"fmt"
	"time"
  "strings"
	colores "image/color"
)

type Move struct {
	face Face
	iters int
}

type Face string

const (
	F Face = "F"
	U      = "U"
	D      = "D"
	L      = "L"
	R      = "R"
	B      = "B"
	X      = "x"
	Y      = "y"
	Z      = "z"
)

func NewMove(f Face, iters int) Move {
  return Move{f, iters}
}

func (m Move) String() string {
	if m.iters == 2 {
		return string(m.face) + "2"
	} else if m.iters == 3 {
		return string(m.face) + "'"
	} else {
		return string(m.face)
	}
}

func ParseMoves(input string) (moves []Move) {
  input = strings.Trim(input, "\n")
  input = strings.Replace(input, "(", "", -1)
  input = strings.Replace(input, ")", "", -1)

  for _, m := range strings.Split(input, " ") {
    if len(m) == 0 {
      continue
    }
    iters := 1
    if strings.Contains(m, "2") {
      iters = 2
    }
    if strings.Contains(m, "'") {
      iters = 3
    }
    m := NewMove(Face(m[0]), iters)
    moves = append(moves, m)
  }
  return moves
}

type Color string

const (
	Wh Color = "W"
	Yw       = "Y"
	Bl       = "B"
	Gr       = "G"
	Or       = "O"
	Rd       = "R"
)

var solvedColors = [54]Color{
	Wh, Wh, Wh, Wh, Wh, Wh, Wh, Wh, Wh,
	Or, Or, Or, Or, Or, Or, Or, Or, Or,
	Gr, Gr, Gr, Gr, Gr, Gr, Gr, Gr, Gr,
	Rd, Rd, Rd, Rd, Rd, Rd, Rd, Rd, Rd,
	Bl, Bl, Bl, Bl, Bl, Bl, Bl, Bl, Bl,
	Yw, Yw, Yw, Yw, Yw, Yw, Yw, Yw, Yw,
}

func (c Color) RGBA() colores.RGBA {
	switch c {
	case Wh:
		return colores.RGBA{255, 255, 255, 255}
	case Yw:
		return colores.RGBA{255, 255, 0, 255}
	case Bl:
		return colores.RGBA{0, 0, 255, 255}
	case Gr:
		return colores.RGBA{0, 255, 0, 255}
	case Or:
		return colores.RGBA{255, 140, 30, 255}
	case Rd:
		return colores.RGBA{255, 0, 0, 255}
	}
	return colores.RGBA{255, 255, 255, 255}
}


func (c Color) String() string {
	switch c {
	case Wh:
		return "W"
	case Yw:
		return "\033[38;5;226mY\033[0m"
	case Bl:
		return "\033[34mB\033[0m"
	case Gr:
		return "\033[32mG\033[0m"
	case Or:
		return "\033[38;5;214mO\033[0m"
	case Rd:
		return "\033[31mR\033[0m"
	}
	return ""
}

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

func (c *cube) State() [54]Color {
	return c.colors
}

func (c *cube) History() []Move {
	return c.history
}

func (c *cube) swapFour(a, b, cc, d int) {
	temp := c.colors[d]
	c.colors[d] = c.colors[cc]
	c.colors[cc] = c.colors[b]
	c.colors[b] = c.colors[a]
	c.colors[a] = temp
}

func (c *cube) Rotate(m Move) {
  c.rotate(m)
	c.history = append(c.history, m)
  c.undoHistory = []Move{}
}

func (c *cube) rotate(m Move) {
	for range m.iters {
		switch m.face {
		case F:
			c.swapFour(20, 26, 24, 18)
			c.swapFour(23, 25, 21, 19)
			c.swapFour(6, 27, 47, 17)
			c.swapFour(7, 30, 46, 14)
			c.swapFour(8, 33, 45, 11)
		case U:
			c.swapFour(0, 2, 8, 6)
			c.swapFour(1, 5, 7, 3)
			c.swapFour(38, 29, 20, 11)
			c.swapFour(37, 28, 19, 10)
			c.swapFour(36, 27, 18, 9)
		case D:
			c.swapFour(45, 47, 53, 51)
			c.swapFour(46, 50, 52, 48)
			c.swapFour(15, 24, 33, 42)
			c.swapFour(16, 25, 34, 43)
			c.swapFour(17, 26, 35, 44)
		case L:
			c.swapFour(9, 11, 17, 15)
			c.swapFour(10, 14, 16, 12)
			c.swapFour(0, 18, 45, 44)
			c.swapFour(3, 21, 48, 41)
			c.swapFour(6, 24, 51, 38)
		case R:
			c.swapFour(27, 29, 35, 33)
			c.swapFour(28, 32, 34, 30)
			c.swapFour(20, 2, 42, 47)
			c.swapFour(23, 5, 39, 50)
			c.swapFour(26, 8, 36, 53)
		case B:
			c.swapFour(36, 38, 44, 42)
			c.swapFour(37, 41, 43, 39)
			c.swapFour(29, 0, 15, 53)
			c.swapFour(32, 1, 12, 52)
			c.swapFour(35, 2, 9, 51)
		case X:
			c.swapFour(0, 44, 45, 18)
			c.swapFour(1, 43, 46, 19)
			c.swapFour(2, 42, 47, 20)
			c.swapFour(3, 41, 48, 21)
			c.swapFour(4, 40, 49, 22)
			c.swapFour(5, 39, 50, 23)
			c.swapFour(6, 38, 51, 24)
			c.swapFour(7, 37, 52, 25)
			c.swapFour(8, 36, 53, 26)
			c.swapFour(27, 29, 35, 33)
			c.swapFour(28, 32, 34, 30)
			c.swapFour(11, 9, 15, 17)
			c.swapFour(14, 10, 12, 16)
		case Y:
			c.swapFour(44, 35, 26, 17)
			c.swapFour(43, 34, 25, 16)
			c.swapFour(42, 33, 24, 15)
			c.swapFour(41, 32, 23, 14)
			c.swapFour(40, 31, 22, 13)
			c.swapFour(39, 30, 21, 12)
			c.swapFour(38, 29, 20, 11)
			c.swapFour(37, 28, 19, 10)
			c.swapFour(36, 27, 18, 9)
			c.swapFour(0, 2, 8, 6)
			c.swapFour(1, 5, 7, 3)
			c.swapFour(47, 45, 51, 53)
			c.swapFour(50, 46, 48, 52)
		case Z:
			c.swapFour(0, 29, 53, 15)
			c.swapFour(1, 32, 52, 12)
			c.swapFour(2, 35, 51, 9)
			c.swapFour(3, 28, 50, 16)
			c.swapFour(4, 31, 49, 13)
			c.swapFour(5, 34, 48, 10)
			c.swapFour(6, 27, 47, 17)
			c.swapFour(7, 30, 46, 14)
			c.swapFour(8, 33, 45, 11)
			c.swapFour(18, 20, 26, 24)
			c.swapFour(19, 23, 25, 21)
			c.swapFour(36, 42, 44, 38)
			c.swapFour(39, 43, 41, 37)
		default:
			fmt.Printf("Unknown move: %v\n", m)
			return
		}
	}
	c.lastMoveTime = time.Now()
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

func (c *cube) LastMoveTime() time.Time {
	return c.lastMoveTime
}

//            0  1  2
//            3  4  5
//            6  7  8
// 9  10 11 | 18 19 20 | 27 28 29 | 36 37 38
// 12 13 14 | 21 22 23 | 30 31 32 | 39 40 41
// 15 16 17 | 24 25 26 | 33 34 35 | 42 43 44
//            45 46 47
//            48 49 50
//            51 52 53			c.swapFour(9, 11, 17, 15)

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

func (c cube) Solved() bool {
	return c.colors == solvedColors
}

