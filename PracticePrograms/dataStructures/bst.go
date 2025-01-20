package main

import (
	"errors"
)

// BST is a struct that we will use to implement a Binary Search Tree.
type BST struct {
	root *BSTNode
}

// BSTNode is a struct that we will use to represent the individual nodes of
// our Binary Search Tree.
type BSTNode struct {
	val   int
	left  *BSTNode
	right *BSTNode
}

// IsEmpty is a method that will return a bool representing if the BST is empty.
func (b *BST) IsEmpty() bool {
	return b.root == nil
}

func (b *BST) InOrder() []int {
	if b.IsEmpty() {
		var tree []int

		return tree
	}

	return b.root.inOrder()
}

func (n *BSTNode) inOrder() []int {
	var tree []int

	if n.left != nil {
		for _, n := range n.left.inOrder() {
			tree = append(tree, n)
		}
	}

	tree = append(tree, n.val)

	if n.right != nil {
		for _, n := range n.right.inOrder() {
			tree = append(tree, n)
		}
	}

	return tree
}

func (b *BST) PreOrder() []int {
	if b.IsEmpty() {
		var tree []int

		return tree
	}

	return b.root.preOrder()
}

func (n *BSTNode) preOrder() []int {
	var tree []int

	tree = append(tree, n.val)

	if n.left != nil {
		for _, n := range n.left.preOrder() {
			tree = append(tree, n)
		}
	}

	if n.right != nil {
		for _, n := range n.right.preOrder() {
			tree = append(tree, n)
		}
	}

	return tree
}

func (b *BST) PostOrder() []int {
	if b.IsEmpty() {
		var tree []int

		return tree
	}

	return b.root.postOrder()
}

func (n *BSTNode) postOrder() []int {
	var tree []int

	if n.left != nil {
		for _, n := range n.left.preOrder() {
			tree = append(tree, n)
		}
	}

	if n.right != nil {
		for _, n := range n.right.preOrder() {
			tree = append(tree, n)
		}
	}

	tree = append(tree, n.val)

	return tree
}

func (b *BST) Insert(val int) error {
	node := BSTNode{
		val: val,
	}

	if b.root == nil {
		b.root = &node
		return nil
	}

	return b.root.insert(node)
}

func (n *BSTNode) insert(node BSTNode) error {
	if node.val == n.val {
		return errors.New("Node with value: %d, already exists")
	}

	if node.val < n.val {
		if n.left == nil {
			n.left = &node
			return nil
		}

		return n.left.insert(node)
	}

	if n.right == nil {
		n.right = &node
		return nil
	}

	return n.right.insert(node)
}

func (b *BST) Min() (int, error) {
	if b.IsEmpty() {
		return 0, errors.New("Tree is currently empty, unable to get minimum value")
	}

	current := b.root

	for current.left != nil {
		current = current.left
	}

	return current.val, nil
}

func (b *BST) Max() (int, error) {
	if b.IsEmpty() {
		return 0, errors.New("Tree is currently empty, unable to get maximum value")
	}

	current := b.root

	for current.right != nil {
		current = current.right
	}

	return current.val, nil
}

func (b *BST) Exists(val int) bool {
	if b.IsEmpty() {
		return false
	}

	return b.root.exists(val)
}

func (n *BSTNode) exists(val int) bool {
	if val == n.val {
		return true
	}

	if val < n.val && n.left != nil {
		return n.left.exists(val)
	}

	if val > n.val && n.right != nil {
		return n.right.exists(val)
	}

	return false
}
