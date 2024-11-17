package main

import (
	"reflect"
	"testing"
)

func TestLinkedList(t *testing.T) {
	t.Run("Create Linked List", func(t *testing.T) {
		got := new(LinkedList[int])
		want := []int{}

		if !reflect.DeepEqual(got.contents, want) {
			t.Errorf("expected: '%d', got: '%d'", got.contents, want)
		}
	})
}
