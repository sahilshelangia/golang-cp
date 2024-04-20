package deque

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeque_New(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		deque := New[int]()
		assert.Equal(t, 0, deque.size)
		assert.Nil(t, deque.head)
		assert.Nil(t, deque.tail)
	})
}

func TestDeque_PushBack(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		deque := New[int]()
		deque.PushBack(1)
		assert.Equal(t, 1, deque.head.value)
		assert.Equal(t, 1, deque.tail.value)
		assert.Nil(t, deque.head.prev)
		assert.Nil(t, deque.head.next)
		assert.Equal(t, 1, deque.Size())
		assert.Equal(t, deque.head, deque.tail)

		deque.PushBack(2)
		assert.Equal(t, 2, deque.tail.value)
		assert.Equal(t, 1, deque.head.value)
		assert.Nil(t, deque.head.prev)
		assert.Nil(t, deque.tail.next)
		assert.Equal(t, deque.head.next, deque.tail)
		assert.Equal(t, deque.tail.prev, deque.head)
		assert.Equal(t, 2, deque.Size())
	})
}

func TestDeque_PopBack(t *testing.T) {
	t.Run("success_when_called_on_non_empty_deque", func(t *testing.T) {
		deque := New[int]()
		deque.PushBack(1)
		deque.PushBack(2)

		deque.PopBack()
		assert.Equal(t, 1, deque.size)
		assert.Equal(t, 1, deque.head.value)
		assert.Equal(t, deque.head, deque.tail)
		assert.Nil(t, deque.head.prev)
		assert.Nil(t, deque.head.next)
	})
	t.Run("panic_when_called_on_empty_deque", func(t *testing.T) {
		deque := New[int]()
		defer func() {
			if r := recover(); r == nil {
				assert.Fail(t, "panic expected")
			} else {
				err, _ := r.(string)
				assert.Equal(t, "can't perform operation, because deque is empty", err)
			}
		}()
		deque.PopBack()
	})
}

func TestDeque_PushFront(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		deque := New[int]()
		deque.PushFront(1)
		assert.Equal(t, 1, deque.head.value)
		assert.Equal(t, 1, deque.tail.value)
		assert.Nil(t, deque.head.prev)
		assert.Nil(t, deque.head.next)
		assert.Equal(t, 1, deque.Size())
		assert.Equal(t, deque.head, deque.tail)

		deque.PushFront(2)
		assert.Equal(t, 2, deque.head.value)
		assert.Equal(t, 1, deque.tail.value)
		assert.Nil(t, deque.head.prev)
		assert.Nil(t, deque.tail.next)
		assert.Equal(t, deque.head.next, deque.tail)
		assert.Equal(t, deque.tail.prev, deque.head)
		assert.Equal(t, 2, deque.Size())
	})
}

func TestDeque_PopFront(t *testing.T) {
	t.Run("success_when_called_on_non_empty_deque", func(t *testing.T) {
		deque := New[int]()
		deque.PushBack(1)
		deque.PushBack(2)

		deque.PopFront()
		assert.Equal(t, 1, deque.size)
		assert.Equal(t, 2, deque.head.value)
		assert.Equal(t, deque.head, deque.tail)
		assert.Nil(t, deque.head.prev)
		assert.Nil(t, deque.head.next)
	})
	t.Run("panic_when_called_on_empty_deque", func(t *testing.T) {
		deque := New[int]()
		defer func() {
			if r := recover(); r == nil {
				assert.Fail(t, "panic expected")
			} else {
				err, _ := r.(string)
				assert.Equal(t, "can't perform operation, because deque is empty", err)
			}
		}()
		deque.PopFront()
	})
}

func TestDeque_Front(t *testing.T) {
	t.Run("success_when_called_on_non_empty_deque", func(t *testing.T) {
		deque := New[int]()
		deque.PushFront(1)
		deque.PushBack(2)
		assert.Equal(t, 1, deque.Front())
	})
	t.Run("panic_when_called_on_empty_deque", func(t *testing.T) {
		deque := New[int]()
		defer func() {
			if r := recover(); r == nil {
				assert.Fail(t, "panic expected")
			} else {
				err, _ := r.(string)
				assert.Equal(t, "can't perform operation, because deque is empty", err)
			}
		}()
		deque.Front()
	})
}

func TestDeque_Back(t *testing.T) {
	t.Run("success_when_called_on_non_empty_deque", func(t *testing.T) {
		deque := New[int]()
		deque.PushBack(1)
		deque.PushBack(2)
		assert.Equal(t, 2, deque.Back())
	})
	t.Run("panic_when_called_on_empty_deque", func(t *testing.T) {
		deque := New[int]()
		defer func() {
			if r := recover(); r == nil {
				assert.Fail(t, "panic expected")
			} else {
				err, _ := r.(string)
				assert.Equal(t, "can't perform operation, because deque is empty", err)
			}
		}()
		deque.Back()
	})
}

func TestDeque_Size(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		deque := New[int]()
		assert.Equal(t, 0, deque.Size())

		deque.PushBack(1)
		assert.Equal(t, 1, deque.Size())
	})
}

func TestDeque_Empty(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		deque := New[int]()
		assert.True(t, deque.Empty())

		deque.PushBack(1)
		assert.False(t, deque.Empty())
	})
}
