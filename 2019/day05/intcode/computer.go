package intcode

import "fmt"

type Computer struct {
	Halted bool

	InstructionPointer Address

	Input  func() int64
	Output func(v int64)

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

type Param struct {
	Immediate bool
	Value     Address
}

type (
	Instr interface {
		Exec(cpu *Computer) error
	}

	// [Store] := [A] + [B]
	Add struct {
		A, B  Param
		Store Param
	}

	// [Store] := [A] * [B]
	Multiply struct {
		A, B  Param
		Store Param
	}

	// [Store] := <-input
	Input struct {
		Store Param
	}

	// output <- [Load]
	Output struct {
		Load Param
	}

	Halt struct{}
)

func (op Add) Exec(cpu *Computer) error {
	a, aerr := cpu.ValueOf(op.A)
	b, berr := cpu.ValueOf(op.B)
	if aerr != nil || berr != nil {
		return fmt.Errorf("invalid arguments %+v: %v, %v", op, aerr, berr)
	}

	return cpu.Store(op.Store, a+b)
}

func (op Multiply) Exec(cpu *Computer) error {
	a, aerr := cpu.ValueOf(op.A)
	b, berr := cpu.ValueOf(op.B)
	if aerr != nil || berr != nil {
		return fmt.Errorf("invalid arguments %+v: %v, %v", op, aerr, berr)
	}

	return cpu.Store(op.Store, a*b)
}

func (op Input) Exec(cpu *Computer) error {
	v, err := cpu.ReadInput()
	if err != nil {
		return err
	}
	return cpu.Store(op.Store, v)
}

func (op Output) Exec(cpu *Computer) error {
	a, aerr := cpu.ValueOf(op.Load)
	if aerr != nil {
		return fmt.Errorf("invalid arguments %+v: %v", op, aerr)
	}

	return cpu.WriteOutput(a)
}

func (op Halt) Exec(cpu *Computer) error {
	cpu.Halted = true
	return nil
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
	return 0 <= param.Value && param.Value < int64(len(cpu.Code))
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

func DecodeInstr(code Code) (instr Instr, advance int64, err error) {
	if len(code) == 0 {
		return Halt{}, 0, fmt.Errorf("code missing")
	}

	// ABCDE
	//  1002
	//
	// DE - two digit opcode
	//  C - mode of 1st param
	//  B - mode of 2nd param
	//  A - mode of 3rd param

	full := code[0]
	opcode := OpCode(full % 100)

	imm1 := 1 == (full/100)%10
	imm2 := 1 == (full/1000)%10
	imm3 := 1 == (full/10000)%10

	switch opcode {
	case OpAdd:
		if len(code) < 4 {
			return Halt{}, 0, fmt.Errorf("add requires 3 arguments")
		}
		return Add{A: Param{imm1, code[1]}, B: Param{imm2, code[2]}, Store: Param{imm3, code[3]}}, 4, nil

	case OpMultiply:
		if len(code) < 4 {
			return Halt{}, 0, fmt.Errorf("multiply requires 3 arguments")
		}
		return Multiply{A: Param{imm1, code[1]}, B: Param{imm2, code[2]}, Store: Param{imm3, code[3]}}, 4, nil

	case OpInput:
		if len(code) < 2 {
			return Halt{}, 0, fmt.Errorf("input requires 1 arguments")
		}
		return Input{Store: Param{imm1, code[1]}}, 2, nil

	case OpOutput:
		if len(code) < 2 {
			return Halt{}, 0, fmt.Errorf("output requires 1 arguments")
		}

		return Output{Load: Param{imm1, code[1]}}, 2, nil

	case OpHalt:
		return Halt{}, 1, nil

	default:
		return Halt{}, 0, fmt.Errorf("unknown opcode %d", code[0])
	}
}
