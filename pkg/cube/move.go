package cube

import (
  "strings"
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

