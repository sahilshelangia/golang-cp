package stack

import (
	"errors"
)

// Stack represents the stack data structure.
type Stack[T any] struct {
	items []T
}

// New returns a new instance of Stack.
func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push adds an item to the top of the stack.
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop adds an item to the top of the stack.
func (s *Stack[T]) Pop() (T, error) {
	if len(s.items) == 0 {
		// If T is a value type, return its zero value
		var zeroValue T
		return zeroValue, errors.New("stack is empty")
	}
	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return item, nil
}

// Top returns the top element of the stack. return error, if stack is empty.
func (s *Stack[T]) Top() (T, error) {
	if len(s.items) == 0 {
		// If T is a value type, return its zero value
		var zeroValue T
		return zeroValue, errors.New("stack is empty")
	}

	return s.items[len(s.items)-1], nil
}

// Size returns the number of items in the stack.
func (s *Stack[T]) Size() int {
	return len(s.items)
}

// Empty returns true if Stack is empty, else returns false.
func (s *Stack[T]) Empty() bool {
	return len(s.items) == 0
}
