package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	data := must(os.ReadFile(flag.Arg(0)))
	data = bytes.ReplaceAll(data, []byte{' '}, nil)

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
