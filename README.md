# Rubik's Cube Go Simulator

This is a Rubik's Cube simulator written in Go.

## Usage

### Terminal

Run this to enter a CLI rubik's cube. You can enter multiple moves in [standard Rubik's Cube notation](https://jperm.net/3x3/moves).

```
go run ./cmd/terminal
```

### Flat

This version uses the ebiten game engine to render a flat cube similar to the terminal output.

```
go run ./cmd/flat
```

Controls:

- `J K L ;` for the 4 vertical sides and `I` and `,` for the horizontal sides.
- `u` and `Ctrl + r` for undo and redo respectively.
