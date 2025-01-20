package main

import "testing"

func TestBST(t *testing.T) {
	t.Run("Test Binary Search Tree", func(t *testing.T) {
		bst := BST{}
		assertTrue(t, bst.IsEmpty())

		var emptyTree []int
		assertDeepEqual(t, bst.InOrder(), emptyTree)
		assertDeepEqual(t, bst.PreOrder(), emptyTree)
		assertDeepEqual(t, bst.PostOrder(), emptyTree)

		assertNotError(t, bst.Insert(5))
		assertNotError(t, bst.Insert(2))
		assertNotError(t, bst.Insert(11))
		assertNotError(t, bst.Insert(12))
		assertError(t, bst.Insert(2))
	})
}
