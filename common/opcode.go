package common

const (
	// OpReturn is the operation used to return values from functions
	OpReturn = iota
	// OpConstant is the operation to denote the presence of a constant value
	OpConstant
	// OpNegate is operation used to negate values
	OpNegate
)
