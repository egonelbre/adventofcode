package main

import (
	"strings"
)

type Number struct {
	Digits []int8
}

func NumberFromString(s string) Number {
	n := Number{
		Digits: make([]int8, len(s)),
	}
	for i, r := range s {
		n.Digits[i] = int8(r - '0')
	}
	return n
}

func (a Number) FFT(pattern []int8) Number {
	r := Number{
		make([]int8, len(a.Digits)),
	}
	for i := range a.Digits {
		var t int64
		for k, b := range a.Digits {
			t += int64(b) * int64(Repeat(pattern, i, k))
		}
		r.Digits[i] = Downsize(t)
	}
	return r
}

func Repeat(pattern []int8, out, in int) int8 {
	p := ((in + 1) / (out + 1))
	//return int8(p % len(pattern))
	return pattern[p%len(pattern)]
}

func PickRepeated(i, n int) int {
	return i % n
}

func Downsize(t int64) int8 {
	if t < 0 {
		return int8(-t % 10)
	}
	return int8(t % 10)
}

func (a Number) String() string {
	var b strings.Builder
	for _, v := range a.Digits {
		_ = b.WriteByte(byte(v) + '0')
	}
	return b.String()
}

func (a Number) StringN(n int) string {
	var b strings.Builder
	if n > len(a.Digits) {
		n = len(a.Digits)
	}
	for _, v := range a.Digits[:n] {
		_ = b.WriteByte(byte(v) + '0')
	}
	return b.String()
}
