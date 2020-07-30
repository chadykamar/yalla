package common

import (
	"strings"
	"testing"
)

func TestDisassemble(t *testing.T) {

	c := NewChunk([]byte{OpReturn})
	builder := strings.Builder{}

	d := Disassembler{c, &builder}

	d.Disassemble("test chunk")

	expectedString := `== test chunk ==
0000 OpReturn`

	actualString := builder.String()
	if expectedString != actualString {
		t.Errorf("Expected:\n%s\n\nGot:\n%s", expectedString, actualString)
	}

}
