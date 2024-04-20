package deque

const _emptyDequeErrorString = "can't perform operation, because deque is empty"

// Deque represents the double ended queue. It's implemented using doubly linked list.
type Deque[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

type node[T any] struct {
	next  *node[T]
	prev  *node[T]
	value T
}

// New creates a new instance of Deque.
// Time complexity: O(1)
func New[T any]() *Deque[T] {
	return &Deque[T]{}
}

// PushBack pushes element in back of the deque.
// Time complexity: O(1)
func (deque *Deque[T]) PushBack(item T) {
	newNode := &node[T]{
		value: item,
	}

	if deque.size == 0 {
		deque.head = newNode
		deque.tail = newNode
		deque.size = 1
		return
	}

	deque.tail.next = newNode
	newNode.prev = deque.tail
	deque.tail = newNode
	deque.size++
}

// PopBack removes last element of the queue.
// Time complexity: O(1)
func (deque *Deque[T]) PopBack() {
	if deque.size == 0 {
		panic(_emptyDequeErrorString)
	}
	deque.size--
	deque.tail = deque.tail.prev
	if deque.tail != nil {
		deque.tail.next = nil
	}
}

// PushFront pushes element in front of the deque.
// Time complexity: O(1)
func (deque *Deque[T]) PushFront(item T) {
	newItem := &node[T]{
		value: item,
	}

	if deque.size == 0 {
		deque.head = newItem
		deque.tail = newItem
		deque.size = 1
		return
	}

	deque.head.prev = newItem
	newItem.next = deque.head
	deque.head = newItem
	deque.size++
}

// PopFront removes first element of the queue.
// Time complexity: O(1)
func (deque *Deque[T]) PopFront() {
	if deque.size == 0 {
		panic(_emptyDequeErrorString)
	}
	deque.size--
	deque.head = deque.head.next
	if deque.head != nil {
		deque.head.prev = nil
	}
}

// Back returns the last element of the deque.
// Time complexity: O(1)
func (deque *Deque[T]) Back() T {
	if deque.size == 0 {
		panic(_emptyDequeErrorString)
	}
	return deque.tail.value
}

// Front returns the first element of the deque.
// Time complexity: O(1)
func (deque *Deque[T]) Front() T {
	if deque.size == 0 {
		panic(_emptyDequeErrorString)
	}
	return deque.head.value
}

// Size returns the number of elements in the deque.
// Time complexity: O(1)
func (deque *Deque[T]) Size() int {
	return deque.size
}

// Empty returns true when deque is empty, else returns false.
// Time complexity: O(1)
func (deque *Deque[T]) Empty() bool {
	return deque.size == 0
}
