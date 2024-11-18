package main

import (
	"testing"
)

func TestQueue(t *testing.T) {
	t.Run("Test Integers", func(t *testing.T) {
		got := new(Queue[int])

		AssertTrue(t, got.IsEmpty())
		got.Enqueue(5)
		AssertFalse(t, got.IsEmpty())
		AssertEqual(t, got.contents[0], 5)
		got.Enqueue(9)
		got.Enqueue(2)
		got.Enqueue(1)
		AssertEqual(t, len(got.contents), 4)
		value, _ := got.Dequeue()
		AssertEqual(t, value, 5)
	})
	t.Run("Test Strings", func(t *testing.T) {
		got := new(Queue[string])

		AssertTrue(t, got.IsEmpty())
		got.Enqueue("hello")
		AssertFalse(t, got.IsEmpty())
		AssertEqual(t, got.contents[0], "hello")
		got.Enqueue("world")
		got.Enqueue("!")
		AssertEqual(t, len(got.contents), 3)
		value, _ := got.Dequeue()
		AssertEqual(t, value, "hello")
	})
	t.Run("Test Floats", func(t *testing.T) {
		got := new(Queue[float32])

		AssertTrue(t, got.IsEmpty())
		got.Enqueue(12.8417)
		AssertFalse(t, got.IsEmpty())
		AssertEqual(t, got.contents[0], 12.8417)
		got.Enqueue(2.1)
		got.Enqueue(0.241211237)
		AssertEqual(t, len(got.contents), 3)
		value, _ := got.Dequeue()
		AssertEqual(t, value, 12.8417)
	})
	t.Run("Con not Dequeue empty queue", func(t *testing.T) {
		q := new(Queue[int])

		_, err := q.Dequeue()

		if err == nil {
			t.Error("Should not be able to Dequeue empty queue")
		}
	})
}
