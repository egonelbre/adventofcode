package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const AlphabetData = `ABCDEFGHIJKLMNOPRSŠZŽTUVÕÄÖÜ`

func main() {
	flag.Parse()
	data := string(must(os.ReadFile(flag.Arg(0))))

	//hackfix
	data = strings.ReplaceAll(data, ", kes kauaks ei jää", "")
	data = strings.ReplaceAll(data, ", et piparkooke küpsetada", "")

	syndmused, naabrid, _ := strings.Cut(data, "Naabrid:")
	syndmused = strings.TrimPrefix(syndmused, "Sündmused:")

	syndmused = strings.TrimSpace(syndmused)
	naabrid = strings.TrimSpace(naabrid)

	eventcost := map[string]int{}
	for _, event := range strings.Split(syndmused, "\n") {
		event = strings.TrimSpace(event)
		if event == "" {
			continue
		}
		name, cost, ok := strings.Cut(event, ": ")
		if !ok {
			panic(event)
		}

		eventcost[name] = must(strconv.Atoi(cost))
	}

	kahtlased := 0

	for _, naaber := range strings.Split(naabrid, "\n") {
		naaber = strings.TrimSpace(naaber)
		if naaber == "" {
			continue
		}

		id, events, ok := strings.Cut(naaber, ": ")
		if !ok {
			panic(naaber)
		}

		total := 0
		for _, event := range strings.Split(events, ", ") {
			cost, ok := eventcost[event]
			if !ok {
				panic(fmt.Sprintf("%q\n%q\n", naaber, event))
			}
			total += cost
		}

		fmt.Println(id, total)
		if total > 10 {
			kahtlased++
		}
	}

	fmt.Println(kahtlased)
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
