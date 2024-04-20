package queue

const _emptyQueueErrorString = "can't perform operation, because queue is empty"

// Queue represents the queue data structure. It's implemented using linked list.
type Queue[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

type node[T any] struct {
	value T
	next  *node[T]
}

// New returns a new instance of Queue.
// Time complexity: O(1)
func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

// Push adds an element at back of the queue.
// Time complexity: O(1)
func (queue *Queue[T]) Push(item T) {
	newNode := &node[T]{
		value: item,
	}
	if queue.size == 0 {
		queue.head = newNode
		queue.tail = newNode
		queue.size = 1
		return
	}

	queue.tail.next = newNode
	queue.tail = newNode
	queue.size++
}

// Pop removes the first element from the queue.
// Time complexity: O(1)
func (queue *Queue[T]) Pop() {
	if queue.size == 0 {
		panic(_emptyQueueErrorString)
	}
	queue.size--
	queue.head = queue.head.next
}

// Empty returns true if queue is empty, else returns false.
// Time complexity: O(1)
func (queue *Queue[T]) Empty() bool {
	return queue.size == 0
}

// Front returns the front element of the queue.
// Time complexity: O(1)
func (queue *Queue[T]) Front() T {
	if queue.size == 0 {
		panic(_emptyQueueErrorString)
	}
	return queue.head.value
}

// Back returns the last element of the queue.
// Time complexity: O(1)
func (queue *Queue[T]) Back() T {
	if queue.size == 0 {
		panic(_emptyQueueErrorString)
	}
	return queue.tail.value
}

// Size returns the size of the queue.
// Time complexity: O(1)
func (queue *Queue[T]) Size() int {
	return queue.size
}
