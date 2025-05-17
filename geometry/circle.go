package geometry

import (
	"math"

	"github.com/vistormu/go-dsa/constraints"
)

type Circle[T constraints.Number] struct {
	C Point2D[T]
	R T
}

func NewCircle[T constraints.Number](c Point2D[T], r T) Circle[T] {
	return Circle[T]{c, r}
}

func (c Circle[T]) Contains(p Point2D[T]) bool {
	dx := float64(p.X - c.C.X)
	dy := float64(p.Y - c.C.Y)
	return dx*dx+dy*dy <= float64(c.R*c.R)
}

func (c Circle[T]) Distance(p Point2D[T]) float64 {
	dx := float64(p.X - c.C.X)
	dy := float64(p.Y - c.C.Y)
	return math.Sqrt(dx*dx + dy*dy)
}

func (c Circle[T]) Intersects(o Circle[T]) bool {
	rSum := float64(c.R + o.R)
	dx := float64(o.C.X - c.C.X)
	dy := float64(o.C.Y - c.C.Y)
	return dx*dx+dy*dy <= rSum*rSum
}

