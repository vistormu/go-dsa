package math

import (
	"github.com/vistormu/go-dsa/constraints"
)

func MapInterval[T constraints.Number](value, fromMin, fromMax, toMin, toMax T) T {
	if fromMin == fromMax {
		return toMin
	}

	inputRange := fromMax - fromMin
	outputRange := toMax - toMin

	return (value-fromMin)*outputRange/inputRange + toMin
}
