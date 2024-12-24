package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()
	data := string(must(os.ReadFile(flag.Arg(0))))

	people := []Person{}

	for _, line := range strings.Split(data, "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		name, times, _ := strings.Cut(line, ": ")
		person := Person{}
		person.Name = name
		for _, job := range strings.Split(times, ";") {
			jobname, time, _ := strings.Cut(job, "-")
			minutes := must(strconv.Atoi(strings.TrimSuffix(time, " min")))

			person.Jobs = append(person.Jobs, Job{
				Name:    jobname,
				Minutes: minutes,
			})
		}

		people = append(people, person)
	}

	jobcount := len(people[0].Jobs)
	best := math.MaxInt

	var iterate func(who []int, k int)
	iterate = func(who []int, k int) {
		if k == jobcount {
			t := costtime(people, who)
			if t < best {
				fmt.Println(who)
				best = t
			}
			return
		}

		for i := range people {
			iterate(append(who, i), k+1)
		}
	}
	fmt.Println(best)

	iterate(make([]int, 0, 10), 0)

	fmt.Println(people)
	fmt.Println(best)
}

type Job struct {
	Name    string
	Minutes int
}

type Person struct {
	Name string
	Jobs []Job
}

func costtime(people []Person, worker []int) int {
	spent := make([]int, len(people))
	for job, who := range worker {
		spent[who] += people[who].Jobs[job].Minutes
	}

	t := 0
	for _, v := range spent {
		t = max(t, v)
	}
	return t
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
