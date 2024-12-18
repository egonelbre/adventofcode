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
	screen := bytes.Split(data, []byte{'\n'})

	for y := 0; y < len(screen); y += 4 {
		k := 1
		first := true
		fmt.Print("|")
		for x := 0; x < len(screen[y]); x += 5 {
			ok1 := bytes.Equal(screen[y][x:x+4], []byte("|==|"))
			ok2 := bytes.Equal(screen[y+1][x:x+4], []byte("|  |"))
			ok3 := bytes.Equal(screen[y+2][x:x+4], []byte("|==|"))

			if !ok1 || !ok2 || !ok3 {
				if !first {
					fmt.Print(", ")
				}
				fmt.Print(k)
				first = false
			}
			k++
		}
		fmt.Print("|")
	}
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
