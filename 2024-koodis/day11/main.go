package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	data := string(must(os.ReadFile(flag.Arg(0))))

	blocks := strings.Split(data, "\n\n")
	
	friends := strings.Split(blocks[0], "\n")[1:]
	nolikes := strings.Split(blocks[1], "\n")[1:]
	notinvited := strings.Split(blocks[2], "\n")[1:]

	dont := map[string]bool {}
	for _, friend := range notinvited {
		dont[friend] = true
	}
	dislike := map[string][]string{}
	for _, nolike := range nolikes {
		person, people, _ := strings.Cut(nolike, ": ")
		dislike[person] = strings.Split(people, ", ")
	}

	invited := map[string]bool{}
	order :=[]string{}

	next:
	for _, friend := range friends {
		if dont[friend] {
			continue next
		}

		for _, dis := range dislike[friend] {
			if invited[dis] {
				continue next
			}
		}

		order = append(order, friend)
		invited[friend] = true

		for _, dis := range dislike[friend] {
			dont[dis] = true
		}
	}

	fmt.Println(strings.Join(order, ", "))
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
