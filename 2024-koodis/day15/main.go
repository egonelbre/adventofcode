package main

import (
	"image"
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
	code := plan[0]
	plan = plan[1:]

	locations := map[byte]image.Point{}

	for y, line := range plan {
		for x, at := range line {
			if at == 'x' || at == ' ' {
				continue
			}

			locations[at] = image.Pt(x, y)
		}
	}

	start := locations['-']
	code = append(code, '+')
	distance := 0
	for _, shop := range code {
		at, ok := locations[shop]
		if !ok {
			panic("shop")
		}
		dist := at.Sub(start)
		man := abs(dist.X) + abs(dist.Y)
		fmt.Printf("%v %v : %v %v\n", string(shop), at, dist, man)
		distance += man
		start = at
	}

	fmt.Println(distance)
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
