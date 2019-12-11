package main

import (
	"fmt"

	"github.com/egonelbre/adventofcode/2019/day11/intcode"
)

func main() {
	err := CountPanels()
	if err != nil {
		fmt.Println(err)
	}

	err = Paint()
	if err != nil {
		fmt.Println(err)
	}
}

func CountPanels() error {
	panels := map[Vector]int64{}
	err := RunRobot(panels)
	if err != nil {
		return err
	}

	fmt.Println("panels painted", len(panels))
	return nil
}

func Paint() error {
	panels := map[Vector]int64{}
	panels[Vector{0, 0}] = 1

	err := RunRobot(panels)
	if err != nil {
		return err
	}

	min, max := Vector{}, Vector{}
	for panel := range panels {
		min = min.Min(panel)
		max = max.Max(panel)
	}
	size := max.Sub(min).Add(Vector{1, 1})

	image := NewImage(size)
	for panel, color := range panels {
		image.Set(panel.Sub(min), Color(color))
	}

	image.Print()

	return nil
}

func RunRobot(panels map[Vector]int64) error {
	cpu := &intcode.Computer{
		Code: Painter.Clone(),
	}

	dir := Vector{0, -1}
	at := Vector{0, 0}

	for {
		color := panels[at]

		ok, err := WriteValue(cpu, color)
		if err != nil {
			return fmt.Errorf("failed to write: %w", err)
		}
		if !ok {
			return nil
		}

		paint, ok, err := ReadValue(cpu)
		if err != nil {
			return fmt.Errorf("failed to read paint: %w", err)
		}
		if !ok {
			return nil
		}

		panels[at] = paint

		rotate, ok, err := ReadValue(cpu)
		if err != nil {
			return fmt.Errorf("failed to read rotate: %w", err)
		}
		if !ok {
			return nil
		}

		if rotate == 0 {
			dir = RotateLeft(dir)
		} else {
			dir = RotateRight(dir)
		}

		at = at.Add(dir)
	}
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
