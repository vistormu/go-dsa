package math

import (
	"github.com/vistormu/go-dsa/constraints"
)

func Mean[T constraints.Number](values []T) T {
	var sum T
	length := len(values)

	for _, v := range values {
		sum += v
	}

	return sum / T(length)
}
