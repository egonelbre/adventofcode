package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()
	data := string(must(os.ReadFile(flag.Arg(0))))

	lines := strings.Split(data, "\n")
	lines[0], lines[len(lines)-1] = lines[len(lines)-1], lines[0]

	count := len(lines)
	allvisited := uint64((1<<count - 1) &^ 1)

	cityid := map[string]int{}
	idcity := map[int]string{}
	cities := make([][]int64, len(lines))
	for i := range cities {
		cities[i] = make([]int64, len(lines))
	}

	ensure := func(v string) int {
		if n, ok := cityid[v]; ok {
			return n
		}
		id := len(cityid)
		cityid[v] = id
		idcity[id] = v
		return id
	}

	for _, line := range lines {
		line = strings.TrimSpace(line)

		from, destinations, ok := strings.Cut(line, ": ")
		if !ok {
			panic("not ok")
		}

		fromnode := ensure(from)

		for _, to := range strings.Split(destinations, "| ; |") {
			city, time, ok := strings.Cut(to, " - ")
			if !ok {
				panic("not ok")
			}
			tonode := ensure(city)

			time = strings.TrimSuffix(time, "min")
			minutes := must(strconv.ParseInt(time, 10, 64))

			cities[fromnode][tonode] = minutes
		}
	}

	for _, dist := range cities {
		fmt.Println(dist)
	}
	fmt.Println()

	besttime := int64(math.MaxInt64)
	bestpath := []int{}

	var traverse func(path []int, at int, visited uint64, time int64)
	traverse = func(path []int, at int, visited uint64, time int64) {
		if allvisited == visited {
			time += cities[at][0]
			if time < besttime {
				besttime = time
				bestpath = append(slices.Clone(path), 0)
			}
			return
		}
		for k := 1; k < count; k++ {
			if visited&(1<<k) != 0 {
				continue
			}
			traverse(append(path, k), k, visited|(1<<k), time+cities[at][k])
		}
	}

	traverse([]int{0}, 0, 0, 0)

	fmt.Println(besttime)
	fmt.Println(bestpath)
	xs := []string{}
	for _, id := range bestpath {
		xs = append(xs, idcity[id])
	}
	fmt.Println(strings.Join(xs, " -> "))
}

type node struct {
	id    int64
	label string
}

func (n *node) ID() int64 { return n.id }

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
