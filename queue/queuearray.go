package queue

import (
	"errors"
)

type QueueArray[T any] struct {
	data   []T
	length int
}

func NewQueueArray[T any]() *QueueArray[T] {
	return &QueueArray[T]{}
}

func (q *QueueArray[T]) Enqueue(data T) {
	// append the data to the end of the queue
	q.data = append(q.data, data)

	// increment the length of the queue
	q.length++
}

func (q *QueueArray[T]) Dequeue() (T, error) {
	// check if the queue is empty
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("queue is empty")
	}

	// store the data of the first element
	data := q.data[0]

	// remove the first element from the queue
	q.data = q.data[1:]

	// decrement the length of the queue
	q.length--

	return data, nil
}

func (q *QueueArray[T]) ToSlice() []T {
	slice := make([]T, q.length)
	copy(slice, q.data)

	return slice
}

func (q *QueueArray[T]) IsEmpty() bool {
	return q.length == 0
}

func (q *QueueArray[T]) Length() int {
	return q.length
}

func (q *QueueArray[T]) Peek() (T, bool) {
	// check if the queue is empty
	if q.IsEmpty() {
		var zero T
		return zero, false
	}

	// return the first element of the queue
	return q.data[0], true
}

func (q *QueueArray[T]) Clear() {
	// clear the data slice
	q.data = []T{}

	// reset the length of the queue
	q.length = 0
}
