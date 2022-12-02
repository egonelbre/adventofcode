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

	fmt.Println(guide.Score(Basic))
	fmt.Println(guide.Score(Advanced))

	return nil
}

func Basic(_ Weapon, r Rule) Weapon {
	switch r {
	case RuleX:
		return Rock
	case RuleY:
		return Paper
	case RuleZ:
		return Scissors
	}
	panic("invalid")
}

func Advanced(opponent Weapon, r Rule) Weapon {
	switch r {
	case RuleX:
		return opponent.WinsAgainst()
	case RuleY:
		return opponent
	case RuleZ:
		return opponent.LosesAgainst()
	}
	panic("invalid")
}

type StrategyGuide struct {
	Rounds []Round
}

type Strategy func(Weapon, Rule) Weapon

func (guide *StrategyGuide) Score(strategy Strategy) int {
	var total int
	for _, round := range guide.Rounds {
		total += round.Score(strategy)
	}
	return total
}

type Round struct {
	Opponent Weapon
	Rule     Rule
}

func (r Round) Score(strategy Strategy) int {
	player := strategy(r.Opponent, r.Rule)
	return score(r.Opponent, player)
}

func score(opponent, player Weapon) int {
	if opponent == player {
		return player.Score() + 3
	}

	if player.WinsAgainst() == opponent {
		return player.Score() + 6
	}

	return player.Score() + 0
}

type Weapon byte

const (
	Rock     Weapon = 1
	Paper    Weapon = 2
	Scissors Weapon = 3
)

func (weapon Weapon) LosesAgainst() Weapon {
	switch weapon {
	case Rock:
		return Paper
	case Paper:
		return Scissors
	case Scissors:
		return Rock
	}
	panic("invalid")
}

func (weapon Weapon) WinsAgainst() Weapon {
	switch weapon {
	case Rock:
		return Scissors
	case Paper:
		return Rock
	case Scissors:
		return Paper
	}
	panic("invalid")
}

func (weapon Weapon) Score() int { return int(weapon) }

type Rule int

const (
	RuleX Rule = 1
	RuleY Rule = 2
	RuleZ Rule = 3
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
			round.Rule = RuleX
		case "Y":
			round.Rule = RuleY
		case "Z":
			round.Rule = RuleZ
		default:
			return guide, fmt.Errorf("invalid line #%d %q", lineno, line)
		}

		guide.Rounds = append(guide.Rounds, round)
	}

	return guide, nil
}
