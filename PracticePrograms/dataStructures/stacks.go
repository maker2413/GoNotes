package main

import "errors"

// Stack is a struct that we will use to implement a stack.
type Stack[T any] struct {
	contents []T
}

// Push is the method that allows us to push content onto the Stack.
func (s *Stack[T]) Push(n T) {
	s.contents = append(s.contents, n)
}

// Pop is the method that allows us to pop content off of the Stack.
func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("Can not pop a Stack with no items!")
	}

	element := len(s.contents) - 1
	popvalue := s.contents[element]

	s.contents = s.contents[:element]

	return popvalue, nil
}

// IsEmpty is a method that will return a bool representing if the Stack is
// empty.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.contents) <= 0
}
