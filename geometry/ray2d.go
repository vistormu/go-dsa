package geometry

import (
	"fmt"
	"math"

	"github.com/vistormu/go-dsa/constraints"
)

type Ray2D[T constraints.Number] struct {
	P Point2D[T]  // origin
	D Vector2D[T] // direction (need not be unit, but must be non-zero)
}

func NewRay2D[T constraints.Number](p Point2D[T], dir Vector2D[T]) Ray2D[T] {
	return Ray2D[T]{p, dir}
}

func NewRay2DFromPoints[T constraints.Number](a, b Point2D[T]) Ray2D[T] {
	return Ray2D[T]{a, b.Sub(a)}
}

func (r Ray2D[T]) At(t T) Point2D[T] {
	return Point2D[T]{r.P.X + r.D.X*t, r.P.Y + r.D.Y*t}
}

// Distance from point to ray (0 when point lies on the ray)
func (r Ray2D[T]) Distance(q Point2D[T]) float64 {
	vx := float64(q.X - r.P.X)
	vy := float64(q.Y - r.P.Y)
	dx := float64(r.D.X)
	dy := float64(r.D.Y)

	dot := vx*dx + vy*dy
	if dot <= 0 { // behind origin
		return math.Sqrt(vx*vx + vy*vy)
	}

	cross := math.Abs(vx*dy - vy*dx)
	return cross / math.Sqrt(dx*dx+dy*dy)
}

// Intersection of two rays; returns point and true when they meet with t,u ≥ 0
func (r Ray2D[T]) Intersect(o Ray2D[T]) (Point2D[float64], bool) {
	dx1, dy1 := float64(r.D.X), float64(r.D.Y)
	dx2, dy2 := float64(o.D.X), float64(o.D.Y)

	det := dx1*dy2 - dy1*dx2
	if math.Abs(det) < 1e-12 {
		return Point2D[float64]{}, false
	}

	px := float64(o.P.X - r.P.X)
	py := float64(o.P.Y - r.P.Y)

	t := (px*dy2 - py*dx2) / det
	u := (px*dy1 - py*dx1) / det
	if t < 0 || u < 0 {
		return Point2D[float64]{}, false
	}

	return Point2D[float64]{float64(r.P.X) + dx1*t, float64(r.P.Y) + dy1*t}, true
}

func (r Ray2D[T]) String() string {
	return fmt.Sprintf("Ray{P:%v D:%v}", r.P, r.D)
}
