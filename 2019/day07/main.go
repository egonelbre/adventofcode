package main

import (
	"fmt"
	"os"

	"github.com/egonelbre/adventofcode/2019/day07/intcode"
)

func main() {
	cache := Cache{}

	var best [5]int
	var maxOutput int64

	Permutations(5, []int{0, 1, 2, 3, 4}, func(xs []int) {
		var signal int64
		for _, phase := range xs {
			var err error
			signal, err = cache.Amplify(int64(phase), signal)
			if err != nil {
				fmt.Fprintf(os.Stderr, "failed amplification %v: %v\n", xs, err)
				return
			}
		}
		if signal > maxOutput {
			maxOutput = signal
			copy(best[:], xs)
		}
	})

	fmt.Println("phases", best, "max output", maxOutput)
}

// Permutations implements Heap's algorithm.
func Permutations(n int, xs []int, fn func([]int)) {
	if n == 1 {
		fn(xs)
		return
	}

	Permutations(n-1, xs, fn)

	for i := 0; i < n-1; i++ {
		if i&1 == 0 {
			xs[i], xs[n-1] = xs[n-1], xs[i]
		} else {
			xs[0], xs[n-1] = xs[n-1], xs[0]
		}

		Permutations(n-1, xs, fn)
	}
}

type Input struct {
	Phase  int64
	Signal int64
}

type Cache map[Input]int64

func (cache Cache) Amplify(phase, signal int64) (int64, error) {
	cached, ok := cache[Input{phase, signal}]
	if ok {
		return cached, nil
	}

	output, err := Amplify(phase, signal)
	if err != nil {
		return output, err
	}

	cache[Input{phase, signal}] = output

	return output, nil
}

func Amplify(phase, signal int64) (int64, error) {
	var inputIndex int64
	var output int64

	cpu := &intcode.Computer{
		Input: func() int64 {
			inputIndex++
			switch inputIndex {
			case 1:
				return phase
			case 2:
				return signal
			default:
				fmt.Fprintf(os.Stderr, "invalid requested input %v\n", inputIndex)
				return 0
			}
		},
		Output: func(v int64) {
			output = v
		},
		Code: AmplifierControllerSoftware.Clone(),
	}

	err := cpu.Run()
	return output, err
}
