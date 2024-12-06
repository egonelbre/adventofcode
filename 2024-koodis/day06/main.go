package main

import (
	"flag"
	"fmt"
	"maps"
	"os"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()
	data := string(must(os.ReadFile(flag.Arg(0))))

	lines := strings.Split(data, "\n")
	prefix := lines[0]
	lines = lines[1:]

	numbers := map[string]struct{}{}

	for _, line := range lines {
		if strings.HasPrefix(line, prefix) {
			numbers[line] = struct{}{}
		}
	}

	const numlen = 7
	rx := regexp.MustCompile(prefix + "[0-9]{" + strconv.Itoa(7-len(prefix)) + "}")

	for _, line := range rotate(lines) {
		for _, num := range rx.FindAllString(line, -1) {
			numbers[num] = struct{}{}
		}
	}

	result := slices.Collect(maps.Keys(numbers))
	sort.Strings(result)
	fmt.Println(strings.Join(result, ", "))
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func rotate(vs []string) []string {
	xs := make([][]byte, len(vs[0]))
	for _, line := range vs {
		for i, b := range []byte(line) {
			xs[i] = append(xs[i], b)
		}
	}

	rs := []string{}
	for _, x := range xs {
		rs = append(rs, string(x))
	}
	return rs
}
