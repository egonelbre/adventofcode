package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	space, err := ParseMap(Input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse %v\n", err)
		os.Exit(1)
	}

	var best struct {
		Location Vector
		Score    int64
	}

	for y := int64(0); y < space.Size.Y; y++ {
		for x := int64(0); x < space.Size.X; x++ {
			loc := Vector{x, y}
			if space.At(loc) != Asteroid {
				continue
			}

			score := CountAsteroids(space, loc)
			if score > best.Score {
				best.Location = loc
				best.Score = score
			}
		}
	}

	fmt.Printf("%#v\n", best)

	ShootingGallery(space, best.Location, 200)
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

func ShootingGallery(space *Map, at Vector, count int64) {
	type Target struct {
		Location Vector
		Offset   Vector
	}
	type Targets struct {
		Direction Vector
		Targets   []Target
	}

	targets := []*Targets{}
	byDirection := map[Vector]*Targets{}
	for y := int64(0); y < space.Size.Y; y++ {
		for x := int64(0); x < space.Size.X; x++ {
			location := Vector{x, y}
			if location == at {
				continue
			}
			if space.At(location) != Asteroid {
				continue
			}

			offset := location.Sub(at)
			direction := Direction(offset)

			tx, ok := byDirection[direction]
			if !ok {
				tx = &Targets{
					Direction: direction,
				}
				targets = append(targets, tx)
				byDirection[direction] = tx
			}
			tx.Targets = append(tx.Targets, Target{
				Location: location,
				Offset:   offset,
			})

			byDirection[direction].Targets = append(byDirection[direction].Targets)
		}
	}

	for _, tx := range targets {
		sort.Slice(tx.Targets, func(i, k int) bool {
			di := tx.Targets[i].Offset.SquareLength()
			dk := tx.Targets[k].Offset.SquareLength()
			return di < dk
		})
	}
	sort.Slice(targets, func(i, k int) bool {
		ai := targets[i].Direction.Angle()
		ak := targets[k].Direction.Angle()
		return ai < ak
	})

	asteroid := 0
	k := 0
	for len(targets) > 0 {
		tx := targets[k]

		var target Vector
		target, tx.Targets = tx.Targets[0].Location, tx.Targets[1:]

		asteroid++
		fmt.Println("#", asteroid, target)

		if len(tx.Targets) == 0 {
			targets = append(targets[:k], targets[k+1:]...)
		} else {
			k++
		}
		if k >= len(targets) {
			k = 0
		}

	}
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
