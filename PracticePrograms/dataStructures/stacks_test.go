package main

import (
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("Test Integers", func(t *testing.T) {
		got := new(Stack[int])

		AssertTrue(t, got.IsEmpty())
		got.Push(4)
		AssertFalse(t, got.IsEmpty())
		AssertEqual(t, got.contents[0], 4)
		got.Push(8)
		got.Push(2)
		AssertEqual(t, len(got.contents), 3)
		value, _ := got.Pop()
		AssertEqual(t, value, 2)
	})
	t.Run("Test Strings", func(t *testing.T) {
		got := new(Stack[string])

		AssertTrue(t, got.IsEmpty())
		got.Push("hello")
		AssertFalse(t, got.IsEmpty())
		AssertEqual(t, got.contents[0], "hello")
		got.Push("world")
		got.Push("!")
		got.Push("!")
		AssertEqual(t, len(got.contents), 4)
		value, _ := got.Pop()
		AssertEqual(t, value, "!")
	})
	t.Run("Test Floats", func(t *testing.T) {
		got := new(Stack[float32])

		AssertTrue(t, got.IsEmpty())
		got.Push(3.14)
		AssertFalse(t, got.IsEmpty())
		AssertEqual(t, got.contents[0], 3.14)
		got.Push(2.789)
		got.Push(548.1239)
		AssertEqual(t, len(got.contents), 3)
		value, _ := got.Pop()
		AssertEqual(t, value, 548.1239)
	})
	t.Run("Can not Pop empty stack", func(t *testing.T) {
		s := new(Stack[int])

		_, err := s.Pop()

		if err == nil {
			t.Error("Should not be able to Pop an empty stack")
		}
	})
}
