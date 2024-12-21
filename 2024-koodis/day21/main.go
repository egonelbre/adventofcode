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

	plan := bytes.Split(data, []byte("\n"))

	for k := 0; ; k++ {
		changed := flood(plan)
		if !changed {
			break
		}

		// fmt.Println(k)
		// printplan(plan)
	}

	fmt.Println("=== DONE ===")
	printplan(plan)

	for {
		changed1 := floodBlack(plan)
		changed2 := flood(plan)
		if !(changed1 || changed2) {
			break
		}
	}

	fmt.Println("=== DONE ===")
	printplan(plan)
}

func flood(data [][]byte) (changed bool) {
	for y := range data {
		for x, v := range data[y] {
			if v == '*' || v == '0' {
				continue
			}

			fleft, cleft := status(data, x-1, y)
			fright, cright := status(data, x+1, y)
			ftop, ctop := status(data, x, y-1)
			fbottom, cbottom := status(data, x, y+1)

			total := fleft + fright + ftop + fbottom
			if total == 0 {
				continue
			}

			closed := cleft + cright + ctop + cbottom
			if closed == 4 {
				data[y][x] = '0'
				continue
			}

			switch {
			case v == '4':
				data[y][x] = '*'
				changed = true
				continue
			case total == 4 && (v == '1' || v == '2' || v == '3'):
				data[y][x] = '*'
				changed = true
				continue
			case total == 3 && (v == '3' || v == '2'):
				data[y][x] = '*'
				changed = true
				continue
			case total >= 2 && v == '3':
				data[y][x] = '*'
				changed = true
				continue
			}
		}
	}
	return changed
}

func floodBlack(data [][]byte) (changed bool) {
	for y := range data {
		for x, v := range data[y] {
			if v == '*' || v == '0' {
				continue
			}

			fleft, _ := status(data, x-1, y)
			fright, _ := status(data, x+1, y)
			ftop, _ := status(data, x, y-1)
			fbottom, _ := status(data, x, y+1)

			total := fleft + fright + ftop + fbottom
			if total > 0 {
				data[y][x] = '*'
				changed = true
			}
		}
	}
	return changed
}

func status(data [][]byte, x, y int) (flood int, closed int) {
	if y < 0 || y >= len(data) {
		return 0, 1
	}
	line := data[y]
	if x < 0 || x >= len(line) {
		return 0, 1
	}
	v := line[x]
	if v == '*' {
		return 1, 0
	}
	if v == '0' {
		return 0, 1
	}
	return 0, 0
}

func printplan(data [][]byte) {
	flooded := 0
	for _, line := range data {
		fmt.Println(string(line))
		for _, v := range line {
			if v == '*' {
				flooded++
			}
		}
	}
	fmt.Println("flooded:", flooded)
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
