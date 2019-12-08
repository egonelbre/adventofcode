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

type Instr interface {
	Exec(cpu *Computer) error
	Decode(code Code) (Instr, int64, error)
}

type OpCode int64

var Ops = map[OpCode]Instr{}

func RegisterOp(instr Instr, code OpCode) {
	_, exists := Ops[code]
	if exists {
		panic(fmt.Sprintf("code %d already registered", code))
	}

	Ops[code] = instr
}

const (
	OpAdd      = OpCode(1)
	OpMultiply = OpCode(2)

	OpInput  = OpCode(3)
	OpOutput = OpCode(4)

	OpJumpIfTrue  = OpCode(5)
	OpJumpIfFalse = OpCode(6)

	OpLessThan = OpCode(7)
	OpEquals   = OpCode(8)
)

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

func ParseInstructionCode(code int64) (op OpCode, imm1, imm2, imm3 bool) {
	// ABCDE
	//  1002
	//
	// DE - two digit opcode
	//  C - mode of 1st param
	//  B - mode of 2nd param
	//  A - mode of 3rd param

	return OpCode(code % 100),
		1 == (code/100)%10,
		1 == (code/1000)%10,
		1 == (code/10000)%10
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

	opcode, imm1, _, _ := ParseInstructionCode(code[0])
	return opcode,
		Param{imm1, code[1]},
		2,
		nil
}

func Decode2(code Code) (op OpCode, a1, a2 Param, advance int64, err error) {
	if len(code) < 3 {
		return OpHalt, Param{}, Param{}, 0, fmt.Errorf("requires 2 arguments")
	}

	opcode, imm1, imm2, _ := ParseInstructionCode(code[0])
	return opcode,
		Param{imm1, code[1]},
		Param{imm2, code[2]},
		3,
		nil
}

func Decode3(code Code) (op OpCode, a1, a2, a3 Param, advance int64, err error) {
	if len(code) < 4 {
		return OpHalt, Param{}, Param{}, Param{}, 0, fmt.Errorf("requires 3 arguments")
	}

	opcode, imm1, imm2, imm3 := ParseInstructionCode(code[0])
	return opcode,
		Param{imm1, code[1]},
		Param{imm2, code[2]},
		Param{imm3, code[3]},
		4,
		nil
}
