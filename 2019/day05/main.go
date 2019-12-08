package main

import (
	"fmt"

	"github.com/egonelbre/adventofcode/2019/day05/intcode"
)

func main() {
	Part1()
	Part2()
}

func Part1() {
	fmt.Println("=== PART 1 ===")

	var cpu *intcode.Computer
	cpu = &intcode.Computer{
		Input: func() int64 {
			// air condition unit
			return 1
		},
		Output: func(v int64) {
			fmt.Println("@", cpu.InstructionPointer, v)
		},

		Code: Input.Clone(),
	}

	err := cpu.Run()
	fmt.Println("@", cpu.InstructionPointer, "<finished>", err)
}

func Part2() {
	fmt.Println("=== PART 2 ===")

	var cpu *intcode.Computer
	cpu = &intcode.Computer{
		Input: func() int64 {
			// thermal radiator controller
			return 5
		},
		Output: func(v int64) {
			fmt.Println("@", cpu.InstructionPointer, v)
		},

		Code: Input.Clone(),
	}

	err := cpu.Run()
	fmt.Println("@", cpu.InstructionPointer, "<finished>", err)
}
