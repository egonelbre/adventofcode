package intcode

import (
	"fmt"
)

type Computer struct {
	Halted bool

	RelativeBase       Address
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

func (code *Code) ResizeTo(addr Address) {
	if addr < int64(len(*code)) {
		return
	}

	increaseBy := addr + 1 - int64(len(*code))
	*code = append(*code, make(Code, increaseBy)...)
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
	Mode  AddressingMode
	Value Address
}

type AddressingMode byte

const (
	Absolute  = AddressingMode(0)
	Immediate = AddressingMode(1)
	Relative  = AddressingMode(2)
)

func (cpu *Computer) ValueOf(p Param) (int64, error) {
	switch p.Mode {
	case Immediate:
		return p.Value, nil
	case Absolute:
		addr := p.Value
		if addr < 0 {
			return 0, fmt.Errorf("invalid address %v", p)
		}
		if addr >= int64(len(cpu.Code)) {
			return 0, nil
		}
		return cpu.Code[addr], nil
	case Relative:
		addr := cpu.RelativeBase + p.Value
		if addr < 0 {
			return 0, fmt.Errorf("invalid address %v", p)
		}
		if addr >= int64(len(cpu.Code)) {
			return 0, nil
		}
		return cpu.Code[addr], nil
	default:
		return 0, fmt.Errorf("invalid addressing mode %v", p)
	}
}

func (cpu *Computer) Store(at Param, value int64) error {
	switch at.Mode {
	case Immediate:
		return fmt.Errorf("cannot store at immediate %+v", at)
	case Absolute:
		addr := at.Value
		if addr < 0 {
			return fmt.Errorf("invalid address %v", at)
		}
		cpu.Code.ResizeTo(addr)
		cpu.Code[addr] = value
		return nil
	case Relative:
		addr := cpu.RelativeBase + at.Value
		if addr < 0 {
			return fmt.Errorf("invalid address %v", at)
		}
		cpu.Code.ResizeTo(addr)
		cpu.Code[addr] = value
		return nil
	default:
		return fmt.Errorf("invalid addressing mode %v", at)
	}
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
