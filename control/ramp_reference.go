package control

import (
	"github.com/vistormu/go-dsa/constraints"
)

type Ramp[T constraints.Number] struct {
	slope T
	delay T
}

func NewRampReference[T constraints.Number](slope, delay T) Ramp[T] {
	return Ramp[T]{slope, delay}
}

func (r Ramp[T]) Compute(t T) T {
	if t < r.delay {
		return 0
	}
	return r.slope * (t - r.delay)
}
