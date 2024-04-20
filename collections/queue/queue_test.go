package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue_New(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		queue := New[int]()
		assert.Equal(t, 0, queue.size)
		assert.Nil(t, queue.head)
		assert.Nil(t, queue.tail)
	})
}

func TestQueue_Push(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		queue := New[int]()
		queue.Push(1)
		assert.Equal(t, 1, queue.size)
		assert.Equal(t, 1, queue.head.value)
		assert.Nil(t, queue.head.next)

		queue.Push(2)
		assert.Equal(t, 2, queue.size)
		assert.Equal(t, 2, queue.tail.value)
		assert.Equal(t, queue.head.next, queue.tail)
		assert.Nil(t, queue.tail.next)
	})
}

func TestQueue_Pop(t *testing.T) {
	t.Run("success_when_queue_have_elements", func(t *testing.T) {
		queue := New[int]()
		queue.Push(1)
		queue.Push(2)
		queue.Pop()
		assert.Equal(t, 1, queue.size)
		assert.Equal(t, 2, queue.head.value)
		assert.Equal(t, queue.head, queue.tail)
		assert.Nil(t, queue.head.next)
	})
	t.Run("panic_when_queue_pop_called_on_empty_queue", func(t *testing.T) {
		queue := New[int]()
		defer func() {
			if r := recover(); r == nil {
				assert.Fail(t, "panic expected")
			} else {
				err, _ := r.(string)
				assert.Equal(t, "can't perform operation, because queue is empty", err)
			}
		}()
		queue.Pop()
	})
}

func TestQueue_Front(t *testing.T) {
	t.Run("success_when_called_on_non_empty_queue", func(t *testing.T) {
		queue := New[int]()
		queue.Push(1)
		queue.Push(2)
		assert.Equal(t, 1, queue.Front())

		queue.Pop()
		assert.Equal(t, 2, queue.Front())
	})
	t.Run("panic_when_called_on_empty_queue", func(t *testing.T) {
		queue := New[int]()
		defer func() {
			if r := recover(); r == nil {
				assert.Fail(t, "panic expected")
			} else {
				err, _ := r.(string)
				assert.Equal(t, "can't perform operation, because queue is empty", err)
			}
		}()
		queue.Front()
	})
}

func TestQueue_Back(t *testing.T) {
	t.Run("success_when_called_on_non_empty_queue", func(t *testing.T) {
		queue := New[int]()
		queue.Push(1)
		queue.Push(2)
		assert.Equal(t, 2, queue.Back())

		queue.Pop()
		assert.Equal(t, 2, queue.Back())
	})
	t.Run("panic_when_called_on_empty_queue", func(t *testing.T) {
		queue := New[int]()
		defer func() {
			if r := recover(); r == nil {
				assert.Fail(t, "panic expected")
			} else {
				err, _ := r.(string)
				assert.Equal(t, "can't perform operation, because queue is empty", err)
			}
		}()
		queue.Back()
	})
}

func TestQueue_Empty(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		queue := New[int]()
		assert.True(t, queue.Empty())

		queue.Push(1)
		assert.False(t, queue.Empty())
	})
}

func TestQueue_Size(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		queue := New[int]()
		assert.Equal(t, 0, queue.Size())

		queue.Push(1)
		assert.Equal(t, 1, queue.Size())
	})
}
