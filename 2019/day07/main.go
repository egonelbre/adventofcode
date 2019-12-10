package main

import (
	"fmt"
	"os"

	"github.com/egonelbre/adventofcode/2019/day07/intcode"
)

func main() {
	best := Feedforward()
	Feedback(best)
}

// part 1
func Feedforward() []int {
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
	return best[:]
}

// part 2
func Feedback(phases []int) {
	var maxOutput int64

	cpus := make([]*intcode.Computer, len(phases))
	for i, phase := range phases {
		cpus[i] = &intcode.Computer{
			Input:  nil,
			Output: nil,
			Code:   AmplifierControllerSoftware.Clone(),
		}

		err := SetPhase(cpus[i], int64(phase))
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed setting phase %v: %v\n", phases, err)
			return
		}
	}

	var signal int64
feedback:
	for {
		for i, cpu := range cpus {
			output, ok, err := Feed(cpu, signal)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cpu failed setting input %d: %v\n", i, err)
				break feedback
			}

			signal = output
			if !ok {
				fmt.Fprintf(os.Stderr, "cpu %d finished: %v\n", i, err)
				break feedback
			}
		}

		if signal > maxOutput {
			maxOutput = signal
		}

		fmt.Println(signal)
	}

	fmt.Println("max feedback output", maxOutput)
}

func SetPhase(cpu *intcode.Computer, phase int64) error {
	cpu.Halted = false

	cpu.Input = func() int64 {
		cpu.Halted = true
		return phase
	}
	cpu.Output = nil

	return cpu.Run()
}

func Feed(cpu *intcode.Computer, input int64) (output int64, ok bool, err error) {
	cpu.Halted = false

	cpu.Input = func() int64 {
		return input
	}
	cpu.Output = func(v int64) {
		cpu.Halted = true
		output = v
		ok = true
	}

	err = cpu.Run()
	return output, ok, err
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
