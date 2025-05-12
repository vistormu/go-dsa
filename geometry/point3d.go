package geometry

import (
	"fmt"
	"math"

	"github.com/vistormu/go-dsa/constraints"
)

type Point3D[T constraints.Number] struct {
	X, Y, Z T
}

func NewPoint3D[T constraints.Number](x, y, z T) Point3D[T] {
	return Point3D[T]{x, y, z}
}

func (p Point3D[T]) Add(v Vector3D[T]) Point3D[T] {
	return Point3D[T]{p.X + v.X, p.Y + v.Y, p.Z + v.Z}
}

func (p Point3D[T]) Sub(q Point3D[T]) Vector3D[T] {
	return Vector3D[T]{p.X - q.X, p.Y - q.Y, p.Z - q.Z}
}

func (p Point3D[T]) Distance(q Point3D[T]) float64 {
	dx := float64(p.X - q.X)
	dy := float64(p.Y - q.Y)
	dz := float64(p.Z - q.Z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func (p Point3D[T]) String() string {
	return fmt.Sprintf("(%v, %v, %v)", p.X, p.Y, p.Z)
}

