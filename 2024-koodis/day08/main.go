package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()
	data := string(must(os.ReadFile(flag.Arg(0))))

	lines := strings.Split(data, "\n")
	capacity := must(strconv.Atoi(strings.TrimPrefix(lines[0], "Maksimaalne kandevõime: ")))
	products := lines[2:]

	weights := []int{}
	for _, prod := range products {
		weights = append(weights, must(strconv.Atoi(prod)))
	}

	slices.Sort(weights)

	cars := []*Car{}
nextWeight:
	for weighti := len(weights) - 1; weighti >= 0; weighti-- {
		weight := weights[weighti]

		for _, car := range cars {
			if car.Total+weight <= capacity {
				car.Add(weight)
				weights = slices.Delete(weights, weighti, weighti+1)
				continue nextWeight
			}
		}

		cars = append(cars, &Car{
			Total:   weight,
			Ballast: []int{weight},
		})
		weights = slices.Delete(weights, weighti, weighti+1)
	}

	fmt.Println(capacity)
	fmt.Println(weights)
	fmt.Println(cars)

	fmt.Println(len(cars))
}

type Car struct {
	Total   int
	Ballast []int
}

func (c *Car) Add(weight int) {
	c.Total += weight
	c.Ballast = append(c.Ballast, weight)
}

func (c *Car) String() string {
	return fmt.Sprintf("%v %v\n", c.Total, c.Ballast)
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
