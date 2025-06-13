package buffer

import (
	"fmt"
	"slices"
)

type BoundedBuffer[T any] struct {
	capacity int
	buffer   []T
}

func NewBoundedBuffer[T any](capacity int) *BoundedBuffer[T] {
	if capacity <= 0 {
		panic("capacity must be greater than zero")
	}
	return &BoundedBuffer[T]{
		capacity: capacity,
		buffer:   make([]T, 0, capacity),
	}
}

func (bb *BoundedBuffer[T]) Full() bool {
	return len(bb.buffer) >= bb.capacity
}

func (bb *BoundedBuffer[T]) Empty() bool {
	return len(bb.buffer) == 0
}

func (bb *BoundedBuffer[T]) Add(item T) error {
	if bb.Full() {
		return fmt.Errorf("buffer is full, cannot add item: %v", item)
	}

	bb.buffer = append(bb.buffer, item)

	return nil
}

func (bb *BoundedBuffer[T]) Remove() (T, error) {
	if bb.Empty() {
		var zero T
		return zero, fmt.Errorf("buffer is empty, cannot remove item")
	}

	item := bb.buffer[0]
	bb.buffer = bb.buffer[1:]

	return item, nil
}

func (bb *BoundedBuffer[T]) Size() int {
	return len(bb.buffer)
}

func (bb *BoundedBuffer[T]) Capacity() int {
	return bb.capacity
}

func (bb *BoundedBuffer[T]) Clear() {
	bb.buffer = make([]T, 0, bb.capacity)
}

func (bb *BoundedBuffer[T]) ToSlice() []T {
	return slices.Clone(bb.buffer) // Return a copy of the buffer
}
