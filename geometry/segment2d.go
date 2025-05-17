package geometry

import (
	"math"

	"github.com/vistormu/go-dsa/constraints"
)

type Segment2D[T constraints.Number] struct {
	A, B Point2D[T]
}

func NewSegment2D[T constraints.Number](a, b Point2D[T]) Segment2D[T] {
	return Segment2D[T]{a, b}
}

func (s Segment2D[T]) Direction() Vector2D[T] {
	return s.B.Sub(s.A)
}

func (s Segment2D[T]) Length() float64 {
	return s.A.Distance(s.B)
}

func (s Segment2D[T]) Midpoint() Point2D[T] {
	return Point2D[T]{(s.A.X + s.B.X) / 2, (s.A.Y + s.B.Y) / 2}
}

func (s Segment2D[T]) Distance(p Point2D[T]) float64 {
	d := s.Direction()
	lenSq := float64(d.LenSq())
	if lenSq == 0 {
		return p.Distance(s.A)
	}
	t := float64(p.Sub(s.A).Dot(d)) / lenSq
	switch {
	case t <= 0:
		return p.Distance(s.A)
	case t >= 1:
		return p.Distance(s.B)
	default:
		proj := Point2D[float64]{
			float64(s.A.X) + float64(d.X)*t,
			float64(s.A.Y) + float64(d.Y)*t,
		}
		return p.Distance(Point2D[T]{T(proj.X), T(proj.Y)})
	}
}

func (s Segment2D[T]) Intersects(o Segment2D[T]) (Point2D[float64], bool) {
	rx, ry := float64(s.B.X-s.A.X), float64(s.B.Y-s.A.Y)
	sx, sy := float64(o.B.X-o.A.X), float64(o.B.Y-o.A.Y)
	det := rx*sy - ry*sx
	if math.Abs(det) < 1e-12 {
		return Point2D[float64]{}, false
	}
	qpx, qpy := float64(o.A.X-s.A.X), float64(o.A.Y-s.A.Y)
	t := (qpx*sy - qpy*sx) / det
	u := (qpx*ry - qpy*rx) / det
	if t < 0 || t > 1 || u < 0 || u > 1 {
		return Point2D[float64]{}, false
	}
	return Point2D[float64]{float64(s.A.X) + rx*t, float64(s.A.Y) + ry*t}, true
}
