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

	instruction := d.Chunk.code[offset]

	switch instruction {
	case OpReturn:
		return d.simpleInstruction("OpReturn", offset)

	default:
		fmt.Fprintf(d.Writer, "Unknown opcode %d\n", instruction)
		return offset + 1
	}

}

func (d *Disassembler) simpleInstruction(name string, offset int) int {
	fmt.Fprintf(d.Writer, "%s\n", name)
	return offset + 1
}
