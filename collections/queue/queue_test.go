package queue

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue_New(t *testing.T) {
	t.Run("queue_constructor_test", func(t *testing.T) {
		queue := New[int]()
		assert.Equal(t, 0, len(queue.items))
	})
}

func TestQueue_Push(t *testing.T) {
	t.Run("queue_push_test", func(t *testing.T) {
		queue := New[int]()
		queue.Push(1)
		queue.Push(2)
		assert.Equal(t, 2, len(queue.items))
		assert.Equal(t, 2, queue.items[len(queue.items)-1])
	})
}

func TestQueue_Pop(t *testing.T) {
	t.Run("queue_pop_test", func(t *testing.T) {
		queue := New[int]()
		queue.Push(1)
		queue.Push(2)
		element, err := queue.Pop()
		assert.Equal(t, 1, element)
		assert.Nil(t, err)

		element, err = queue.Pop()
		assert.Equal(t, 2, element)
		assert.Nil(t, err)

		element, err = queue.Pop()
		// zero value is returned
		assert.Equal(t, 0, element)
		assert.Equal(t, errors.New("queue is empty"), err)
	})
}

func TestQueue_Empty(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		queue := New[int]()
		assert.True(t, queue.Empty())
	})
}

func TestQueue_Size(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		queue := New[int]()
		assert.Equal(t, 0, queue.Size())
	})
}

func TestQueue_Front(t *testing.T) {
	t.Run("queue_front_test", func(t *testing.T) {
		queue := New[int]()
		queue.Push(1)
		queue.Push(2)
		element, err := queue.Front()
		assert.Equal(t, 1, element)
		assert.Nil(t, err)

		queue.Pop()
		queue.Pop()
		element, err = queue.Front()
		assert.Equal(t, 0, element)
		assert.Equal(t, errors.New("queue is empty"), err)
	})
}

func TestQueue_Back(t *testing.T) {
	t.Run("queue_back_test", func(t *testing.T) {
		queue := New[int]()
		queue.Push(1)
		queue.Push(2)
		element, err := queue.Back()
		assert.Equal(t, 2, element)
		assert.Nil(t, err)

		queue.Pop()
		queue.Pop()
		element, err = queue.Back()
		assert.Equal(t, 0, element)
		assert.Equal(t, errors.New("queue is empty"), err)
	})
}
