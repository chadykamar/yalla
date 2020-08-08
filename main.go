package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/chadykamar/yalla/common"
)

func main() {

	args := os.Args
	argc := len(args)
	if argc == 1 {
		repl()
	} else if argc == 2 {
		runFile(args[1])
	} else {
		fmt.Fprintln(os.Stderr, "Usage: yalla [path]")
	}

}

func repl() {

	fmt.Println("Yalla REPL")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		_, err := reader.ReadString('\n')
		vm := common.NewVirtualMachine(reader, os.Stdout)
		if err != nil {
			fmt.Println("Could not read line")
		}

		vm.Interpret()

	}

}

func runFile(filename string) {

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Could not read file %s", filename)
	}

	text := string(content)

	reader := strings.NewReader(text)
	vm := common.NewVirtualMachine(reader, os.Stdout)

	result, err := vm.Interpret()

	if result == common.InterpretCompileError {
		os.Exit(65)
	} else if result == common.InterpretRuntimeError {
		os.Exit(70)
	}

}
