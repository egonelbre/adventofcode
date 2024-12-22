package main

import (
	"flag"
	"fmt"
	"maps"
	"os"
	"slices"
	"strconv"
	"strings"
)

type trap struct {
	wood  int32
	nails int32
	rope  int32
}

func (t trap) canbuild(b trap) bool {
	return t.wood >= b.wood && t.nails >= b.nails && t.rope >= b.rope
}

func (t trap) sub(b trap) trap {
	t.wood -= b.wood
	t.nails -= b.nails
	t.rope -= b.rope
	return t
}

func main() {
	flag.Parse()
	data := string(must(os.ReadFile(flag.Arg(0))))

	trapstext, resourcesdata, _ := strings.Cut(data, "Ressursid:\n")

	traps := []trap{}
	for _, line := range strings.Split(trapstext, "\n")[1:] {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		_, res, _ := strings.Cut(line, ": ")
		var trap trap
		resx := strings.Split(res, ", ")
		trap.wood = int32(must(strconv.Atoi(strings.TrimPrefix(resx[0], "puit:"))))
		trap.nails = int32(must(strconv.Atoi(strings.TrimPrefix(resx[1], "naelad:"))))
		trap.rope = int32(must(strconv.Atoi(strings.TrimPrefix(resx[2], "köis:"))))

		traps = append(traps, trap)
	}

	resources := strings.Split(resourcesdata, "\n")
	var available trap
	available.wood = int32(must(strconv.Atoi(strings.TrimPrefix(resources[0], "puit: "))))
	available.nails = int32(must(strconv.Atoi(strings.TrimPrefix(resources[1], "naelad: "))))
	available.rope = int32(must(strconv.Atoi(strings.TrimPrefix(resources[2], "köis: "))))

	fmt.Println(traps)
	fmt.Println(available)

	best := map[uint32]struct{}{}
	bestcost := uint32(0)

	var fill func(at int, built, count uint32, free trap)
	fill = func(at int, built, count uint32, free trap) {
		if at >= len(traps) {
			if bestcost < count {
				bestcost = count
				best = map[uint32]struct{}{}
				best[built] = struct{}{}
			} else if bestcost == count {
				best[built] = struct{}{}
			}
			return
		}

		fill(at+1, built, count, free)
		if free.canbuild(traps[at]) {
			fill(at+1, built|(1<<at), count+1, free.sub(traps[at]))
		}
	}

	fill(0, 0, 0, available)

	fmt.Println("TRAPS:", bestcost)

	values := slices.Collect(maps.Keys(best))
	slices.Sort(values)
	for _, built := range values {
		fmt.Printf("%-30v [%32b]\n", builttostr(built), built)
	}
}

func builttostr(v uint32) string {
	s := ""
	for k := 0; k < 32; k++ {
		if v&(1<<k) != 0 {
			if s != "" {
				s += ", "
			}
			s += strconv.Itoa(k)
		}
	}
	return s
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
