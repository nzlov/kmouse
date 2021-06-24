package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

var (
	hold   = 0
	kmouse = false
	x, y   = 0, 0
	w, h   = 0, 0
	kw, kh = 0, 0
)

func main() {
	eChan := robotgo.Start()
	defer robotgo.End()

	for e := range eChan {
		switch e.Rawcode {
		case 59:
			if e.Kind == 4 {
				hold = hold | 1
			} else {
				hold = hold ^ 1
			}
		case 55:
			if e.Kind == 4 {
				hold = hold | 2
			} else {
				hold = hold ^ 2
			}
		case 39:
			if e.Kind == 5 {
				if hold == 3 {
					kmouse = !kmouse
					if kmouse {
						w, h = robotgo.GetScreenSize()
						x, y = 0, 0
						rset(w, h)
					}
				} else {
					w, h = robotgo.GetScreenSize()
					x, y = 0, 0
					rset(w, h)
				}
			}
		case 33:
			if e.Kind == 5 && kmouse {
				robotgo.MouseClick("left")
			}
		case 30:
			if e.Kind == 5 && kmouse {
				robotgo.MouseClick("right")
			}
		default:
			if kmouse && e.Kind == 5 {
				if pos, ok := keyM[e.Rawcode]; ok {
					mv(pos[0], pos[1])
				}
			}
		}
	}
}

func rset(w, h int) {
	fmt.Println("rset", w, h)

	kw = w / 10
	kh = h / 3

	robotgo.MoveMouseSmooth(x+w/2, y+h/2, 0.05, 1.0)
	fmt.Println("move_mouse_smooth")
}

// mv kx,ky为按键位置索引,q:0,0;p:9,0;z:2,0
func mv(kx, ky int) {
	fmt.Println("mv k:", kx, kw, ky, kh)
	x = x + kx*kw
	y = y + ky*kh
	fmt.Println("mv xy:", x, y)
	w, h = kw, kh
	rset(w, h)
}

var keyM = map[uint16][]int{
	12: {0, 0},
	13: {1, 0},
	14: {2, 0},
	15: {3, 0},
	17: {4, 0},
	16: {5, 0},
	32: {6, 0},
	34: {7, 0},
	31: {8, 0},
	35: {9, 0},

	0:  {0, 1},
	1:  {1, 1},
	2:  {2, 1},
	3:  {3, 1},
	5:  {4, 1},
	4:  {5, 1},
	38: {6, 1},
	40: {7, 1},
	37: {8, 1},
	41: {9, 1},

	6:  {0, 2},
	7:  {1, 2},
	8:  {2, 2},
	9:  {3, 2},
	11: {4, 2},
	45: {5, 2},
	46: {6, 2},
	43: {7, 2},
	47: {8, 2},
	44: {9, 2},
}
