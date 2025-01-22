package main

import "testing"

func TestBST(t *testing.T) {
	t.Run("Test Binary Search Tree", func(t *testing.T) {
		bst := BST{}
		assertTrue(t, bst.IsEmpty())

		// Test empty tree
		var emptyTree []int
		assertDeepEqual(t, bst.InOrder(), emptyTree)
		assertDeepEqual(t, bst.PreOrder(), emptyTree)
		assertDeepEqual(t, bst.PostOrder(), emptyTree)
		_, err := bst.Min()
		assertError(t, err)
		_, err = bst.Max()
		assertError(t, err)
		assertFalse(t, bst.Exists(1))
		assertEqual(t, bst.Height(), 0)

		// Test inserting tree
		assertNotError(t, bst.Insert(5))
		assertEqual(t, bst.Height(), 1)

		assertNotError(t, bst.Insert(2))
		assertNotError(t, bst.Insert(11))
		assertNotError(t, bst.Insert(12))
		assertError(t, bst.Insert(2))

		// Test printing tree
		assertDeepEqual(t, bst.InOrder(), []int{2, 5, 11, 12})
		assertDeepEqual(t, bst.PreOrder(), []int{5, 2, 11, 12})
		assertDeepEqual(t, bst.PostOrder(), []int{2, 11, 12, 5})

		// Test min and max values in tree
		got, err := bst.Min()
		assertEqual(t, got, 2)
		assertNotError(t, err)
		got, err = bst.Max()
		assertEqual(t, got, 12)
		assertNotError(t, err)

		// Test searching in tree
		assertTrue(t, bst.Exists(5))
		assertFalse(t, bst.Exists(6))
		assertTrue(t, bst.Exists(2))

		// Test height of tree
		assertEqual(t, bst.Height(), 3)
		assertNotError(t, bst.Insert(3))
		assertNotError(t, bst.Insert(4))
		assertEqual(t, bst.Height(), 4)

		// Test deleting of tree node
		assertDeepEqual(t, bst.InOrder(), []int{2, 3, 4, 5, 11, 12})
		assertNotError(t, bst.Delete(2))
		assertDeepEqual(t, bst.InOrder(), []int{3, 4, 5, 11, 12})
	})
}
