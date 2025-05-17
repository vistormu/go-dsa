package geometry

import (
	"github.com/vistormu/go-dsa/constraints"
)

type Line3D[T constraints.Number] struct {
	P Point3D[T]
	D Vector3D[T]
}

func NewLine3D[T constraints.Number](p Point3D[T], d Vector3D[T]) Line3D[T] {
	return Line3D[T]{p, d}
}

func NewLine3DFromPoints[T constraints.Number](a, b Point3D[T]) Line3D[T] {
	return Line3D[T]{a, b.Sub(a)}
}

func (l Line3D[T]) At(t T) Point3D[T] {
	return Point3D[T]{
		l.P.X + l.D.X*t,
		l.P.Y + l.D.Y*t,
		l.P.Z + l.D.Z*t,
	}
}

func (a Vector3D[T]) Cross(b Vector3D[T]) Vector3D[T] {
	return Vector3D[T]{
		a.Y*b.Z - a.Z*b.Y,
		a.Z*b.X - a.X*b.Z,
		a.X*b.Y - a.Y*b.X,
	}
}

func (l Line3D[T]) Distance(q Point3D[T]) float64 {
	v := q.Sub(l.P)
	num := l.D.Cross(v).Len()
	den := l.D.Len()
	if den == 0 {
		return 0
	}
	return num / den
}

func (l Line3D[T]) Len() float64 { return l.D.Len() }

