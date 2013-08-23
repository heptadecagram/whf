package main

import (
	"math/rand"
	"time"
)

import "github.com/nsf/termbox-go"

type Graphic struct {
	color termbox.Attribute
	image rune
	size  int
}

type coord struct {
	y, x int
}

func StatusMsg(m string) {
	for i, r := range m {
		termbox.SetCell(i+1, 0, r, 0, 0)
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	rand.Seed(time.Now().UnixNano())

	sy, sx := termbox.Size()
	pcolor := termbox.ColorDefault


	my, mx := int(rand.Int31n(int32(sy))), int(rand.Int31n(int32(sx-1)))+1
	termbox.SetCell(my, mx, 'y', 0, 0)
	ty, tx := int(rand.Int31n(int32(sy))), int(rand.Int31n(int32(sx-1)))+1
	dead := false

	y, x := 6, 5
	termbox.SetCell(y, x, '@', pcolor, 0)
	termbox.Flush()

	e := termbox.PollEvent()
	for {
		switch e.Type {
		case termbox.EventResize:
			sy, sx = termbox.Size()
		case termbox.EventKey:
			termbox.SetCell(y, x, ' ', 0, 0)
			switch e.Ch {
			case 0:
				if my == y && mx == x {
					dead = true
					StatusMsg("You got  it!")
				} else if !dead {
					termbox.SetCell(my, mx, ' ', 0, 0)

					if ty < my {
						my--
					} else if my < ty {
						my++
					} else if tx < mx {
						mx--
					} else if mx < tx {
						mx++
					}
					if my == ty && mx == tx {
						ty, tx = int(rand.Int31n(int32(sy))), int(rand.Int31n(int32(sx-1)))+1
					}

					termbox.SetCell(my, mx, 'y', 0, 0)
				}
				switch e.Key {
				case termbox.KeyEsc:
					return
				case termbox.KeyArrowUp:
					x--
					if x < 1 {
						x = 1
					}
				case termbox.KeyArrowDown:
					x++
					if x >= sx {
						x = sx - 1
					}
				case termbox.KeyArrowLeft:
					y--
					if y < 0 {
						y = 0
					}
				case termbox.KeyArrowRight:
					y++
					if y >= sy {
						y = sy - 1
					}
				}
			case 'y':
				pcolor = termbox.ColorYellow
			case 'g':
				pcolor = termbox.ColorGreen
			case 'b':
				pcolor = termbox.ColorBlue
			case 'r':
				pcolor = termbox.ColorRed
			case 'p':
				pcolor = termbox.ColorMagenta
			case 's':
				pcolor = termbox.ColorWhite
			case 'w':
				pcolor = termbox.ColorDefault
			case 'q':
				return
			}
			if my == y && mx == x {
				dead = true
				StatusMsg("You got  it!")
			}

			termbox.SetCell(y, x, '@', pcolor, 0)
		}

		if dead {
			sy, sx = termbox.Size()
			sx /= 2
			sy /= 2

			var endmsg string
			endmsg = "YOU WIN"

			sy -= len(endmsg) / 2

			for i, r := range endmsg {
				termbox.SetCell(sy+i, sx, r, 0, 0)
			}
			termbox.Flush()
			for {
				e = termbox.PollEvent()

				if e.Ch == 'q' || e.Key == termbox.KeySpace {
					return
				}
			}
		}
		termbox.Flush()
		e = termbox.PollEvent()
	}

}
