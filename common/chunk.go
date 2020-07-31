package common

// Chunk contains bytecode and stores constants and line information
// during lexing
type Chunk struct {
	code      []byte
	constants []Value
	lines     []int
}

// NewChunk initializes a Chunk
func NewChunk(code []byte) *Chunk {

	c := Chunk{}
	c.code = code

	return &c

}

func (c *Chunk) Write(data byte, line int) {
	c.code = append(c.code, data)
	c.lines = append(c.lines, line)
}

// AddConstant returns the index where the constant value as appended so that we can
// locate that same constant later
func (c *Chunk) AddConstant(value Value) byte {
	c.constants = append(c.constants, value)
	return byte(len(c.constants) - 1)

}
