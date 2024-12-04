package main

import (
	"flag"
	"fmt"
	"os"
	"unicode"
)

const AlphabetData = `ABCDEFGHIJKLMNOPRSŠZŽTUVÕÄÖÜ`

func main() {
	flag.Parse()
	data := string(must(os.ReadFile(flag.Arg(0))))

	ALPHABET := []rune(AlphabetData)

	index := map[rune]int{}
	for i, a := range ALPHABET {
		index[a] = i
	}

	for rot := range len(ALPHABET) - 1 {
		fmt.Printf("%2d: ", rot)
		for _, b := range data {
			isup := unicode.IsUpper(b)
			if !isup {
				b = unicode.ToUpper(b)
			}

			ix, ok := index[b]
			if ok {
				b = ALPHABET[(ix+rot)%len(ALPHABET)]
			}

			if !isup {
				b = unicode.ToLower(b)
			}

			fmt.Print(string(b))
		}
		fmt.Println()
	}

	location := 1
	level := 1
	for _, b := range data {
		switch b {
		case 'v':
			location = location*2 - 1
		case 'p':
			location = location * 2
		}
		level++
	}

	fmt.Println(location)
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

/*
                       1
                 /           \
           1                       2
        /     \                 /     \
     1           2           3           4
    / \         / \         / \         / \
  1     2     3     4     5     6     7     8
 / \   / \   / \   / \   / \   / \   / \   / \
1   2 3   4 5   6 7   8 9  10 11 12 13 14 15 16
*/
