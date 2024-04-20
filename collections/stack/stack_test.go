package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		stack := New[int]()
		assert.Equal(t, 0, stack.size)
		assert.Nil(t, stack.head)
	})
}

func TestStack_Push(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		stack := New[int]()
		stack.Push(1)
		assert.Equal(t, 1, stack.head.value)
		assert.Nil(t, stack.head.next)

		stack.Push(2)
		assert.Equal(t, 2, stack.head.value)
	})
}

func Test_Pop(t *testing.T) {
	t.Run("success_when_called_on_empty_stack", func(t *testing.T) {
		stack := New[int]()
		stack.Push(1)
		stack.Push(2)
		stack.Pop()
		assert.Equal(t, 1, stack.head.value)
		stack.Pop()
		assert.Nil(t, stack.head)
	})

	t.Run("panic_when_called_on_empty_stack", func(t *testing.T) {
		stack := New[int]()
		defer func() {
			if r := recover(); r == nil {
				assert.Fail(t, "panic expected")
			} else {
				err, _ := r.(string)
				assert.Equal(t, "can't perform operation, because stack is empty", err)
			}
		}()
		stack.Pop()
	})
}

func Test_Top(t *testing.T) {
	t.Run("success_when_called_on_empty_stack", func(t *testing.T) {
		stack := New[int]()
		stack.Push(1)
		stack.Push(2)
		assert.Equal(t, 2, stack.Top())
		stack.Pop()
		assert.Equal(t, 1, stack.Top())
	})

	t.Run("panic_when_called_on_empty_stack", func(t *testing.T) {
		stack := New[int]()
		defer func() {
			if r := recover(); r == nil {
				assert.Fail(t, "panic expected")
			} else {
				err, _ := r.(string)
				assert.Equal(t, "can't perform operation, because stack is empty", err)
			}
		}()
		stack.Top()
	})
}

func Test_Size(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		stack := New[int]()
		stack.Push(1)
		assert.Equal(t, 1, stack.Size())
	})
}

func Test_Empty(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		stack := New[int]()
		assert.True(t, stack.Empty())
	})
	t.Run("non_empty_stack", func(t *testing.T) {
		stack := New[int]()
		stack.Push(1)
		assert.False(t, stack.Empty())
	})
}
