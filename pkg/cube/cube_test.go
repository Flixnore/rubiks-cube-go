package cube

import (
  "testing"
)

func TestBasic(t *testing.T) {
  c := NewSolved()
  for _, move := range ParseMoves("F F' U U' R R' L L' B B' D D'") {
    c.Rotate(move)
  }
  for _, move := range ParseMoves("F F F F D D D D") {
    c.Rotate(move)
  }

  if !c.Solved() {
    t.Errorf("Expected solved cube, got:\n%v", c)
  }
}

func TestScrambleAndSolve(t *testing.T) {
  c := NewSolved()
  for _, move := range ParseMoves("U' L2 F2 D' L' D U2 R U' R' U2 R2 U F' L' U R'") {
    c.Rotate(move)
  }

  for _, move := range ParseMoves("R' U' F2 D F' D' B L F U B' R B2 L2 D2 R2 D R2 B2") {
    c.Rotate(move)
  }

  if !c.Solved() {
    t.Errorf("Expected solved cube, got:\n%v", c)
  }
}

func TestUndo(t *testing.T) {
  c := NewSolved()
  for _, move := range ParseMoves("U' L2 F2 D' L' D U2 R U' R' U2 R2 U F' L' U R'") {
    c.Rotate(move)
  }

  for range len(c.History()) {
    c.Undo()
  }

  if !c.Solved() {
    t.Errorf("Expected solved cube, got:\n%v", c)
  }
}

func TestRedo(t *testing.T) {
  c := NewSolved()
  moveStr := "U' L2 F2 D' L' D U2 R U' R' U2 R2 U F' L' U R'"
  for _, move := range ParseMoves(moveStr) {
    c.Rotate(move)
  }

  preState := c.State()

  for range len(moveStr) {
    c.Undo()
  }

  for range len(moveStr) {
    c.Redo()
  }
  
  postState := c.State()

  if preState != postState {
    t.Errorf("Expected pre undo-redo state to match post state, got:\n%v", c)
  }
}
