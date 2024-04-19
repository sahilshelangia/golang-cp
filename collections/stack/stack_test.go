package stack

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		stack := New[int]()
		assert.Equal(t, len(stack.items), 0)
	})
}

func TestStack_Push(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		stack := New[int]()
		stack.Push(1)
		assert.Equal(t, 1, stack.items[0])
	})
}

func Test_Pop(t *testing.T) {
	t.Run("return_element_when_stack_is_not_empty", func(t *testing.T) {
		stack := New[int]()
		stack.Push(1)
		element, err := stack.Pop()
		assert.Equal(t, 1, element)
		assert.Nil(t, err)
	})

	t.Run("return_error_when_stack_is_empty", func(t *testing.T) {
		stack := New[int]()
		element, err := stack.Pop()
		assert.Equal(t, 0, element)
		assert.Equal(t, errors.New("stack is empty"), err)
	})
}

func Test_Top(t *testing.T) {
	t.Run("return_element_when_stack_is_not_empty", func(t *testing.T) {
		stack := New[int]()
		stack.Push(1)
		element, err := stack.Top()
		assert.Equal(t, 1, element)
		assert.Nil(t, err)
	})

	t.Run("return_error_when_stack_is_empty", func(t *testing.T) {
		stack := New[int]()
		element, err := stack.Top()
		assert.Equal(t, 0, element)
		assert.Equal(t, errors.New("stack is empty"), err)
	})
}

func Test_Size(t *testing.T) {
	t.Run("size_equal_to_1", func(t *testing.T) {
		stack := New[int]()
		stack.Push(1)
		assert.Equal(t, 1, stack.Size())
	})
}

func Test_Empty(t *testing.T) {
	t.Run("empty_stack", func(t *testing.T) {
		stack := New[int]()
		assert.True(t, stack.Empty())
	})
	t.Run("non_empty_stack", func(t *testing.T) {
		stack := New[int]()
		stack.Push(1)
		assert.False(t, stack.Empty())
	})
}
