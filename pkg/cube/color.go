package cube

import (
	colores "image/color"
)

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
