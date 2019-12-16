package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}

		defer pprof.StopCPUProfile()
	}

	pattern := []int8{0, 1, 0, -1}

	fmt.Println("====")
	Run(NumberFromString("12345678"), pattern, 4, 8)

	fmt.Println("===")
	input := NumberFromString(Input)
	Run(input, pattern, 100, 8)

	fmt.Println("====")
	offset := input.Slice(0, 7).Int()
	input10000 := input.Repeat(10000)
	result := Run(input10000, pattern, 100, 8)
	fmt.Println(result.Slice(offset, offset+8))
}

func Run(num Number, pattern []int8, steps int, limit int) Number {
	fmt.Println("S", 0, num.StringN(limit))
	for i := 0; i < steps; i++ {
		num = num.FFT(pattern)
		fmt.Println("S", i+1, num.StringN(limit))
	}
	return num
}
