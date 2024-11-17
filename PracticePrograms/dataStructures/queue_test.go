package main

import (
	"reflect"
	"testing"
)

func TestQueue(t *testing.T) {
	t.Run("Enqueue Integers", func(t *testing.T) {
		got := new(Queue[int])
		want := []int{5, 9, 2, 1}

		got.Enqueue(5)
		got.Enqueue(9)
		got.Enqueue(2)
		got.Enqueue(1)

		if !reflect.DeepEqual(got.contents, want) {
			t.Errorf("expected: '%d', got '%d'", want, got.contents)
		}
	})
	t.Run("Enqueue Strings", func(t *testing.T) {
		got := new(Queue[string])
		want := []string{"hello", "world", "!"}

		got.Enqueue("hello")
		got.Enqueue("world")
		got.Enqueue("!")

		if !reflect.DeepEqual(got.contents, want) {
			t.Errorf("expected: '%s', got '%s'", want, got.contents)
		}
	})
	t.Run("Dequeue", func(t *testing.T) {
		got := new(Queue[int])
		want := []int{9, 2, 1}

		got.Enqueue(5)
		got.Enqueue(9)
		got.Enqueue(2)
		got.Enqueue(1)

		dequeue, err := got.Dequeue()

		if err != nil {
			t.Fatal(err)
		}

		if dequeue != 5 {
			t.Errorf("expected to dequeue: 5, got: %d", dequeue)
		}

		if !reflect.DeepEqual(got.contents, want) {
			t.Errorf("expected: '%d', got '%d'", want, got.contents)
		}
	})
}
