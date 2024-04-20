package stack

const _emptyStackErrorString = "can't perform operation, because stack is empty"

// Stack represents the stack data structure. It's implement with linked list.
type Stack[T any] struct {
	head *node[T]
	size int
}

type node[T any] struct {
	next  *node[T]
	value T
}

// New returns a new instance of Stack.
// Time complexity: O(1)
func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push adds an item to the top of the stack.
// Time complexity: O(1)
func (stack *Stack[T]) Push(item T) {
	newNode := &node[T]{
		value: item,
	}
	if stack.size == 0 {
		stack.head = newNode
		stack.size = 1
		return
	}

	newNode.next = stack.head
	stack.head = newNode
	stack.size++
}

// Pop removes an item from top of the stack.
// Time complexity: O(1)
func (stack *Stack[T]) Pop() {
	if stack.size == 0 {
		panic(_emptyStackErrorString)
	}
	stack.head = stack.head.next
	stack.size--
}

// Top returns the top element of the stack. return error, if stack is empty.
// Time complexity: O(1)
func (stack *Stack[T]) Top() T {
	if stack.size == 0 {
		panic(_emptyStackErrorString)
	}
	return stack.head.value
}

// Size returns the number of items in the stack.
// Time complexity: O(1)
func (stack *Stack[T]) Size() int {
	return stack.size
}

// Empty returns true if Stack is empty, else returns false.
// Time complexity: O(1)
func (stack *Stack[T]) Empty() bool {
	return stack.size == 0
}
