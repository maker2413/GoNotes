package main

import "testing"

func TestBST(t *testing.T) {
	t.Run("Test Binary Search Tree", func(t *testing.T) {
		bst := BST{}
		AssertTrue(t, bst.IsEmpty())
	})
}
