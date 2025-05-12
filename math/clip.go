package math

import (
	"github.com/vistormu/go-dsa/constraints"
)

func Clip[T constraints.Number](value, min, max T) T {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
