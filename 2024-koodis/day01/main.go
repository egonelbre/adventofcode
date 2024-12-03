package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()
	data := string(must(os.ReadFile(flag.Arg(0))))

	type Price struct {
		Name string
		Cost float64
	}

	pricing := map[string]float64{}
	target := ""
	for _, line := range strings.Split(data, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		name, amount, ok := strings.Cut(line, ": ")
		if !ok {
			target = line
			break
		}

		pricing[name] = must(strconv.ParseFloat(amount, 64))
	}

	s := ""
	for product := range pricing {
		if s != "" {
			s += "|"
		}
		s += regexp.QuoteMeta(product)
	}
	s = "(" + s + ")"

	rx := regexp.MustCompile(s)
	total := 0.0
	for _, match := range rx.FindAllString(target, -1) {
		total += pricing[match]
	}

	fmt.Println(total)
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
