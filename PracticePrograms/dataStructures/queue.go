package main

import "errors"

// Queue is a struc that we will use to implement a queue.
type Queue[T any] struct {
	contents []T
}

// Enqueue is the method that allows us to add content to the Queue.
func (q *Queue[T]) Enqueue(n T) {
	q.contents = append(q.contents, n)
}

// Enqueue is the method that allows us to remove content from the Queue.
func (q *Queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("Can not dequeue an empty Queue")
	}

	dequeuevalue := q.contents[0]
	q.contents = q.contents[1:]

	return dequeuevalue, nil
}

// IsEmpty is a method that will return a bool representing if the Queue is
// empty.
func (q *Queue[T]) IsEmpty() bool {
	return len(q.contents) <= 0
}
