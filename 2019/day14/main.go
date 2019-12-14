package main

import (
	"fmt"
	"os"
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
	fmt.Println(state)
}
