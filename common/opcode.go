package common

const (
	// OpReturn is the operation used to return values from functions
	OpReturn = iota
	// OpConstant is the operation to denote the presence of a constant value
	OpConstant
	// OpNegate is the operation used to negate values
	OpNegate
	// OpAdd is the operation used to add two values
	OpAdd
	// OpSubtract is the operation used to subract two values
	OpSubtract
	// OpMultiply is the operation used to multiply two values
	OpMultiply
	//OpDivide is the operation used to divide two values
	OpDivide
)
