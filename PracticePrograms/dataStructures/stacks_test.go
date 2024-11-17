package main

import (
	"reflect"
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("Push Integers", func(t *testing.T) {
		got := new(Stack[int])
		want := []int{4, 8, 2}

		got.Push(4)
		got.Push(8)
		got.Push(2)
		if !reflect.DeepEqual(got.contents, want) {
			t.Errorf("expected: '%d', but got '%d'", want, got.contents)
		}
	})
	t.Run("Push Strings", func(t *testing.T) {
		got := new(Stack[string])
		want := []string{"hello", "world", "!"}

		got.Push("hello")
		got.Push("world")
		got.Push("!")
		if !reflect.DeepEqual(got.contents, want) {
			t.Errorf("expected: '%s', but got '%s'", want, got.contents)
		}
	})
	t.Run("Push Floats", func(t *testing.T) {
		got := new(Stack[float32])
		want := []float32{3.14, 2.789, 548.1239}

		got.Push(3.14)
		got.Push(2.789)
		got.Push(548.1239)
		if !reflect.DeepEqual(got.contents, want) {
			t.Errorf("expected: '%f', but got '%f'", want, got.contents)
		}
	})
	t.Run("Pop", func(t *testing.T) {
		got := new(Stack[int])
		want := []int{4, 8, 2}

		got.Push(4)
		got.Push(8)
		got.Push(2)
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
		s := new(Stack[int])

		_, err := s.Pop()

		if err == nil {
			t.Errorf("Should not be able to pop an empty stack")
		}
	})
}
