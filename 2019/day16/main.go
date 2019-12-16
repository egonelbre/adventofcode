package main

import "fmt"

func main() {
	pattern := []int8{0, 1, 0, -1}
	for i := 0; i < 8; i++ {
		for k := 0; k < 8; k++ {
			fmt.Printf("%3d", Repeat(pattern, i, k))
		}
		fmt.Println()
	}

	fmt.Println("====")
	Run(NumberFromString("12345678"), pattern, 4, 8)

	fmt.Println("===")
	input := NumberFromString(Input)
	Run(input, pattern, 100, 8)
}

func Run(num Number, pattern []int8, steps int, limit int) {
	fmt.Println("S", 0, num.StringN(limit))
	for i := 0; i < steps; i++ {
		num = num.FFT(pattern)
		fmt.Println("S", i+1, num.StringN(limit))
	}
}
