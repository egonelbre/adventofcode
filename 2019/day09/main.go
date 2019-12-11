package main

import (
	"fmt"

	"github.com/egonelbre/adventofcode/2019/day09/intcode"
)

func main() {
	BoostTestSequence()
	Part2()
}

func BoostTestSequence() {
	fmt.Println("=== PART 2 ===")

	var cpu *intcode.Computer
	cpu = &intcode.Computer{
		Input: func() int64 {
			return 1
		},
		Output: func(v int64) {
			fmt.Println("@", cpu.InstructionPointer, v)
		},
		Code: BOOST.Clone(),
	}

	err := cpu.Run()
	fmt.Println("@", cpu.InstructionPointer, "<finished>", err)
}

func Part2() {
}
