package control

import (
	"github.com/vistormu/go-dsa/constraints"
	"math"
)

type Triangular[T constraints.Number] struct {
	amp    float64
	freq   float64
	phi    float64
	offset float64
}

func NewTriangularReference[T constraints.Number](amp, freq, phi, offset T) Triangular[T] {
	return Triangular[T]{
		amp:    float64(amp),
		freq:   float64(freq),
		phi:    float64(phi),
		offset: float64(offset),
	}
}

func (tr Triangular[T]) Compute(t T) T {
	return T(tr.amp*(2/math.Pi*math.Asin(math.Sin(2*math.Pi*tr.freq*float64(t)+tr.phi))) + tr.offset)
}
