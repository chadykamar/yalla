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
	c.Write(OpReturn, 123)
	builder := strings.Builder{}

	d := Disassembler{c, &builder}

	d.Disassemble("test chunk")

	expectedString := `== test chunk ==
0000  123 OpConstant          0 '1.2'
0002    | OpReturn
`

	actualString := builder.String()

	assert.Equal(t, expectedString, actualString)
	// if expectedString != actualString {
	// 	t.Errorf("Expected:\n%s\n\nGot:\n%s", expectedString, actualString)
	// }

}
