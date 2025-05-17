package control

import (
	"github.com/vistormu/go-dsa/constraints"
	"math"
)

type Sine[T constraints.Number] struct {
	amp    T
	freq   T
	phi    T
	offset T
}

func NewSineReference[T constraints.Number](amp, freq, phi, offset T) Sine[T] {
	return Sine[T]{amp, freq, phi, offset}
}

func (s Sine[T]) Compute(t T) T {
	return T(float64(s.amp)*math.Sin(2*math.Pi*float64(s.freq*t+s.phi)) + float64(s.offset))
}
