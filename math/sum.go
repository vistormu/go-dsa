package math

import (
	"github.com/vistormu/go-dsa/constraints"
)

func Sum[T constraints.Number](values []T) T {
	result := T(0)
	for _, value := range values {
		result += value
	}

	return result
}
