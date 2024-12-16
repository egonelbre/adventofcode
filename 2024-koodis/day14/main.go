package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
)

const (
	Floor     = '_'
	Table     = 'L'
	Door      = 'U'
	Furniture = 'O'
)

func canPlaceNear(b byte) bool {
	return b != Door && b != Furniture
}

func main() {
	flag.Parse()
	data := must(os.ReadFile(flag.Arg(0)))

	plan := bytes.Split(data, []byte{'\n'})
	count := 0
	for y, line := range plan {
		for x, at := range line {
			if at != Floor {
				continue
			}

			a, b, c, d := line[x-1], line[x+1], plan[y-1][x], plan[y+1][x]

			if !canPlaceNear(a) || !canPlaceNear(b) || !canPlaceNear(c) || !canPlaceNear(d) {
				continue
			}
			if a != Table && b != Table && c != Table && d != Table {
				continue
			}

			plan[y][x] = 'X'
			count++
		}
	}

	for _, line := range plan {
		fmt.Println(string(line))
	}

	fmt.Println(count)
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
