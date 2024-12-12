package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	flag.Parse()
	data := string(must(os.ReadFile(flag.Arg(0))))

	database, input, _ := strings.Cut(data, "\n\n")

	type Category struct {
		Height string
		Age string
		Cloth string
	}

	allowedPeople := map[Category]bool{}
	for _, line := range strings.Split(database, "\n") {
		line = strings.TrimSpace(line)
		if line == "" { continue }

		tokens := strings.Split(line, "-")
		allowedPeople[Category{
			Height: tokens[1],
			Age: tokens[2],
			Cloth: tokens[3],
		}] = true
	}

	rx := regexp.MustCompile(`(\d+) cm, (\d+) aastat vana, seljas on (.*)`)

	allowed, disallowed := 0, 0
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if line == "" { continue }

		match := rx.FindStringSubmatch(line)

		cat := Category{
			Height: match[1],
			Age: match[2],
			Cloth: match[3],
		}
		fmt.Println(cat)

		if allowedPeople[cat] {
			delete(allowedPeople, cat)
			allowed++
		} else {
			disallowed++
		}
	}

	fmt.Printf("Sissepääs lubatud: %d. Sissepääs keelatud: %d.\n", allowed, disallowed)
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
