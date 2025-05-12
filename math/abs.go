package math

import (
	"github.com/vistormu/go-dsa/constraints"
)

func Abs[T constraints.Number](value T) T {
	if value < 0 {
		return -value
	}
	return value
}
