package filter

import (
	"github.com/vistormu/go-dsa/constraints"
	"slices"
)

type MedianFilter[T constraints.Number] struct {
	windowSize int
	values     []T
}

func NewMedianFilter[T constraints.Number](windowSize int) *MedianFilter[T] {
	return &MedianFilter[T]{
		windowSize: windowSize,
		values:     make([]T, 0, windowSize),
	}
}

func (mf *MedianFilter[T]) Compute(value T) T {
	mf.values = append(mf.values, value)

	if len(mf.values) > mf.windowSize {
		mf.values = mf.values[1:]
	}

	sortedValues := make([]T, len(mf.values))
	copy(sortedValues, mf.values)

	slices.Sort(sortedValues)

	n := len(sortedValues)
	if n%2 == 1 {
		return sortedValues[n/2]
	}
	return (sortedValues[n/2-1] + sortedValues[n/2]) / 2
}
