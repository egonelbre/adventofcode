package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"os"
	"slices"
)

const (
	Wall    = '#'
	Start   = 'L'
	End     = 'K'
	Empty   = ' '
	Visited = '+'
)

func main() {
	flag.Parse()
	data := must(os.ReadFile(flag.Arg(0)))

	plan := bytes.Split(data, []byte{'\n'})

	var parent [][]image.Point
	for _, line := range plan {
		parent = append(parent, make([]image.Point, len(line)))
	}

	start := image.Point{}
	end := image.Point{}
	for y, line := range plan {
		for x, at := range line {
			if at == Start {
				start = image.Pt(x, y)
			}
			if at == End {
				end = image.Pt(x, y)
			}
			if at == '.' {
				line[x] = Empty
			}
		}
	}

	fmt.Printf("start:%v end:%v\n", start, end)

	queue := []image.Point{start}
	next := []image.Point{}
	found := false

	visit := func(from, to image.Point, direction byte) {
		place := plan[to.Y][to.X]

		if place == End {
			found = true
			plan[to.Y][to.X] = direction
			parent[to.Y][to.X] = from
			return
		}
		if place != Empty {
			return
		}

		plan[to.Y][to.X] = direction
		parent[to.Y][to.X] = from
		next = append(next, to)
	}

	for len(queue) > 0 && !found {
		for _, at := range queue {
			visit(at, at.Add(image.Point{X: 0, Y: 1}), '^')
			visit(at, at.Add(image.Point{X: 0, Y: -1}), 'v')
			visit(at, at.Add(image.Point{X: 1, Y: 0}), '<')
			visit(at, at.Add(image.Point{X: -1, Y: 0}), '>')
			if found {
				break
			}
		}
		queue, next = next, queue[:0]
	}

	for _, line := range plan {
		fmt.Println(string(line))
	}

	if !found {
		panic("did not find end")
	}

	path := []image.Point{end}
	at := end
	for at != start {
		if at == (image.Point{}) {
			panic("wrong")
		}
		at = parent[at.Y][at.X]
		path = append(path, at)
	}

	slices.Reverse(path)
	fmt.Println(path)

	heading := path[1].Sub(path[0])
	fmt.Println(heading)
	last := path[0]
	// VVPVP

	for _, at := range path[1:] {
		move := at.Sub(last)
		last = at
		if move == heading {
			continue
		}
		rotate := heading.X*move.Y - heading.Y*move.X
		heading = move
		if rotate == -1 {
			fmt.Print("V")
		} else {
			fmt.Print("P")
		}
	}
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
