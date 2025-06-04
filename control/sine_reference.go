package control

import (
	"github.com/vistormu/go-dsa/constraints"
	"math"
)

type Sine[T constraints.Number] struct {
	amp    float64
	freq   float64
	phi    float64
	offset float64
}

func NewSineReference[T constraints.Number](amp, freq, phi, offset T) Sine[T] {
	return Sine[T]{
		amp:    float64(amp),
		freq:   float64(freq),
		phi:    float64(phi),
		offset: float64(offset),
	}
}

func (s Sine[T]) Compute(t T) T {
	return T(s.amp*math.Sin(2*math.Pi*s.freq*float64(t)+s.phi) + s.offset)
}
