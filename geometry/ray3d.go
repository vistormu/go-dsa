package geometry

import (
	"github.com/vistormu/go-dsa/constraints"
)

type Ray3D[T constraints.Number] struct {
	P Point3D[T]
	D Vector3D[T] // direction
}

func NewRay3D[T constraints.Number](p Point3D[T], dir Vector3D[T]) Ray3D[T] {
	return Ray3D[T]{p, dir}
}

func NewRay3DFromPoints[T constraints.Number](a, b Point3D[T]) Ray3D[T] {
	return Ray3D[T]{a, b.Sub(a)}
}

func (r Ray3D[T]) At(t T) Point3D[T] {
	return Point3D[T]{r.P.X + r.D.X*t, r.P.Y + r.D.Y*t, r.P.Z + r.D.Z*t}
}

func (r Ray3D[T]) Distance(q Point3D[T]) float64 {
	v := q.Sub(r.P)
	if dot := float64(v.Dot(r.D)); dot <= 0 {
		return q.Distance(r.P)
	}
	num := r.D.Cross(v).Len()
	den := r.D.Len()
	if den == 0 {
		return 0
	}
	return num / den
}
