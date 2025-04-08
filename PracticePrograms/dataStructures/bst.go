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

// Height is a method that is a warpper around the node method height().
func (b *BST) Height() int {
	if b.IsEmpty() {
		return 0
	}

	return b.root.height()
}

// height is a node method that returns the height or depth of the tree.
func (n *BSTNode) height() int {
	if n.left == nil && n.right == nil {
		return 1
	}

	leftHeight := 0
	rightHeight := 0
	if n.left != nil {
		leftHeight = n.left.height()
	}
	if n.right != nil {
		rightHeight = n.right.height()
	}

	if leftHeight > rightHeight {
		return leftHeight + 1
	} else {
		return rightHeight + 1
	}
}

// InOrder is a method that is a wrapper around the node method inOrder().
func (b *BST) InOrder() []int {
	if b.IsEmpty() {
		var tree []int

		return tree
	}

	return b.root.inOrder()
}

// inOrder is a node method that is called recursively to traverse the tree
// in order: https://www.geeksforgeeks.org/inorder-traversal-of-binary-tree/.
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

// PreOrder is a method that is a wrapper around the node method preOrder().
func (b *BST) PreOrder() []int {
	if b.IsEmpty() {
		var tree []int

		return tree
	}

	return b.root.preOrder()
}

// preOrder is a node method that is called recursively to traverse the tree
// in pre order: https://www.geeksforgeeks.org/preorder-traversal-of-binary-tree/.
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

// PostOrder is a method that is a wrapper around the node method postOrder().
func (b *BST) PostOrder() []int {
	if b.IsEmpty() {
		var tree []int

		return tree
	}

	return b.root.postOrder()
}

// postOrder is a node method that is called recursively to traverse the tree
// in post order: https://www.geeksforgeeks.org/postorder-traversal-of-binary-tree/.
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

// Insert is a method that is a wrapper around the node method insert().
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

// insert is a node method that is called recursively to insert a new node into
// the tree if it doesn't already exist. If the node already exists an error is
// returned.
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

// Min is a method that returns the minimum value in the tree.
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

// Max is a method that returns the maximum value in the tree.
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

// Exists is a method that is a wrapper around the node method exists().
func (b *BST) Exists(val int) bool {
	if b.IsEmpty() {
		return false
	}

	return b.root.exists(val)
}

// exists is a node method that is called recursively to traverse the tree
// and return a bool representing if a val exists currently in the tree.
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

// Delete is a method that is a wrapper around the node method delete().
func (b *BST) Delete(val int) error {
	if b.IsEmpty() {
		return errors.New("Tree is already empty")
	}

	if b.root.val == val {
		if b.root.right != nil {
			temp := b.root
			b.root = b.root.right
			return temp.delete(val)
		} else if b.root.left != nil {
			temp := b.root
			b.root = b.root.left
			return temp.delete(val)
		} else {
			b.root = nil
			return nil
		}
	}

	return b.root.delete(val)
}

func (n *BSTNode) delete(val int) error {
	// if self.val is None:
	//         return None
	// if val < self.val:
	//     if self.left:
	//         self.left = self.left.delete(val)
	//     return self
	// if val > self.val:
	// 	if self.right:
	// 		self.right = self.right.delete(val)
	// 	return self
	// if self.right is None:
	// 	return self.left
	// if self.left is None:
	// 	return self.right
	// min_larger_node = self.right
	// while min_larger_node.left:
	// 	min_larger_node = min_larger_node.left
	// self.val = min_larger_node.val
	// self.right = self.right.delete(min_larger_node.val)
	// return self

	return nil
}
