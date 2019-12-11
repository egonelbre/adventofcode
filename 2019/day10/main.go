package main

import (
	"fmt"
	"os"
)

func main() {
	space, err := ParseMap(Input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse %v\n", err)
		os.Exit(1)
	}

	offsets := []Vector{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	for y := int64(1); y < space.Size.Y; y++ {
		for x := int64(1); x < space.Size.X; x++ {
			if GCD(x, y) != 1 {
				continue
			}

			offsets = append(offsets,
				Vector{x, y},
				Vector{x, -y},
				Vector{-x, y},
				Vector{-x, -y},
			)
		}
	}

	var best struct {
		X, Y  int64
		Score int64
	}

	for y := int64(0); y < space.Size.Y; y++ {
		for x := int64(0); x < space.Size.X; x++ {
			if space.At(Vector{x, y}) != Asteroid {
				continue
			}

			score := CountAsteroids(space, Vector{x, y}, offsets)
			if score > best.Score {
				best.X, best.Y = x, y
				best.Score = score
			}
		}
	}

	fmt.Printf("%#v\n", best)
}

func CountAsteroids(space *Map, at Vector, offsets []Vector) int64 {
	var count int64
	for _, offset := range offsets {
		see := at.Add(offset)
		for ; space.Contains(see); see = see.Add(offset) {
			tile := space.At(see)
			if tile == Asteroid {
				count++
				break
			}
		}
	}
	return count
}

func GCD(a, b int64) int64 {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
	return a
}
