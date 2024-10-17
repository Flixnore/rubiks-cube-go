package cube

import (
  "fmt"
  "time"
)

func (q *cube) swapFour(a, b, c, d int) {
	temp := q.colors[d]
	q.colors[d] = q.colors[c]
	q.colors[c] = q.colors[b]
	q.colors[b] = q.colors[a]
	q.colors[a] = temp
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
