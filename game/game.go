package game

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Opcode int
type Operand int

const (
	Acc Opcode = iota
	Jmp
	Nop
)

type Instruction struct {
	opcode  Opcode
	operand Operand
}

type AddressingMode int

const (
	Ind AddressingMode = 0
	Imm                = 1
	Rel                = 2
)

type Program struct {
	text   []Instruction
	memory []Instruction
	ip     int
	acc    int
	bp     int
	halted bool
	reader *bufio.Reader
	output []int
	debug  bool
}

func (o Opcode) String() string {
	switch o {
	case Acc:
		return "acc"
	case Jmp:
		return "jmp"
	case Nop:
		return "nop"
	default:
		return "Unknown"
	}
}

func (a AddressingMode) String() string {
	switch a {
	case Ind:
		return "Ind"
	case Imm:
		return "Imm"
	case Rel:
		return "Rel"
	default:
		return "Unknown"
	}
}

func (p *Program) SetDebug(val bool) {
	p.debug = val
}

func (p *Program) SetReader(reader io.Reader) {
	p.reader = bufio.NewReader(reader)
}

func (p *Program) SetReaderFromInts(ints ...int) {
	var sb strings.Builder
	for _, i := range ints {
		fmt.Fprintf(&sb, "%d\n", i)
	}
	p.SetReader(strings.NewReader(sb.String()))
}

func (p *Program) Reset() {
	p.acc = 0
	p.ip = 0
	p.bp = 0
	p.output = nil
	p.reader = nil
	p.halted = false
	/*p.memory = make([]Instruction, len(p.text))
	for i := range p.text {
		p.memory[i] = p.text[i]
	}*/
}

func opcodeFromString(opcodeStr string) Opcode {
	switch opcodeStr {
	case "acc":
		return Acc
	case "jmp":
		return Jmp
	case "nop":
		return Nop
	default:
		panic(fmt.Sprintf("invalid opcode: %s", opcodeStr))
	}
}

func (p *Program) PatchOpcode(ip int, opcode Opcode) {
	p.text[ip] = Instruction{opcode: opcode, operand: p.text[ip].operand}
}

func (p *Program) Len() int {
	return len(p.text)
}

func (p *Program) GetHalted() bool {
	return p.halted
}

func (p *Program) InitStateFromFile(filename string) {
	dat, err := ioutil.ReadFile(filename)
	check(err)

	stringArray := strings.Split(string(dat), "\n")

	// copy text section
	p.text = make([]Instruction, len(stringArray))
	for i := 0; i < len(stringArray); i++ {
		if len(stringArray[i]) == 0 || stringArray[i][0] == '\n' {
			continue
		}
		var opcodeStr string
		var operand int
		fmt.Sscanf(stringArray[i], "%s %d\n", &opcodeStr, &operand)
		instruction := Instruction{opcode: opcodeFromString(opcodeStr), operand: Operand(operand)}
		p.text[i] = instruction
	}

	p.Reset()
	return
}

func (p *Program) InitStateFromProgram(other *Program) {
	if len(p.text) != len(other.text) {
		p.text = make([]Instruction, len(other.text))
	}
	copy(p.text, other.text)

	if len(p.output) != len(other.output) {
		p.output = make([]int, len(other.output))
	}
	copy(p.output, other.output)

	p.acc = other.acc
	p.ip = other.ip
	p.bp = other.bp
	p.halted = other.halted
}

func (p *Program) SetAcc(acc int) {
	p.acc = acc
}

func (p *Program) SetIp(ip int) {
	p.ip = ip
}

func (p *Program) GetIp() int {
	return p.ip
}

func (p *Program) GetAcc() int {
	return p.acc
}

func (p *Program) IncrementIp(amount int) {
	p.ip += amount
}

func (p *Program) GetOpcode(ip int) Opcode {
	return p.text[ip].opcode
}

func (p *Program) GetOperand(ip int) Operand {
	return p.text[ip].operand
}

func (p *Program) Step() {
	if p.ip >= len(p.text) {
		p.halted = true
	}
	if p.halted {
		return
	}
	opcode := p.GetOpcode(p.ip)
	switch opcode {
	case Acc:
		operand := p.GetOperand(p.ip)
		if p.debug {
			fmt.Printf("acc %d\n", int(operand))
		}
		p.acc += int(operand)
		p.ip += 1
	case Jmp:
		operand := p.GetOperand(p.ip)
		if p.debug {
			fmt.Printf("jmp %d\n", int(operand))
		}
		p.ip += int(operand)
	case Nop:
		p.ip += 1
	default:
		panic(fmt.Sprintf("encountered unknown opcode: %d", opcode))
	}
}

func (p *Program) StepBy(steps int) {
	for i := 0; i < steps; i++ {
		p.Step()
	}
}

func (p *Program) Run() []int {
	for !p.halted {
		p.Step()
	}

	return p.output
}
