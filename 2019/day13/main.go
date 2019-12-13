package main

import (
	"fmt"
	"os"
	"time"

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
	CountBlocks()
	PlayGame()
}

func CountBlocks() {
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

func PlayGame() {
	SegmentDisplay := g.Vector{-1, 0}

	var out int64
	var at g.Vector

	var score int64
	display := g.NewSparseImage(Empty)

	cpu := &intcode.Computer{
		Code: ArcadeCabinet.Clone(),
		Input: func() int64 {
			ball, okb := display.Find(Ball)
			paddle, okp := display.Find(Paddle)
			if !okb || !okp {
				fmt.Fprintln(os.Stderr, "unable to find ball/paddle")
				display.Image().Print(colors)
				return 0
			}

			if false {
				print("\033[H\033[2J")
				fmt.Println("Score", score)
				display.Image().Print(colors)
				time.Sleep(30 * time.Millisecond)
			}

			return ball.Sub(paddle).Sign().X
		},
		Output: func(v int64) {
			switch out {
			case 0:
				at.X = v
			case 1:
				at.Y = v
			case 2:
				if at == SegmentDisplay {
					score = v
				} else {
					display.Set(at, g.Color(v))
				}
			}

			out = (out + 1) % 3
		},
	}
	cpu.Code[0] = 2

	err := cpu.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fmt.Println("score", score)
	display.Image().Print(colors)
}
