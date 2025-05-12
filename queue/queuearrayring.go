package queue

import (
	"errors"
)

type QueueArrayRing[T any] struct {
	data     []T
	head     int
	tail     int
	size     int
	capacity int
}

func NewQueueArrayRing[T any](capacity int) *QueueArrayRing[T] {
	if capacity <= 0 {
		panic("capacity must be greater than 0")
	}

	return &QueueArrayRing[T]{
		data:     make([]T, capacity),
		head:     0,
		tail:     0,
		size:     0,
		capacity: capacity,
	}
}

func (q *QueueArrayRing[T]) Enqueue(value T) error {
	if q.size == q.capacity {
		return errors.New("queue is full")
	}

	q.data[q.tail] = value
	q.tail = (q.tail + 1) % q.capacity
	q.size++

	return nil
}

func (q *QueueArrayRing[T]) Dequeue() (T, error) {
	if q.size == 0 {
		var zeroValue T
		return zeroValue, errors.New("queue is empty")
	}

	value := q.data[q.head]
	q.head = (q.head + 1) % q.capacity
	q.size--

	return value, nil
}

func (q *QueueArrayRing[T]) Peek() (T, error) {
	if q.size == 0 {
		var zeroValue T
		return zeroValue, errors.New("queue is empty")
	}

	return q.data[q.head], nil
}

func (q *QueueArrayRing[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *QueueArrayRing[T]) IsFull() bool {
	return q.size == q.capacity
}

func (q *QueueArrayRing[T]) Size() int {
	return q.size
}

func (q *QueueArrayRing[T]) Capacity() int {
	return q.capacity
}
