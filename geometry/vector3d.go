package geometry

import (
	"fmt"
	"math"

	"github.com/vistormu/go-dsa/constraints"
)

type Vector3D[T constraints.Number] struct {
	X, Y, Z T
}

func NewVector3D[T constraints.Number](x, y, z T) Vector3D[T] {
	return Vector3D[T]{x, y, z}
}

func (v Vector3D[T]) Add(o Vector3D[T]) Vector3D[T] {
	return Vector3D[T]{v.X + o.X, v.Y + o.Y, v.Z + o.Z}
}

func (v Vector3D[T]) Sub(o Vector3D[T]) Vector3D[T] {
	return Vector3D[T]{v.X - o.X, v.Y - o.Y, v.Z - o.Z}
}

func (v Vector3D[T]) Mul(s T) Vector3D[T] {
	return Vector3D[T]{v.X * s, v.Y * s, v.Z * s}
}

func (v Vector3D[T]) Dot(o Vector3D[T]) T {
	return v.X*o.X + v.Y*o.Y + v.Z*o.Z
}

func (v Vector3D[T]) LenSq() T {
	return v.Dot(v)
}

func (v Vector3D[T]) Len() float64 {
	return math.Sqrt(float64(v.LenSq()))
}

func (v Vector3D[T]) Distance(o Vector3D[T]) float64 {
	return v.Sub(o).Len()
}

func (v Vector3D[T]) Norm() Vector3D[T] {
	l := v.Len()
	if l == 0 {
		return v
	}
	inv := 1 / l
	return Vector3D[T]{
		T(float64(v.X) * inv),
		T(float64(v.Y) * inv),
		T(float64(v.Z) * inv),
	}
}

func (v Vector3D[T]) String() string {
	return fmt.Sprintf("(%v, %v, %v)", v.X, v.Y, v.Z)
}

