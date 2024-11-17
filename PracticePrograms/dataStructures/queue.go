package main

import "errors"

type Queue[T any] struct {
	contents []T
}

func (q *Queue[T]) Enqueue(n T) {
	q.contents = append(q.contents, n)
}

func (q *Queue[T]) Dequeue() (T, error) {
	if len(q.contents) <= 0 {
		var zero T
		return zero, errors.New("Can not dequeue an empty Queue")
	}

	dequeuevalue := q.contents[0]
	q.contents = q.contents[1:]

	return dequeuevalue, nil
}
