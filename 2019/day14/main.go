package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	reactions, err := Parse(InputReactions)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	state := map[string]int64{
		"FUEL": 1,
	}
	reactions.Reduce(state)
	fmt.Println("1 => ", state)

	fuel := FindMaximumFuel(reactions, 1000000000000)
	fmt.Println(fuel)

	state = map[string]int64{
		"FUEL": fuel,
	}
	reactions.Reduce(state)
	fmt.Println(fuel, " => ", state)
}

func FindMaximumFuel(reactions *Reactions, ore int64) int64 {
	return int64(sort.Search(int(ore), func(fuel int) bool {
		state := map[string]int64{
			"FUEL": int64(fuel),
		}
		reactions.Reduce(state)
		return state["ORE"] >= ore
	})) - 1
}
