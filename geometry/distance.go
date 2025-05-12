package geometry

import (
	"github.com/vistormu/go-dsa/constraints"
	"math"
)

func Distance[T constraints.Number](a, b any) T {
	return 0
}

func point2dDistance[T constraints.Number](a, b Point2D[T]) T {
	return T(math.Sqrt(math.Pow(float64(b.X-a.X), 2) + math.Pow(float64(b.Y-a.Y), 2)))
}

func point3dDistance[T constraints.Number](a, b Point3D[T]) T {
	return T(math.Sqrt(math.Pow(float64(b.X-a.X), 2) + math.Pow(float64(b.Y-a.Y), 2) + math.Pow(float64(b.Z-a.Z), 2)))
}
