package geometry

import (
	"math"

	"github.com/vistormu/go-dsa/constraints"
)

type Line2D[T constraints.Number] struct {
	P Point2D[T]
	D Vector2D[T]
}

func NewLine2D[T constraints.Number](p Point2D[T], d Vector2D[T]) Line2D[T] {
	return Line2D[T]{p, d}
}

func NewLine2DFromPoints[T constraints.Number](a, b Point2D[T]) Line2D[T] {
	return Line2D[T]{a, b.Sub(a)}
}

func (l Line2D[T]) At(t T) Point2D[T] {
	return Point2D[T]{l.P.X + l.D.X*t, l.P.Y + l.D.Y*t}
}

func (l Line2D[T]) Distance(q Point2D[T]) float64 {
	vx := float64(q.X - l.P.X)
	vy := float64(q.Y - l.P.Y)
	dx := float64(l.D.X)
	dy := float64(l.D.Y)

	num := math.Abs(vx*dy - vy*dx)
	den := math.Sqrt(dx*dx + dy*dy)
	if den == 0 {
		return 0
	}
	return num / den
}

func (l Line2D[T]) Intersection(o Line2D[T]) (Point2D[float64], bool) {
	dx1 := float64(l.D.X)
	dy1 := float64(l.D.Y)
	dx2 := float64(o.D.X)
	dy2 := float64(o.D.Y)

	det := dx1*dy2 - dy1*dx2
	if math.Abs(det) < 1e-12 {
		return Point2D[float64]{}, false
	}

	px := float64(o.P.X - l.P.X)
	py := float64(o.P.Y - l.P.Y)
	t := (px*dy2 - py*dx2) / det

	return Point2D[float64]{
		float64(l.P.X) + dx1*t,
		float64(l.P.Y) + dy1*t,
	}, true
}
