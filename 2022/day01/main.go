package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

const input1 = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

type Elf struct {
	Food  []Food
	Total Energy
}

func (elf *Elf) Init() {
	for _, food := range elf.Food {
		elf.Total += food.Calories
	}
}

type Food struct {
	Calories Energy
}

type Energy int64

func ParseInput(input string) ([]Elf, error) {
	input = strings.TrimSpace(input) + "\n"
	elves := []Elf{}

	var elf Elf
	for lineno, line := range strings.Split(input, "\n") {
		if line == "" {
			elf.Init()
			elves = append(elves, elf)
			elf = Elf{}
			continue
		}
		calories, err := strconv.Atoi(line)
		if err != nil {
			return elves, fmt.Errorf("invalid line %d %q: %w", lineno, line, err)
		}
		elf.Food = append(elf.Food, Food{
			Calories: Energy(calories),
		})
	}

	return elves, nil
}

func SumTopN(elves []Elf, n int) Energy {
	slices.SortFunc(elves, func(a, b Elf) bool {
		return a.Total > b.Total
	})

	var total Energy
	for _, elf := range elves[:n] {
		total += elf.Total
	}
	return total
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	data, err := os.ReadFile("input1.txt")
	if err != nil {
		return fmt.Errorf("loading input failed: %w", err)
	}

	elves, err := ParseInput(string(data))
	if err != nil {
		return fmt.Errorf("parsing input failed: %w", err)
	}

	fmt.Println(SumTopN(elves, 1))
	fmt.Println(SumTopN(elves, 3))

	return nil
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
