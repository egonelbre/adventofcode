package main

import (
	"fmt"
	"os"

	"github.com/egonelbre/adventofcode/2019/day17/g"
	"github.com/egonelbre/adventofcode/2019/day17/intcode"
)

const (
	Unknown  = g.Color(0)
	Empty    = g.Color('.')
	Scaffold = g.Color('#')

	RobotUp       = g.Color('^')
	RobotDown     = g.Color('v')
	RobotLeft     = g.Color('<')
	RobotRight    = g.Color('>')
	RobotTumbling = g.Color('X')
)

var colors = map[g.Color]rune{
	Unknown:       ' ',
	Empty:         '.',
	Scaffold:      '#',
	RobotUp:       '^',
	RobotDown:     'v',
	RobotLeft:     '<',
	RobotRight:    '>',
	RobotTumbling: 'X',
}

func HasScaffolding(tile g.Color) bool {
	return tile == Scaffold ||
		tile == RobotUp ||
		tile == RobotDown ||
		tile == RobotLeft ||
		tile == RobotRight
}

func main() {
	world, err := ScanWorld()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	world.Image().Print(colors)
	ComputeAlignment(world)
}

func ComputeAlignment(world *g.SparseImage) {
	var totalAlignment int64
	for p, tile := range world.Data {
		if !HasScaffolding(tile) {
			continue
		}

		surrounded := HasScaffolding(world.At(p.Add(g.Vector{0, -1}))) &&
			HasScaffolding(world.At(p.Add(g.Vector{0, +1}))) &&
			HasScaffolding(world.At(p.Add(g.Vector{-1, 0}))) &&
			HasScaffolding(world.At(p.Add(g.Vector{+1, 0})))
		if !surrounded {
			continue
		}

		totalAlignment += p.X * p.Y
	}
	fmt.Println("total alignment", totalAlignment)
}

func ScanWorld() (*g.SparseImage, error) {
	m := g.NewSparseImage(Unknown)
	cursor := g.Vector{}
	cpu := &intcode.Computer{
		Code: ASCIIProgram.Clone(),
		Output: func(v int64) {
			if v == '\n' {
				cursor.X = 0
				cursor.Y++
				return
			}
			m.Set(cursor, g.Color(v))
			cursor.X++
		},
	}

	err := cpu.Run()
	return m, err
}
