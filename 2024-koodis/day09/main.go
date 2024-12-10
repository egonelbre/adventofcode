package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	data := must(os.ReadFile(flag.Arg(0)))
	matrix := bytes.Split(data, []byte{'\n'})
	for i, line := range matrix {
		line = bytes.TrimSpace(line)
		for k, b := range line {
			line[k] = b - '0'
		}
		matrix[i] = line
	}

	width, height := len(matrix[0]), len(matrix)

	hole := 0
	for y := 0; y < height-1; y++ {
		for x := 0; x < width-1; x++ {
			total := matrix[y][x] + matrix[y][x+1] + matrix[y+1][x] + matrix[y+1][x+1]
			if total == 4 {
				hole++
			}
		}
	}
	fmt.Println(hole)
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
