package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	verbose = flag.Bool("verbose", false, "verbose output")
)

type Elf struct {
	Name int
	Next *Elf
}

func NewCircle(n int) (count int, first, precenter *Elf) {
	elves := make([]Elf, n)

	var prev *Elf
	first = &elves[0]
	first.Name = 1
	prev = first

	centern := 1 + n/2
	for i := 2; i <= n; i++ {
		next := &elves[i-1]
		next.Name = i
		if i == centern-1 {
			precenter = next
		}
		prev.Next = next
		prev = next
	}
	prev.Next = first
	return n, first, precenter
}

func StealLeft(count int, active, precenter *Elf) *Elf {
	for active != active.Next {
		active.Next = active.Next.Next
		active = active.Next
	}
	return active
}

func StealAcross(count int, active, precenter *Elf) *Elf {
	for active != active.Next {
		center := precenter.Next
		precenter.Next = center.Next
		if count&1 == 1 {
			precenter = precenter.Next
		}

		active = active.Next
		count--
	}
	return active
}

func StealAcrossAlternate(count int) int {
	elves := make([]int, count)
	next := make([]int, count)
	for i := range elves {
		elves[i] = i + 1
	}

	for count > 1 {
		next = next[:0:cap(next)]
		j := 0
		k := -1
		d := 0
		for i, elf := range elves {
			if elf != 0 {
				j = i + (count-d)/2 + d
				if j < count {
					k = i
					j = j % count
					elves[j] = 0
					d++
				} else {
					next = append(next, elf)
				}
			}
		}

		next = append(next, elves[:k+1]...) // O(n)
		elves, next = next, elves
		count = len(elves)
	}

	return elves[0]
}

func main() {
	flag.Parse()

	// fmt.Println(StealLeft(NewCircle(5)))
	// fmt.Println(StealLeft(NewCircle(3001330)))
	fmt.Println(StealAcross(NewCircle(5)))
	fmt.Println(StealAcross(NewCircle(6)))

	N := 64 * 3001330

	start := time.Now()
	fmt.Println(StealAcross(NewCircle(N)))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println(StealAcrossAvo(N))
	fmt.Println(time.Since(start))
}
