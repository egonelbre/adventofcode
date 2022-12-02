package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	input, err := os.ReadFile("input1.txt")
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	guide, err := ParseStrategyGuide(string(input))
	if err != nil {
		return fmt.Errorf("failed to parse guide: %w", err)
	}

	fmt.Println(guide.Score())

	return nil
}

type StrategyGuide struct {
	Rounds []Round
}

func (guide *StrategyGuide) Score() int {
	var total int
	for _, round := range guide.Rounds {
		total += round.Score()
	}
	return total
}

type Round struct {
	Opponent Weapon
	Response Weapon
}

func (r Round) Score() int {
	if r.Opponent == r.Response {
		return r.Response + 3
	}

	switch r {
	case Round{Rock, Paper},
		Round{Paper, Scissors},
		Round{Scissors, Rock}:
		return r.Response + 6
	case Round{Paper, Rock},
		Round{Scissors, Paper},
		Round{Rock, Scissors}:
		return r.Response + 0
	}

	panic("invalid")
}

type Weapon = int

const (
	Rock     Weapon = 1
	Paper    Weapon = 2
	Scissors Weapon = 3
)

func ParseStrategyGuide(input string) (StrategyGuide, error) {
	input = strings.TrimSpace(input)

	guide := StrategyGuide{}
	for lineno, line := range strings.Split(input, "\n") {
		op, ans, found := strings.Cut(line, " ")
		if !found {
			return guide, fmt.Errorf("invalid line #%d %q", lineno, line)
		}

		round := Round{}
		switch op {
		case "A":
			round.Opponent = Rock
		case "B":
			round.Opponent = Paper
		case "C":
			round.Opponent = Scissors
		default:
			return guide, fmt.Errorf("invalid line #%d %q", lineno, line)
		}

		switch ans {
		case "X":
			round.Response = Rock
		case "Y":
			round.Response = Paper
		case "Z":
			round.Response = Scissors
		default:
			return guide, fmt.Errorf("invalid line #%d %q", lineno, line)
		}

		guide.Rounds = append(guide.Rounds, round)
	}

	return guide, nil
}
