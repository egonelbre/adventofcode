package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"os"
)

func main() {
	flag.Parse()
	data := bytes.Split(must(os.ReadFile(flag.Arg(0))), []byte{'\n'})
	for i, line := range data {
		data[i] = bytes.TrimSpace(line)
	}

	first := true
	var bounds image.Rectangle

	for y, line := range data {
		for x, p := range line {
			if p != 'X' {
				continue
			}

			fmt.Println(x, y)

			r := image.Rect(x-1, y-1, x+1, y+1)
			if first {
				bounds = r
				first = false
			} else {
				bounds = bounds.Union(r)
			}
		}
	}

	fmt.Println(bounds)
	size := bounds.Size()
	fmt.Println(size, 2*size.X+2*size.Y)
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
