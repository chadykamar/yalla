package common

import (
	"fmt"
	"io"
	"log"
	"os"
)

// InterpretResult exists
type InterpretResult int

const (
	// InterpretOk signifies when the VM interprets a Chunk successfully
	InterpretOk InterpretResult = iota
	// InterpretCompileError signifies when the VM runs into an error at compile time
	InterpretCompileError
	// InterpretRuntimeError signifies when the VM runs into an error at runtime
	InterpretRuntimeError
)

var vm VirtualMachine

// VirtualMachine executes programs and holds their state
type VirtualMachine struct {
	Chunk              *Chunk
	InstructionCounter byte
	Stack              Stack
	Disassembler
}

// NewVirtualMachine will initialize a VirtualMachine struct
func NewVirtualMachine(writer io.Writer) *VirtualMachine {
	var disassembler Disassembler
	if writer != nil {
		disassembler = Disassembler{writer}
	} else {
		disassembler = Disassembler{os.Stdout}
	}

	vm := VirtualMachine{InstructionCounter: 0, Stack: Stack{}, Disassembler: disassembler}
	return &vm

}

// Interpret will execute the given bytecode
func (vm *VirtualMachine) Interpret(chunk *Chunk) (InterpretResult, error) {
	vm.Chunk = chunk
	vm.InstructionCounter = 0
	return vm.run()

}

func (vm *VirtualMachine) printStack() {
	fmt.Fprintf(vm.Writer, "          ")
	for i := 0; i < len(vm.Stack.values); i++ {
		fmt.Fprintf(vm.Writer, "[ %v ]", vm.Stack.values[i])
	}
	fmt.Println("")
}

func (vm *VirtualMachine) run() (InterpretResult, error) {

	for {
		// if len(vm.Stack.values) > 0 {
		vm.printStack()
		// }
		vm.disassembleInstruction(*vm.Chunk, int(vm.InstructionCounter))
		instruction := vm.readByte()

		switch instruction {
		case OpReturn:
			{
				val, err := vm.Stack.Pop()
				if err != nil {
					log.Fatal(err)
				}
				fmt.Fprintf(vm.Writer, "%v\n", val)

			}
		case OpConstant:
			{
				constant := vm.readConstant()
				vm.Stack.Push(constant)

			}
		case OpNegate:
			{
				val, err := vm.Stack.Pop()
				if err != nil {
					log.Fatal(err)
				}
				vm.Stack.Push(-val)
			}
		case OpAdd:
			vm.add()
		case OpSubtract:
			vm.subtract()
		case OpMultiply:
			vm.multiply()
		case OpDivide:
			vm.divide()

		}
	}

}

func (vm *VirtualMachine) readConstant() Value {
	return vm.Chunk.constants[vm.readByte()]
}

func (vm *VirtualMachine) readByte() byte {
	b := vm.Chunk.code[vm.InstructionCounter]
	vm.InstructionCounter++
	return b
}

func (vm *VirtualMachine) add() {
	a, err := vm.Stack.Pop()
	if err != nil {
		log.Fatal(err) // TODO Handle this error
	}
	b, err := vm.Stack.Pop()
	if err != nil {
		log.Fatal(err) // TODO Handle this error
	}
	vm.Stack.Push(a + b)
}
func (vm *VirtualMachine) subtract() {
	a, err := vm.Stack.Pop()
	if err != nil {
		log.Fatal(err) // TODO Handle this error
	}
	b, err := vm.Stack.Pop()
	if err != nil {
		log.Fatal(err) // TODO Handle this error
	}
	vm.Stack.Push(a - b)
}
func (vm *VirtualMachine) multiply() {
	a, err := vm.Stack.Pop()
	if err != nil {
		log.Fatal(err) // TODO Handle this error
	}
	b, err := vm.Stack.Pop()
	if err != nil {
		log.Fatal(err) // TODO Handle this error
	}
	vm.Stack.Push(a * b)
}
func (vm *VirtualMachine) divide() {
	a, err := vm.Stack.Pop()
	if err != nil {
		log.Fatal(err) // TODO Handle this error
	}
	b, err := vm.Stack.Pop()
	if err != nil {
		log.Fatal(err) // TODO Handle this error
	}
	vm.Stack.Push(a / b)
}
