package common

import (
	"errors"
)

// Stack is the virtual machine's stack
type Stack struct {
	values []Value
}

// Peek returns the Value on top of the stack
func (s *Stack) Peek() (Value, error) {
	var val Value
	if len(s.values) == 0 {
		return val, errors.New("Invalid Operation Error: Tried to peek empty stack")
	}
	return s.values[len(s.values)-1], nil
}

// Pop removes the value on top of the stack and returns it
func (s *Stack) Pop() (Value, error) {
	var val Value
	n := len(s.values)
	if n == 0 {
		return val, errors.New("Invalid Operation Error: Tried to pop empty stack")
	}
	val = s.values[n-1]
	s.values = s.values[:n-1]
	return val, nil
}

// Push adds a new value on top of the stack
func (s *Stack) Push(value Value) {
	s.values = append(s.values, value)

}
