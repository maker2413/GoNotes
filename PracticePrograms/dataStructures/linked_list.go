package main

import "errors"

// Node is a struct that will represent the individual nodes of our linked list.
type Node struct {
	content any
	next    *Node
}

// LinkedList is a struct that we will use to implement a LinkedList of Nodes.
type LinkedList struct {
	head *Node
	tail *Node
}

// IsEmpty is a method that will return a bool representing if the LinkedList
// is empty.
func (l LinkedList) IsEmpty() bool {
	return l.head == nil
}

// Len is a method that will return the length of our LinkedList.
func (l LinkedList) Len() int {
	if !l.IsEmpty() {
		currentNode := l.head
		count := 0
		for currentNode != nil {
			count++
			currentNode = currentNode.next
		}
		return count
	}
	return 0
}

// InsertFront is a method that allows us to add content to the head of our
// LinkedList.
func (l *LinkedList) InsertFront(n any) {
	switch l.Len() {
	case 0:
		newNode := &Node{content: n}
		l.head = newNode
		l.tail = newNode
	default:
		newNode := &Node{content: n}
		newNode.next = l.head
		l.head = newNode
	}
}

// InsertBack is a method that allows us to add content to the tail of our
// LinkedList.
func (l *LinkedList) InsertBack(n any) {
	switch l.Len() {
	case 0:
		newNode := &Node{content: n}
		l.head = newNode
		l.tail = newNode
	default:
		newNode := &Node{content: n}
		l.tail.next = newNode
		l.tail = newNode
	}
}

// RemoveFront is a method that allows us to remove content from the head of our
// LinkedList.
func (l *LinkedList) RemoveFront() (any, error) {
	var frontNode = &Node{}
	switch l.Len() {
	case 0:
		var zero any
		return zero, errors.New("Can not remove from empty list")
	case 1:
		frontNode = l.head
		l.head = nil
		l.tail = nil
	default:
		frontNode = l.head
		l.head = l.head.next
	}
	return frontNode.content, nil
}

// RemoveBack is a method that allows us to remove content from the tail of our
// LinkedList.
func (l *LinkedList) RemoveBack() (any, error) {
	var backNode = &Node{}
	switch l.Len() {
	case 0:
		var zero any
		return zero, errors.New("Can not remove from empty list")
	case 1:
		backNode = l.tail
		l.tail = nil
		l.head = nil
	default:
		backNode = l.head
		for backNode.next != l.tail {
			backNode = backNode.next
		}
		l.tail = backNode
		backNode = backNode.next
		l.tail.next = nil
	}
	return backNode.content, nil
}
