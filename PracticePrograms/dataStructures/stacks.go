package main

import "errors"

type Stack[T any] struct {
	contents []T
}

func (s *Stack[T]) Push(n T) {
	s.contents = append(s.contents, n)
}

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

func (s *Stack[T]) IsEmpty() bool {
	return len(s.contents) <= 0
}
