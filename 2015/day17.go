// +build ignore

package main

import (
	"fmt"
	"sort"
)

func bitcount(v uint32, n int) int {
	count := 0
	for i := 0; i < n; i++ {
		if v&(1<<uint(i)) != 0 {
			count++
		}
	}
	return count
}

var min int
var countmin int

func count(index int, sum int, unused uint32, containers []int) int {
	if sum > 150 || unused == 0 {
		return 0
	}
	if sum == 150 {
		//fmt.Printf("%020b\n", unused)
		x := len(containers) - bitcount(unused, len(containers))
		for i := 0; i < len(containers); i++ {
			if unused&(1<<uint(i)) != 0 {
				fmt.Print(containers[i], " ")
			}
		}
		fmt.Println()

		if x < min {
			min = x
			countmin = 0
		}
		if x == min {
			countmin++
		}
		return 1
	}
	total := 0
	for i := index; i < len(containers); i++ {
		s := byte(index)
		if unused&(1<<s) == 0 {
			continue
		}

		total += count(i+1, sum+containers[i], unused&^(1<<s), containers)
	}
	return total
}

func Count(containers []int) int {
	return count(0, 0, 1<<uint(len(containers))-1, containers)
}

func main() {
	var containers = []int{
		50, 44, 11, 49, 42, 46, 18, 32, 26, 40, 21, 7, 18, 43, 10, 47, 36, 24, 22, 40,
	}
	sort.Ints(containers)
	fmt.Println(len(containers))
	min = len(containers)
	fmt.Println(Count(containers))
	fmt.Println(min)
	fmt.Println(countmin)
}
