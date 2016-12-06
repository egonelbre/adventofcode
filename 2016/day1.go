package main

import (
	"fmt"
	"strconv"
	"strings"
)

type V2 struct {
	X, Y int
}

func (z *V2) Left()  { z.X, z.Y = -z.Y, z.X }
func (z *V2) Right() { z.Y, z.X = -z.X, z.Y }

func (z *V2) Move(v V2, n int) {
	z.X += v.X * n
	z.Y += v.Y * n
}

func Walk(input string) V2 {
	pos := V2{0, 0}
	dir := V2{0, 1}
	for _, cmd := range strings.Split(input, ", ") {
		if cmd[0] == 'L' {
			dir.Left()
		} else {
			dir.Right()
		}
		n, _ := strconv.Atoi(cmd[1:])
		pos.Move(dir, n)
	}
	return pos
}

const full = `L3, R2, L5, R1, L1, L2, L2, R1, R5, R1, L1, L2, R2, R4, L4, L3, L3, R5, L1, R3, L5, L2, R4, L5, R4, R2, L2, L1, R1, L3, L3, R2, R1, L4, L1, L1, R4, R5, R1, L2, L1, R188, R4, L3, R54, L4, R4, R74, R2, L4, R185, R1, R3, R5, L2, L3, R1, L1, L3, R3, R2, L3, L4, R1, L3, L5, L2, R2, L1, R2, R1, L4, R5, R4, L5, L5, L4, R5, R4, L5, L3, R4, R1, L5, L4, L3, R5, L5, L2, L4, R4, R4, R2, L1, L3, L2, R5, R4, L5, R1, R2, R5, L2, R4, R5, L2, L3, R3, L4, R3, L2, R1, R4, L5, R1, L5, L3, R4, L2, L2, L5, L5, R5, R2, L5, R1, L3, L2, L2, R3, L3, L4, R2, R3, L1, R2, L5, L3, R4, L4, R4, R3, L3, R1, L3, R5, L5, R1, R5, R3, L1`

func main() {
	fmt.Println(Walk(`R2, L3`))
	fmt.Println(Walk(`R2, R2, R2`))
	fmt.Println(Walk(`R5, L5, R5, R3`))
	fmt.Println(Walk(full))
}
