package intcode

import "fmt"

type Computer struct {
	Halted bool

	InstructionPointer Address

	Code Code
}

type Code []int64
type Address = int64

func (code Code) Clone() Code {
	return append(Code{}, code...)
}

func (code Code) Adjust(noun, verb int64) Code {
	clone := code.Clone()
	clone[1] = noun
	clone[2] = verb
	return clone
}

type OpCode int64

const (
	OpAdd      = OpCode(1)
	OpMultiply = OpCode(2)
	OpInput    = OpCode(3)
	OpOutput   = OpCode(4)
	OpHalt     = OpCode(99)
)

type (
	Instr interface {
		Exec(cpu *Computer) error
	}

	// [Store] := [A] + [B]
	Add struct {
		A, B  Address
		Store Address
	}

	// [Store] := [A] * [B]
	Multiply struct {
		A, B  Address
		Store Address
	}

	// [Store] := <-input
	Input struct {
		Store Address
	}

	// output <- [Load]
	Output struct {
		Load Address
	}

	Halt struct{}
)

func (op Add) Exec(cpu *Computer) error {
	if !cpu.ValidAddresses(op.A, op.B, op.Store) {
		return fmt.Errorf("invalid address %+v", op)
	}

	cpu.Code[op.Store] = cpu.Code[op.A] + cpu.Code[op.B]

	return nil
}

func (op Multiply) Exec(cpu *Computer) error {
	if !cpu.ValidAddresses(op.A, op.B, op.Store) {
		return fmt.Errorf("invalid address %+v", op)
	}

	cpu.Code[op.Store] = cpu.Code[op.A] * cpu.Code[op.B]

	return nil
}

func (op Input) Exec(cpu *Computer) error {
	if !cpu.ValidAddresses(op.Store) {
		return fmt.Errorf("invalid address %+v", op)
	}

	cpu.Code[op.Store] = cpu.Input()

	return nil
}

func (op Output) Exec(cpu *Computer) error {
	if !cpu.ValidAddresses(op.Load) {
		return fmt.Errorf("invalid address %+v", op)
	}

	cpu.Output(cpu.Code[op.Load])

	return nil
}

func (op Halt) Exec(cpu *Computer) error {
	cpu.Halted = true
	return nil
}

func (cpu *Computer) ValidAddresses(addrs ...Address) bool {
	for _, addr := range addrs {
		if !cpu.ValidAddress(addr) {
			return false
		}
	}
	return true
}

func (cpu *Computer) ValidAddress(addr Address) bool {
	return 0 <= addr && addr < int64(len(cpu.Code))
}

func (cpu *Computer) Step() error {
	if cpu.InstructionPointer >= int64(len(cpu.Code)) {
		cpu.Halted = true
		return fmt.Errorf("program counter overrun")
	}

	instr, advance, err := DecodeInstr(cpu.Code[cpu.InstructionPointer:])
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

func (cpu *Computer) Input() int64 {
	// TODO:
	return 0
}
func (cpu *Computer) Output(v int64) {
	// TODO:
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

func DecodeInstr(code Code) (instr Instr, advance int64, err error) {
	if len(code) == 0 {
		return Halt{}, 0, fmt.Errorf("code missing")
	}

	switch OpCode(code[0]) {
	case OpAdd:
		if len(code) < 4 {
			return Halt{}, 0, fmt.Errorf("add requires 3 arguments")
		}
		return Add{A: code[1], B: code[2], Store: code[3]}, 4, nil

	case OpMultiply:
		if len(code) < 4 {
			return Halt{}, 0, fmt.Errorf("multiply requires 3 arguments")
		}
		return Multiply{A: code[1], B: code[2], Store: code[3]}, 4, nil

	case OpHalt:
		return Halt{}, 1, nil

	default:
		return Halt{}, 0, fmt.Errorf("unknown opcode %d", code[0])
	}
}
