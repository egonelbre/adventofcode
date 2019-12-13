package main

import (
	"fmt"
	"os"

	"github.com/egonelbre/adventofcode/2019/day13/g"
	"github.com/egonelbre/adventofcode/2019/day13/intcode"
)

const (
	Empty  = g.Color(0)
	Wall   = g.Color(1)
	Block  = g.Color(2)
	Paddle = g.Color(3)
	Ball   = g.Color(4)
)

var colors = map[g.Color]rune{
	Empty:  ' ',
	Wall:   '█',
	Block:  '#',
	Paddle: '-',
	Ball:   'o',
}

func main() {
	var out int64
	var at g.Vector
	m := g.NewSparseImage(Empty)

	cpu := &intcode.Computer{
		Code: ArcadeCabinet.Clone(),
		Output: func(v int64) {
			switch out {
			case 0:
				at.X = v
			case 1:
				at.Y = v
			case 2:
				m.Set(at, g.Color(v))
			}

			out = (out + 1) % 3
		},
	}

	err := cpu.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fmt.Println("blocks", m.Count(Block))
	m.Image().Print(colors)
}
