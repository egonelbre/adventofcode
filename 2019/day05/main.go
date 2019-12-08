package main

import (
	"fmt"

	"github.com/egonelbre/adventofcode/2019/day05/intcode"
)

func main() {
	const AirConditionUnitID = 1

	var cpu *intcode.Computer
	cpu = &intcode.Computer{
		Input: func() int64 {
			return AirConditionUnitID
		},
		Output: func(v int64) {
			fmt.Println("@", cpu.InstructionPointer, v)
		},

		Code: Input,
	}

	err := cpu.Run()
	fmt.Println("@", cpu.InstructionPointer, "<finished>", err)
}
