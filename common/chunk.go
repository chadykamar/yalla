package common

// Chunk contains bytecode and stores constants and line information
// during lexing
type Chunk struct {
	code []byte
}

// NewChunk initializes a Chunk
func NewChunk(code []byte) *Chunk {

	c := Chunk{}
	c.code = code

	return &c

}

func (c *Chunk) Write(data []byte) (int, error) {
	c.code = append(c.code, data...)
	return len(data), nil
}
