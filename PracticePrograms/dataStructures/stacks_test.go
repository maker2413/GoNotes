package main

import (
	"reflect"
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("Push", func(t *testing.T) {
		got := Stack{}
		want := []int{4, 8, 2}

		got.Push(4)
		got.Push(8)
		got.Push(2)
		if !reflect.DeepEqual(got.contents, want) {
			t.Errorf("expected: '%d', but got '%d'", want, got.contents)
		}
	})
	t.Run("Pop", func(t *testing.T) {
		got := Stack{[]int{4, 8, 2}}
		want := []int{4, 8, 2}

		got.Push(9)
		popvalue, err := got.Pop()

		if err != nil {
			t.Fatal(err)
		}

		if popvalue != 9 {
			t.Errorf("expected pop to be: 9, got: '%d'", popvalue)
		}

		if !reflect.DeepEqual(got.contents, want) {
			t.Errorf("expected: '%d', but got '%d'", want, got.contents)
		}
	})
	t.Run("Can not Pop empty stack", func(t *testing.T) {
		s := Stack{}

		_, err := s.Pop()

		if err == nil {
			t.Errorf("Should not be able to pop an empty stack")
		}
	})
}
