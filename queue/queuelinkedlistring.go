package queue

import (
	"errors"
)

type QueueLinkedListArray[T any] struct {
	head   *node[T]
	tail   *node[T]
	length int
}

func NewQueueLinkedListArray[T any]() *QueueLinkedListArray[T] {
	return &QueueLinkedListArray[T]{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (q *QueueLinkedListArray[T]) Enqueue(data T) {
	// create a new node
	newNode := &node[T]{data: data}

	// first element
	if q.head == nil {
		// if the queue is empty, set head and tail to the new node
		q.head = newNode
		q.tail = newNode
		newNode.next = newNode
	} else {
		// update tail to point to the new node
		newNode.next = q.head
		q.tail.next = newNode
		q.tail = newNode
	}

	// increment the length of the queue
	q.length++
}

func (q *QueueLinkedListArray[T]) Dequeue() (T, error) {
	// check if the queue is empty
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("queue is empty")
	}

	// store the data of the head node
	data := q.head.data

	// update head to point to the next node
	if q.head == q.tail {
		q.head = nil
		q.tail = nil
	} else {
		q.head = q.head.next
		q.tail.next = q.head
	}

	// decrement the length of the queue
	q.length--

	return data, nil
}

func (q *QueueLinkedListArray[T]) Peek() (T, error) {
	// check if the queue is empty
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("queue is empty")
	}

	// return the data of the head node
	return q.head.data, nil
}

func (q *QueueLinkedListArray[T]) ToSlice() []T {
	slice := make([]T, q.length)
	if q.head == nil {
		return slice
	}

	current := q.head
	for i := range q.length {
		slice[i] = current.data
		current = current.next
	}

	return slice
}

func (q *QueueLinkedListArray[T]) IsEmpty() bool {
	return q.length == 0
}

func (q *QueueLinkedListArray[T]) Length() int {
	return q.length
}

func (q *QueueLinkedListArray[T]) Clear() {
	// set head and tail to nil
	q.head = nil
	q.tail = nil

	// set length to 0
	q.length = 0
}
