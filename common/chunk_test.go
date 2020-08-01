package common

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisassemble(t *testing.T) {

	c := NewChunk([]byte{})

	constantIndex := c.AddConstant(1.2)
	c.Write(OpConstant, 123)
	c.Write(constantIndex, 123)

	constantIndex = c.AddConstant(3.4)
	c.Write(OpConstant, 123)
	c.Write(constantIndex, 123)

	c.Write(OpAdd, 123)

	constantIndex = c.AddConstant(5.6)
	c.Write(OpConstant, 123)
	c.Write(constantIndex, 123)

	c.Write(OpDivide, 123)
	c.Write(OpNegate, 123)

	c.Write(OpReturn, 123)
	builder := strings.Builder{}

	d := Disassembler{&builder}

	d.Disassemble(*c, "test chunk")

	expectedString := `== test chunk ==
0000  123 OpConstant          0 '1.2'
0002    | OpConstant          1 '3.4'
0004    | OpAdd
0005    | OpConstant          2 '5.6'
0007    | OpDivide
0008    | OpNegate
0009    | OpReturn
`

	actualString := builder.String()

	assert.Equal(t, expectedString, actualString)
	// if expectedString != actualString {
	// 	t.Errorf("Expected:\n%s\n\nGot:\n%s", expectedString, actualString)
	// }

}
