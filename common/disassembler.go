package common

import (
	"fmt"
	"io"
)

// Disassembler translates byte code into human-readable instructions to help with debugging
type Disassembler struct {
	Chunk  *Chunk
	Writer io.Writer
}

// Disassemble traces the instructions of the given chunk line by line
func (d *Disassembler) Disassemble(name string) {
	fmt.Fprintf(d.Writer, "== %s ==\n", name)

	for offset := 0; offset < len(d.Chunk.code); {
		offset = d.disassembleInstruction(offset)

	}

}

func (d *Disassembler) disassembleInstruction(offset int) int {
	fmt.Fprintf(d.Writer, "%04d ", offset)

	if offset > 0 && d.Chunk.lines[offset] == d.Chunk.lines[offset-1] {
		fmt.Fprint(d.Writer, "   | ")
	} else {
		fmt.Fprintf(d.Writer, "%4d ", d.Chunk.lines[offset])
	}

	instruction := d.Chunk.code[offset]

	switch instruction {
	case OpReturn:
		return d.simpleInstruction("OpReturn", offset)
	case OpConstant:
		return d.constantInstruction("OpConstant", offset)

	default:
		fmt.Fprintf(d.Writer, "Unknown opcode %d\n", instruction)
		return offset + 1
	}

}

func (d *Disassembler) simpleInstruction(name string, offset int) int {
	fmt.Fprintf(d.Writer, "%s\n", name)
	return offset + 1
}

func (d *Disassembler) constantInstruction(name string, offset int) int {
	constantIndex := d.Chunk.code[offset+1]
	fmt.Fprintf(d.Writer, "%-16s %4d '%v'\n", name, constantIndex, d.Chunk.constants[constantIndex])
	return offset + 2
}
