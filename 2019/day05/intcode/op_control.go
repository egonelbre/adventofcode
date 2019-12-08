package intcode

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

// if [Check] == 0 then cpu.InstructionPointer := [Store]
type JumpIfFalse struct {
	Check  Param
	Target Param
}

func init() { RegisterOp(JumpIfFalse{}, OpCode(6)) }

func (JumpIfFalse) Decode(code Code) (instr Instr, advance int64, err error) {
	var op JumpIfFalse
	_, op.Check, op.Target, advance, err = Decode2(code)
	return op, advance, err
}
