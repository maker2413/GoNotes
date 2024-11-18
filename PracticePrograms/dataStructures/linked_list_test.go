package main

import (
	"reflect"
	"testing"
)

func TestLinkedList(t *testing.T) {
	t.Run("Create Linked List", func(t *testing.T) {
		got := new(LinkedList[int])
		want := []int{1}

		got.contents = append(got.contents, 1)

		if !reflect.DeepEqual(got.contents, want) {
			t.Errorf("expected: '%q', got: '%q'", got.contents, want)
		}
	})
}
