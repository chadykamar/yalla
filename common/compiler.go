package common

import (
	"fmt"
)

// Compiler provides functions for compilation
type Compiler struct {
	scanner *Scanner
}

func (c *Compiler) compile() {

	line := -1
	for {
		token := c.scanner.scanToken()
		if token.line != line {
			fmt.Printf("%4d ", token.line)
			line = token.line
		} else {
			fmt.Print("   | ")
		}
		fmt.Printf("%2d '%s'\n", token.tokenType, token.str)

		if token.tokenType == TokenEOF {
			break
		}
	}
}
