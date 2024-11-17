package main

import "errors"

type Stack struct {
	contents []int
}

func (s *Stack) Push(n int) {
	s.contents = append(s.contents, n)
}

func (s *Stack) Pop() (int, error) {
	if len(s.contents) <= 0 {
		return 0, errors.New("Can not pop a Stack with no items!")
	}

	element := len(s.contents) - 1
	popvalue := s.contents[element]

	s.contents = s.contents[:element]

	return popvalue, nil
}
