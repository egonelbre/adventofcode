package main

import (
	"fmt"
	"strconv"
)

type Seq []int

func (s Seq) String() string {
	const i10 = "0123456789"
	r := ""
	for _, v := range s {
		if v < 10 {
			r += string(i10[v])
		} else {
			r += strconv.Itoa(v)
		}
	}
	return r
}

func process(in Seq) Seq {
	out := make(Seq, 0, len(in)*2)
	count, cursor := 0, 0
	for _, v := range in {
		if v != cursor {
			if count > 0 {
				out = append(out, count, cursor)
			}
			count, cursor = 0, v
		}
		count++
	}
	if count > 0 {
		out = append(out, count, cursor)
	}
	return out
}

func main() {
	fmt.Println(process(Seq{1}))
	fmt.Println(process(Seq{1, 1}))
	fmt.Println(process(Seq{2, 1}))
	fmt.Println(process(Seq{1, 2, 1, 1}))
	fmt.Println(process(Seq{1, 1, 1, 2, 2, 1}))

	input := Seq{1, 1, 1, 3, 2, 2, 2, 1, 1, 3}
	for i := 0; i < 50; i++ {
		input = process(input)
	}
	fmt.Println(len(input))
}
