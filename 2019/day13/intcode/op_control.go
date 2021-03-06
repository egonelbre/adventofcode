package intcode

import "fmt"

// cpu.Halted := true
type Halt struct{}

const OpHalt = OpCode(99)

func init() { RegisterOp(Halt{}, OpCode(99)) }

func (op Halt) Exec(cpu *Computer) error {
	cpu.Halted = true
	return nil
}

func (Halt) Decode(code Code) (instr Instr, advance int64, err error) {
	var halt Halt
	_, advance, err = Decode0(code)
	return halt, advance, err
}

// cpu.RelativeBase := [Param]

type SetRelativeBase struct {
	Value Param
}

func init() { RegisterOp(SetRelativeBase{}, OpCode(9)) }

func (op SetRelativeBase) Exec(cpu *Computer) error {
	value, err := cpu.ValueOf(op.Value)
	if err != nil {
		return fmt.Errorf("invalid arguments %+v: %v", op, err)
	}

	cpu.RelativeBase += value
	return nil
}

func (SetRelativeBase) Decode(code Code) (instr Instr, advance int64, err error) {
	var op SetRelativeBase
	_, op.Value, advance, err = Decode1(code)
	return op, advance, err
}

// if [Check] != 0 then cpu.InstructionPointer := [Store]
type JumpIfTrue struct {
	Check  Param
	Target Param
}

func init() { RegisterOp(JumpIfTrue{}, OpCode(5)) }

func (JumpIfTrue) Decode(code Code) (instr Instr, advance int64, err error) {
	var op JumpIfTrue
	_, op.Check, op.Target, advance, err = Decode2(code)
	return op, advance, err
}

func (op JumpIfTrue) Exec(cpu *Computer) error {
	check, err := cpu.ValueOf(op.Check)
	if err != nil {
		return fmt.Errorf("invalid arguments %+v: %v", op, err)
	}
	if check != 0 {
		return cpu.JumpTo(op.Target)
	}
	return nil
}

// if [Check] == 0 then cpu.InstructionPointer := [Store]
type JumpIfFalse struct {
	Check  Param
	Target Param
}

func init() { RegisterOp(JumpIfFalse{}, OpCode(6)) }

func (op JumpIfFalse) Exec(cpu *Computer) error {
	check, err := cpu.ValueOf(op.Check)
	if err != nil {
		return fmt.Errorf("invalid arguments %+v: %v", op, err)
	}
	if check == 0 {
		return cpu.JumpTo(op.Target)
	}
	return nil
}

func (JumpIfFalse) Decode(code Code) (instr Instr, advance int64, err error) {
	var op JumpIfFalse
	_, op.Check, op.Target, advance, err = Decode2(code)
	return op, advance, err
}
