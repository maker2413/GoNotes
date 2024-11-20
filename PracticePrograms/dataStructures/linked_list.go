package main

import "errors"

type Node struct {
	content any
	next    *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (l LinkedList) IsEmpty() bool {
	return l.head == nil
}

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
