package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var rxChemical = regexp.MustCompile(`\b(\d+)\s+([A-Z]+)\b`)

func Parse(input string) (*Reactions, error) {
	reactions := NewReactions()

	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		chemstrs := rxChemical.FindAllStringSubmatch(line, -1)
		if len(chemstrs) < 2 {
			return nil, fmt.Errorf("invalid line %q", line)
		}

		var chemicals []Chemical
		for _, chemstr := range chemstrs {
			count, err := strconv.Atoi(chemstr[1])
			if err != nil {
				return nil, fmt.Errorf("invalid line %q, chem %q", line, chemstr)
			}

			chemicals = append(chemicals, Chemical{
				Name:  chemstr[2],
				Count: int64(count),
			})
		}

		last := len(chemicals) - 1
		reactions.Add(Reaction{
			Input:  chemicals[:last],
			Output: chemicals[last],
		})
	}

	reactions.Sort()

	return reactions, nil
}
