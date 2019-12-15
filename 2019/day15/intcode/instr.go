package intcode

import "fmt"

type OpCode int64

var Ops = map[OpCode]Instr{}

func RegisterOp(instr Instr, code OpCode) {
	_, exists := Ops[code]
	if exists {
		panic(fmt.Sprintf("code %d already registered", code))
	}

	Ops[code] = instr
}

func Decode(code Code) (instr Instr, advance int64, err error) {
	opcode, _, err := Decode0(code)
	if err != nil {
		return Halt{}, 0, fmt.Errorf("failed to decode opcode: %w", err)
	}

	decoder, ok := Ops[opcode]
	if !ok {
		return Halt{}, 0, fmt.Errorf("unknown opcode %v", opcode)
	}

	return decoder.Decode(code)
}

func ParseInstructionCode(code int64) (op OpCode, m1, m2, m3 AddressingMode) {
	// ABCDE
	//  1002
	//
	// DE - two digit opcode
	//  C - mode of 1st param
	//  B - mode of 2nd param
	//  A - mode of 3rd param

	return OpCode(code % 100),
		AddressingMode((code / 100) % 10),
		AddressingMode((code / 1000) % 10),
		AddressingMode((code / 10000) % 10)
}

func Decode0(code Code) (op OpCode, advance int64, err error) {
	if len(code) < 1 {
		return OpHalt, 0, fmt.Errorf("requires opcode")
	}

	opcode, _, _, _ := ParseInstructionCode(code[0])
	return opcode, 1, nil
}

func Decode1(code Code) (op OpCode, a1 Param, advance int64, err error) {
	if len(code) < 2 {
		return OpHalt, Param{}, 0, fmt.Errorf("requires 1 argument")
	}

	opcode, m1, _, _ := ParseInstructionCode(code[0])
	return opcode,
		Param{m1, code[1]},
		2,
		nil
}

func Decode2(code Code) (op OpCode, a1, a2 Param, advance int64, err error) {
	if len(code) < 3 {
		return OpHalt, Param{}, Param{}, 0, fmt.Errorf("requires 2 arguments")
	}

	opcode, m1, m2, _ := ParseInstructionCode(code[0])
	return opcode,
		Param{m1, code[1]},
		Param{m2, code[2]},
		3,
		nil
}

func Decode3(code Code) (op OpCode, a1, a2, a3 Param, advance int64, err error) {
	if len(code) < 4 {
		return OpHalt, Param{}, Param{}, Param{}, 0, fmt.Errorf("requires 3 arguments")
	}

	opcode, m1, m2, m3 := ParseInstructionCode(code[0])
	return opcode,
		Param{m1, code[1]},
		Param{m2, code[2]},
		Param{m3, code[3]},
		4,
		nil
}
