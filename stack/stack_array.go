package stack

import (
	"errors"
)

type StackArray[T any] struct {
	elements []T
}

func NewStackArray[T any]() *StackArray[T] {
	return &StackArray[T]{elements: []T{}}
}

func (s *StackArray[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

func (s *StackArray[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue, errors.New("stack is empty")
	}

	value := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]

	return value, nil
}

func (s *StackArray[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue, errors.New("stack is empty")
	}

	return s.elements[len(s.elements)-1], nil
}

func (s *StackArray[T]) Clear() {
	s.elements = []T{}
}

func (s *StackArray[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *StackArray[T]) Length() int {
	return len(s.elements)
}
