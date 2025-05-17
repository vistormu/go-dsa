package geometry

import (
	"math"

	"github.com/vistormu/go-dsa/constraints"
)

type Triangle2D[T constraints.Number] struct {
	A, B, C Point2D[T]
}

func NewTriangle2D[T constraints.Number](a, b, c Point2D[T]) Triangle2D[T] {
	return Triangle2D[T]{a, b, c}
}

func (t Triangle2D[T]) Area() float64 {
	return math.Abs(
		float64((t.B.X-t.A.X)*(t.C.Y-t.A.Y)-
			(t.C.X-t.A.X)*(t.B.Y-t.A.Y)),
	) / 2
}

func (t Triangle2D[T]) Perimeter() float64 {
	return t.A.Distance(t.B) +
		t.B.Distance(t.C) +
		t.C.Distance(t.A)
}

func sign[T constraints.Number](p1, p2, p3 Point2D[T]) float64 {
	return float64(p1.X-p3.X)*float64(p2.Y-p3.Y) -
		float64(p2.X-p3.X)*float64(p1.Y-p3.Y)
}

func (t Triangle2D[T]) ContainsPoint(p Point2D[T]) bool {
	b1 := sign(p, t.A, t.B) < 0
	b2 := sign(p, t.B, t.C) < 0
	b3 := sign(p, t.C, t.A) < 0
	return (b1 == b2) && (b2 == b3)
}
