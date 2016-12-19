package main

import (
	"flag"
	"fmt"
)

var (
	verbose = flag.Bool("verbose", false, "verbose output")
)

type Elf struct {
	Name     int
	Presents int
	Next     *Elf
}

func NewCircle(n int) (count int, first, precenter *Elf) {
	var prev *Elf
	first = &Elf{Name: 1, Presents: 1}
	prev = first

	centern := 1 + n/2
	for i := 2; i <= n; i++ {
		next := &Elf{Name: i, Presents: 1}
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
		if *verbose {
			fmt.Println(
				active.Name,
				"<< ", active.Next.Name, "[", active.Next.Presents, "]",
				"=", active.Presents+active.Next.Presents)
		}
		active.Presents += active.Next.Presents
		active.Next = active.Next.Next
		active = active.Next
	}
	return active
}

func StealAcross(count int, active, precenter *Elf) *Elf {
	for active != active.Next {
		center := precenter.Next
		if *verbose {
			fmt.Println(
				active.Name,
				"<< ", center.Name, "[", center.Presents, "]",
				"=", active.Presents+center.Presents)
		}
		active.Presents += center.Presents

		precenter.Next = center.Next
		if count&1 == 1 {
			precenter = precenter.Next
		}

		active = active.Next
		count--
	}
	return active
}

func main() {
	flag.Parse()

	// fmt.Println(StealLeft(NewCircle(5)))
	// fmt.Println(StealLeft(NewCircle(3001330)))
	fmt.Println(StealAcross(NewCircle(5)))
	fmt.Println(StealAcross(NewCircle(6)))
	fmt.Println(StealAcross(NewCircle(3001330)))
}
