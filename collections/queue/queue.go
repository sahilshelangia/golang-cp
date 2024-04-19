package queue

import (
	"errors"
)

// Queue represents the queue data structure.
type Queue[T any] struct {
	items []T
}

// New returns a new instance of Queue.
func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

// Push adds an element back of the queue.
func (q *Queue[T]) Push(item T) {
	q.items = append(q.items, item)
}

// Pop removes the first element from the queue.
func (q *Queue[T]) Pop() (T, error) {
	if len(q.items) == 0 {
		var zeroValue T
		return zeroValue, errors.New("queue is empty")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

// Empty returns true if queue is empty, else returns false.
func (q *Queue[T]) Empty() bool {
	return len(q.items) == 0
}

// Size returns the size of the queue.
func (q *Queue[T]) Size() int {
	return len(q.items)
}

// Front returns the front element of the queue.
func (q *Queue[T]) Front() (T, error) {
	if len(q.items) == 0 {
		var zeroValue T
		return zeroValue, errors.New("queue is empty")
	}
	return q.items[0], nil
}

// Back returns the last element of the queue.
func (q *Queue[T]) Back() (T, error) {
	if len(q.items) == 0 {
		var zeroValue T
		return zeroValue, errors.New("queue is empty")
	}
	return q.items[len(q.items)-1], nil
}
