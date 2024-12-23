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

	ok := 0
	for _, line := range bytes.Split(data, []byte("\n")) {
		line = bytes.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		if valid(line) {
			ok++
		}
	}
	fmt.Println(ok)
}

func valid(v []byte) bool {
	stack := 0
	for _, b := range v {
		if b == '(' {
			stack++
		} else if b == ')' {
			stack--
			if stack < 0 {
				return false
			}
		} else {
			panic("invalid input " + string(v))
		}
	}
	return stack == 0
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
