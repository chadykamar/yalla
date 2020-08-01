package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/chadykamar/yalla/common"
)

func main() {
	c := common.NewChunk([]byte{})

	constantIndex := c.AddConstant(1.2)
	c.Write(common.OpConstant, 123)
	c.Write(constantIndex, 123)

	constantIndex = c.AddConstant(3.4)
	c.Write(common.OpConstant, 123)
	c.Write(constantIndex, 123)

	c.Write(common.OpAdd, 123)

	constantIndex = c.AddConstant(5.6)
	c.Write(common.OpConstant, 123)
	c.Write(constantIndex, 123)

	c.Write(common.OpDivide, 123)
	c.Write(common.OpNegate, 123)

	c.Write(common.OpReturn, 123)
	builder := strings.Builder{}

	d := common.Disassembler{Writer: &builder}
	d.Disassemble(*c, "test chunk")
	fmt.Println(builder.String())

	vm := common.NewVirtualMachine(os.Stdout)

	vm.Interpret(c)

}
