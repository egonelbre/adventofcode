package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()
	data := string(must(os.ReadFile(flag.Arg(0))))
	values := []int{}
	for _, v := range strings.Split(data, ", ") {
		values = append(values, must(strconv.Atoi(v)))
	}

	best := []int{values[0]}
	current := []int{values[0]}

	for _, v := range values[1:] {
		if current[len(current)-1] < v {
			current = append(current, v)
			if len(best) < len(current) {
				best = current
			}
		} else {
			current = []int{v}
		}
	}

	fmt.Println("connected subslice", best)
	fmt.Println("connected subslice length", len(best))

	nodes := []*node{}
	for _, v := range values {
		n := &node{value: v}
		for _, from := range nodes {
			if from.value < n.value {
				from.smaller = append(from.smaller, n)
			}
		}
		nodes = append(nodes, n)
	}

	for i := len(nodes) - 1; i >= 0; i-- {
		node := nodes[i]
		for _, smaller := range node.smaller {
			node.dist = max(node.dist, 1+smaller.dist)
		}
	}

	max := nodes[0]
	for _, n := range nodes {
		if max.dist < n.dist {
			max = n
		}
	}
	fmt.Println("distance", max.dist+1)
	printpath(max)
	fmt.Println()
}

func printpath(n *node) {
	fmt.Print(n.value, ", ")
	for _, child := range n.smaller {
		if n.dist == 1+child.dist {
			printpath(child)
			return
		}
	}
}

type node struct {
	value   int
	dist    int
	smaller []*node
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
