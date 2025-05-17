package stack

import "errors"

type doublyNode[T comparable] struct {
	data T
	next *doublyNode[T]
	prev *doublyNode[T]
}

type UniqueStack[T comparable] struct {
	head   *doublyNode[T]
	tail   *doublyNode[T]
	length int
	index  map[T]*doublyNode[T]
}

func NewUniqueStack[T comparable]() *UniqueStack[T] {
	return &UniqueStack[T]{index: make(map[T]*doublyNode[T])}
}

func (s *UniqueStack[T]) IsEmpty() bool { return s.length == 0 }

func (s *UniqueStack[T]) Length() int { return s.length }

func (s *UniqueStack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("stack is empty")
	}
	return s.tail.data, nil
}

func (s *UniqueStack[T]) Clear() {
	s.head, s.tail = nil, nil
	s.length = 0
	s.index = make(map[T]*doublyNode[T])
}

func (s *UniqueStack[T]) Push(v T) {
	if n, ok := s.index[v]; ok {
		s.unlink(n)
	}

	newNode := &doublyNode[T]{data: v, prev: s.tail}
	if s.tail != nil {
		s.tail.next = newNode
	} else {
		s.head = newNode
	}
	s.tail = newNode

	s.index[v] = newNode
	s.length++
}

func (s *UniqueStack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("stack is empty")
	}
	n := s.tail
	s.unlink(n)
	delete(s.index, n.data)
	return n.data, nil
}

func (s *UniqueStack[T]) ToSlice() []T {
	slice := make([]T, 0, s.length)
	for n := s.head; n != nil; n = n.next {
		slice = append(slice, n.data)
	}
	return slice
}

func (s *UniqueStack[T]) unlink(n *doublyNode[T]) {
	if n.prev != nil {
		n.prev.next = n.next
	} else {
		s.head = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	} else {
		s.tail = n.prev
	}
	s.length--
}
