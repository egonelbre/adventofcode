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
	for _, x := range Primes {
		for _, y := range Primes {
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

var Primes = []int64{1, 2, 3, 5, 7, 11, 13, 17, 19}
