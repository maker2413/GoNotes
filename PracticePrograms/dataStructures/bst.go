package main

import (
	"errors"
	"fmt"
)

type BST struct {
	root *BSTNode
}

type BSTNode struct {
	val   int
	left  *BSTNode
	right *BSTNode
}

func (b *BST) IsEmpty() bool {
	return b.root == nil
}

func (b *BST) InOrder() []int {
	var tree []int

	return tree
}

func (b *BST) PreOrder() []int {
	var tree []int

	return tree
}

func (b *BST) PostOrder() []int {
	var tree []int

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

	if node.val > n.val {
		if n.right == nil {
			n.right = &node
			return nil
		}

		return n.right.insert(node)
	}

	return errors.New(fmt.Sprintf("Not able to insert: %d", node.val))
}
