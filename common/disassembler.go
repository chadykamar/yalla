package common

import (
	"fmt"
	"io"
)

// Disassembler translates byte code into human-readable instructions to help with debugging
type Disassembler struct {
	Writer io.Writer
}

// Disassemble traces the instructions of the given chunk line by line
func (d *Disassembler) Disassemble(chunk Chunk, name string) {
	fmt.Fprintf(d.Writer, "== %s ==\n", name)

	for offset := 0; offset < len(chunk.code); {
		offset = d.disassembleInstruction(chunk, offset)

	}

}

func (d *Disassembler) disassembleInstruction(chunk Chunk, offset int) int {
	fmt.Fprintf(d.Writer, "%04d ", offset)

	if offset > 0 && chunk.lines[offset] == chunk.lines[offset-1] {
		fmt.Fprint(d.Writer, "   | ")
	} else {
		fmt.Fprintf(d.Writer, "%4d ", chunk.lines[offset])
	}

	instruction := chunk.code[offset]

	switch instruction {
	case OpReturn:
		return d.simpleInstruction("OpReturn", offset)
	case OpConstant:
		return d.constantInstruction(chunk, "OpConstant", offset)

	default:
		fmt.Fprintf(d.Writer, "Unknown opcode %d\n", instruction)
		return offset + 1
	}

}

func (d *Disassembler) simpleInstruction(name string, offset int) int {
	fmt.Fprintf(d.Writer, "%s\n", name)
	return offset + 1
}

func (d *Disassembler) constantInstruction(chunk Chunk, name string, offset int) int {
	constantIndex := chunk.code[offset+1]
	fmt.Fprintf(d.Writer, "%-16s %4d '%v'\n", name, constantIndex, chunk.constants[constantIndex])
	return offset + 2
}
