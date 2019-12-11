package main

import (
	"fmt"
	"math"

	"github.com/egonelbre/adventofcode/2019/day11/intcode"
)

func main() {
	err := CountPanels()
	if err != nil {
		fmt.Println(err)
	}
}

func CountPanels() error {
	cpu := &intcode.Computer{
		Code: Painter.Clone(),
	}

	dir := Vector{0, -1}
	at := Vector{0, 0}
	colors := map[Vector]int64{}

	for {
		color := colors[at]

		ok, err := WriteValue(cpu, color)
		if err != nil {
			return fmt.Errorf("failed to write: %w", err)
		}
		if !ok {
			break
		}

		paint, ok, err := ReadValue(cpu)
		if err != nil {
			return fmt.Errorf("failed to read paint: %w", err)
		}
		if !ok {
			break
		}

		colors[at] = paint

		rotate, ok, err := ReadValue(cpu)
		if err != nil {
			return fmt.Errorf("failed to read rotate: %w", err)
		}
		if !ok {
			break
		}

		if rotate == 0 {
			dir = RotateLeft(dir)
		} else {
			dir = RotateRight(dir)
		}

		at = at.Add(dir)
	}

	fmt.Println("panels painted", len(colors))
	return nil
}

func WriteValue(cpu *intcode.Computer, value int64) (ok bool, err error) {
	cpu.Halted = false

	cpu.Input = func() int64 {
		cpu.Halted = true
		ok = true
		return value
	}
	cpu.Output = nil

	err = cpu.Run()
	return ok, err
}

func ReadValue(cpu *intcode.Computer) (output int64, ok bool, err error) {
	cpu.Halted = false

	cpu.Input = nil
	cpu.Output = func(v int64) {
		cpu.Halted = true
		output = v
		ok = true
	}

	err = cpu.Run()
	return output, ok, err
}

type Vector struct {
	X, Y int64
}

func (a Vector) IsZero() bool {
	return a == Vector{}
}

func (a Vector) Add(b Vector) Vector {
	return Vector{a.X + b.X, a.Y + b.Y}
}
func (a Vector) Sub(b Vector) Vector {
	return Vector{a.X - b.X, a.Y - b.Y}
}
func (a Vector) SquareLength() int64 {
	return a.X*a.X + a.Y*a.Y
}
func (a Vector) Angle() float64 {
	tx := float64(a.X)
	ty := float64(a.Y)
	return math.Atan2(-tx, ty)
}

func RotateLeft(v Vector) Vector {
	return Vector{v.Y, -v.X}
}

func RotateRight(v Vector) Vector {
	return Vector{-v.Y, v.X}
}
