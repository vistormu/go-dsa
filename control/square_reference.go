package control

import (
	"github.com/vistormu/go-dsa/constraints"
	"math"
)

type Square[T constraints.Number] struct {
	amp    float64
	freq   float64
	phi    float64
	offset float64
}

func NewSquareReference[T constraints.Number](amp, freq, phi, offset T) Square[T] {
	return Square[T]{
		amp:    float64(amp),
		freq:   float64(freq),
		phi:    float64(phi),
		offset: float64(offset),
	}
}

func (s Square[T]) Compute(t T) T {
	// Generate a square wave using the sign of the sine wave
	return T(s.amp*math.Copysign(1, math.Sin(2*math.Pi*s.freq*float64(t)+s.phi)) + s.offset)
}

