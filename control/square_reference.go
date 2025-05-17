package control

import (
	"github.com/vistormu/go-dsa/constraints"
	"math"
)

type Square[T constraints.Number] struct {
	amp    T
	freq   T
	phi    T
	offset T
}

func NewSquareReference[T constraints.Number](amp, freq, phi, offset T) Square[T] {
	return Square[T]{amp, freq, phi, offset}
}

func (s Square[T]) Compute(t T) T {
	return T(float64(s.amp)*math.Copysign(1, math.Sin(2*math.Pi*float64(s.freq*t+s.phi))) + float64(s.offset))
}
