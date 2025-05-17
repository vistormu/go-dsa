package stack

import (
	"errors"
)

type node[T any] struct {
	data T
	next *node[T]
}

type StackLinkedList[T any] struct {
	head   *node[T]
	length int
}

func NewStackLinkedList[T any]() *StackLinkedList[T] {
	return &StackLinkedList[T]{}
}

func (s *StackLinkedList[T]) Push(data T) {
	// create a new node
	newNode := &node[T]{data: data}

	// set the new node's next to the current head
	newNode.next = s.head

	// update head to point to the new node
	s.head = newNode

	// increment the length of the stack
	s.length++
}

func (s *StackLinkedList[T]) Pop() (T, error) {
	// check if the stack is empty
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("stack is empty")
	}

	// store the data of the head node
	data := s.head.data

	// update head to point to the next node
	s.head = s.head.next

	// decrement the length of the stack
	s.length--

	return data, nil
}

func (s *StackLinkedList[T]) Peek() (T, error) {
	// check if the stack is empty
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("stack is empty")
	}

	return s.head.data, nil
}

func (s *StackLinkedList[T]) IsEmpty() bool {
	return s.length == 0
}

func (s *StackLinkedList[T]) Length() int {
	return s.length
}

func (s *StackLinkedList[T]) Clear() {
	s.head = nil
	s.length = 0
}

func (s *StackLinkedList[T]) ToSlice() []T {
	slice := make([]T, 0, s.length)
	current := s.head
	for current != nil {
		slice = append(slice, current.data)
		current = current.next
	}
	return slice
}
