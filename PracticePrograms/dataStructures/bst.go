package main

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
