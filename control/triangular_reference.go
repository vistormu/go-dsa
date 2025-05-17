package control

import (
	"github.com/vistormu/go-dsa/constraints"
	"math"
)

type Triangular[T constraints.Number] struct {
	amp    T
	freq   T
	phi    T
	offset T
}

func NewTriangularReference[T constraints.Number](amp, freq, phi, offset T) Triangular[T] {
	return Triangular[T]{amp, freq, phi, offset}
}

func (tr Triangular[T]) Compute(t T) T {
	return T(float64(tr.amp)*(1+math.Asin(math.Sin(2*math.Pi*float64(tr.freq*t+tr.phi)))*2/math.Pi) + float64(tr.offset))
}
