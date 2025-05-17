package control

import (
	"github.com/vistormu/go-dsa/constraints"
)

type Step[T constraints.Number] struct {
	amp   T
	delay T
}

func NewStepReference[T constraints.Number](amp, delay T) Step[T] {
	return Step[T]{amp, delay}
}

func (s Step[T]) Compute(t T) T {
	if t < s.delay {
		return 0
	}
	return s.amp
}
