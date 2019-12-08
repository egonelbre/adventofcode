package intcode

import "fmt"

type Computer struct {
	Halted bool

	InstructionPointer Address

	Input  func() int64
	Output func(v int64)

	Code Code
}

type Address = int64

type Code []int64

func (code Code) Clone() Code {
	return append(Code{}, code...)
}

func (code Code) Adjust(noun, verb int64) Code {
	clone := code.Clone()
	clone[1] = noun
	clone[2] = verb
	return clone
}

type Instr interface {
	Exec(cpu *Computer) error
	Decode(code Code) (Instr, int64, error)
}

type Param struct {
	Immediate bool
	Value     Address
}

func (cpu *Computer) ValueOf(p Param) (int64, error) {
	if p.Immediate {
		return p.Value, nil
	}
	if !cpu.ValidParam(p) {
		return 0, fmt.Errorf("invalid param %v", p)
	}
	return cpu.Code[p.Value], nil
}

func (cpu *Computer) Store(at Param, value int64) error {
	if at.Immediate {
		return fmt.Errorf("cannot store at immediate %+v", at)
	}
	if !cpu.ValidParam(at) {
		return fmt.Errorf("invalid store address %+v", at)
	}
	cpu.Code[at.Value] = value
	return nil
}

func (cpu *Computer) JumpTo(target Param) error {
	addr, err := cpu.ValueOf(target)
	if err != nil {
		return err
	}
	if !cpu.ValidAddress(addr) {
		return fmt.Errorf("invalid target address %v", addr)
	}

	cpu.InstructionPointer = addr
	return nil
}

func (cpu *Computer) ValidParams(params ...Param) bool {
	for _, param := range params {
		if !cpu.ValidParam(param) {
			return false
		}
	}
	return true
}

func (cpu *Computer) ValidParam(param Param) bool {
	if param.Immediate {
		return true
	}
	return cpu.ValidAddress(param.Value)
}

func (cpu *Computer) ValidAddress(addr Address) bool {
	return 0 <= addr && addr < int64(len(cpu.Code))
}

func (cpu *Computer) Step() error {
	if cpu.InstructionPointer >= int64(len(cpu.Code)) {
		cpu.Halted = true
		return fmt.Errorf("program counter overrun")
	}

	instr, advance, err := Decode(cpu.Code[cpu.InstructionPointer:])
	if err != nil {
		cpu.Halted = true
		return err
	}
	cpu.InstructionPointer += advance

	if err := instr.Exec(cpu); err != nil {
		cpu.Halted = true
		return err
	}

	return nil
}

func (cpu *Computer) ReadInput() (int64, error) {
	if cpu.Input == nil {
		return 0, fmt.Errorf("input module missing")
	}

	return cpu.Input(), nil
}

func (cpu *Computer) WriteOutput(v int64) error {
	if cpu.Output == nil {
		return fmt.Errorf("output module missing")
	}

	cpu.Output(v)
	return nil
}

func (cpu *Computer) Run() error {
	for !cpu.Halted {
		err := cpu.Step()
		if err != nil {
			return err
		}
	}
	return nil
}
