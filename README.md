# Rubik's Cube Go Simulator

This is a Rubik's Cube simulator written in Go.

## Usage

### Terminal

Run this to enter a CLI rubik's cube. You can enter multiple moves in [standard Rubik's Cube notation](https://jperm.net/3x3/moves).

```
go run ./cmd/terminal
```

![Peek 2024-11-11 13-56](https://github.com/user-attachments/assets/e4dcf78d-412c-4c1b-a228-0b15fb555f87)

### Flat

This version uses the ebiten game engine to render a flat cube similar to the terminal output.

```
go run ./cmd/flat
```

Controls:

- `J K L ;` for the 4 vertical sides and `I` and `,` for the horizontal sides. `Shift + letter` will do the rotation counterclockwise.
- `u` and `Ctrl + r` for undo and redo respectively.
- `z x c` for cube rotations. `c` takes the place of the `y` rotation because of convenient keyboard positioning.
