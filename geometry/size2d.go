package geometry

import (
	"fmt"

	"github.com/vistormu/go-dsa/constraints"
)

type Size2D[T constraints.Number] struct {
	W, H T
}

func NewSize2D[T constraints.Number](w, h T) Size2D[T] {
	return Size2D[T]{w, h}
}

func (s Size2D[T]) Add(o Size2D[T]) Size2D[T] {
	return Size2D[T]{s.W + o.W, s.H + o.H}
}

func (s Size2D[T]) Sub(o Size2D[T]) Size2D[T] {
	return Size2D[T]{s.W - o.W, s.H - o.H}
}

func (s Size2D[T]) Mul(k T) Size2D[T] {
	return Size2D[T]{s.W * k, s.H * k}
}

func (s Size2D[T]) Area() T {
	return s.W * s.H
}

func (s Size2D[T]) IsZero() bool {
	return s.W == 0 && s.H == 0
}

func (s Size2D[T]) String() string {
	return fmt.Sprintf("(%v × %v)", s.W, s.H)
}

