package intcode

import "fmt"

// [Store] := <-input
type Input struct {
	Store Param
}

func init() { RegisterOp(Input{}, OpCode(3)) }

func (op Input) Exec(cpu *Computer) error {
	v, err := cpu.ReadInput()
	if err != nil {
		return err
	}
	return cpu.Store(op.Store, v)
}

func (Input) Decode(code Code) (instr Instr, advance int64, err error) {
	var op Input
	_, op.Store, advance, err = Decode1(code)
	return op, advance, err
}

// output <- [Load]
type Output struct {
	Load Param
}

func init() { RegisterOp(Output{}, OpCode(4)) }

func (op Output) Exec(cpu *Computer) error {
	a, aerr := cpu.ValueOf(op.Load)
	if aerr != nil {
		return fmt.Errorf("invalid arguments %+v: %v", op, aerr)
	}

	return cpu.WriteOutput(a)
}

func (Output) Decode(code Code) (instr Instr, advance int64, err error) {
	var op Output
	_, op.Load, advance, err = Decode1(code)
	return op, advance, err
}
