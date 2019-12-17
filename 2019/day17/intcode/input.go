package intcode

func WriteValue(cpu *Computer, value int64) (ok bool, err error) {
	cpu.Halted = false

	cpu.Input = func() int64 {
		cpu.Halted = true
		ok = true
		return value
	}
	cpu.Output = nil

	err = cpu.Run()
	return ok, err
}

func ReadValue(cpu *Computer) (output int64, ok bool, err error) {
	cpu.Halted = false

	cpu.Input = nil
	cpu.Output = func(v int64) {
		cpu.Halted = true
		output = v
		ok = true
	}

	err = cpu.Run()
	return output, ok, err
}
