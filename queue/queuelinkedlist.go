package queue

import (
	"errors"
)

type QueueLinkedList[T any] struct {
	head   *node[T]
	tail   *node[T]
	length int
}

func NewQueueLinkedList[T any]() *QueueLinkedList[T] {
	return &QueueLinkedList[T]{}
}

func (q *QueueLinkedList[T]) Enqueue(data T) {
	// create a new node
	newNode := &node[T]{data: data}

	// update tail to point to the new node
	if q.tail != nil {
		q.tail.next = newNode
	}
	q.tail = newNode

	// if the queue is empty, set head to the new node
	if q.head == nil {
		q.head = newNode
	}

	// increment the length of the queue
	q.length++
}

func (q *QueueLinkedList[T]) Dequeue() (T, error) {
	// check if the queue is empty
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("queue is empty")
	}

	// store the data of the head node
	data := q.head.data

	// update head to point to the next node
	q.head = q.head.next

	// if the queue is now empty, set tail to nil
	if q.head == nil {
		q.tail = nil
	}

	// decrement the length of the queue
	q.length--

	return data, nil
}

func (q *QueueLinkedList[T]) Peek() (T, error) {
	// check if the queue is empty
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("queue is empty")
	}

	// return the data of the head node
	return q.head.data, nil
}

func (q *QueueLinkedList[T]) IsEmpty() bool {
	return q.length == 0
}

func (q *QueueLinkedList[T]) Length() int {
	return q.length
}

func (q *QueueLinkedList[T]) Clear() {
	// set head and tail to nil
	q.head = nil
	q.tail = nil

	// set length to 0
	q.length = 0
}
