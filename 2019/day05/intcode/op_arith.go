package intcode

import "fmt"

// [Store] := [A] + [B]
type Add struct {
	A, B  Param
	Store Param
}

func init() { RegisterOp(Add{}, OpCode(1)) }

func (op Add) Exec(cpu *Computer) error {
	a, aerr := cpu.ValueOf(op.A)
	b, berr := cpu.ValueOf(op.B)
	if aerr != nil || berr != nil {
		return fmt.Errorf("invalid arguments %+v: %v, %v", op, aerr, berr)
	}

	return cpu.Store(op.Store, a+b)
}

func (Add) Decode(code Code) (instr Instr, advance int64, err error) {
	var op Add
	_, op.A, op.B, op.Store, advance, err = Decode3(code)
	return op, advance, err
}

// [Store] := [A] * [B]
type Multiply struct {
	A, B  Param
	Store Param
}

func init() { RegisterOp(Multiply{}, OpCode(2)) }

func (op Multiply) Exec(cpu *Computer) error {
	a, aerr := cpu.ValueOf(op.A)
	b, berr := cpu.ValueOf(op.B)
	if aerr != nil || berr != nil {
		return fmt.Errorf("invalid arguments %+v: %v, %v", op, aerr, berr)
	}

	return cpu.Store(op.Store, a*b)
}

func (Multiply) Decode(code Code) (instr Instr, advance int64, err error) {
	var op Multiply
	_, op.A, op.B, op.Store, advance, err = Decode3(code)
	return op, advance, err
}

// [Store] := [A] < [B] ? 1 : 0
type LessThan struct {
	A, B  Param
	Store Param
}

func init() { RegisterOp(LessThan{}, OpCode(7)) }

func (LessThan) Decode(code Code) (instr Instr, advance int64, err error) {
	var op LessThan
	_, op.A, op.B, op.Store, advance, err = Decode3(code)
	return op, advance, err
}

// [Store] := [A] == [B] ? 1 : 0
type Equals struct {
	A, B  Param
	Store Param
}

func init() { RegisterOp(Equals{}, OpCode(8)) }

func (Equals) Decode(code Code) (instr Instr, advance int64, err error) {
	var op Equals
	_, op.A, op.B, op.Store, advance, err = Decode3(code)
	return op, advance, err
}
