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

	var best struct {
		X, Y  int64
		Score int64
	}

	for y := int64(0); y < space.Size.Y; y++ {
		for x := int64(0); x < space.Size.X; x++ {
			if space.At(Vector{x, y}) != Asteroid {
				continue
			}

			score := CountAsteroids(space, Vector{x, y})
			if score > best.Score {
				best.X, best.Y = x, y
				best.Score = score
			}
		}
	}

	fmt.Printf("%#v\n", best)
}

func CountAsteroids(space *Map, at Vector) int64 {
	counted := map[Vector]int64{}
	for y := int64(0); y < space.Size.Y; y++ {
		for x := int64(0); x < space.Size.X; x++ {
			loc := Vector{x, y}
			if loc == at {
				continue
			}
			if space.At(loc) != Asteroid {
				continue
			}

			offset := loc.Sub(at)
			direction := Direction(offset)

			counted[direction]++
		}
	}
	return int64(len(counted))
}

func Direction(offset Vector) Vector {
	gcd := GCD(offset.X, offset.Y)
	if gcd == 0 {
		return Vector{}
	}

	dir := offset
	dir.X /= gcd
	dir.Y /= gcd

	return dir
}

func GCD(a, b int64) int64 {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	if a == 0 {
		return b
	} else if b == 0 {
		return a
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
