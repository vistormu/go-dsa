package geometry

import (
	"fmt"
	"math"

	"github.com/vistormu/go-dsa/constraints"
)

type Point2D[T constraints.Number] struct {
	X, Y T
}

func NewPoint2D[T constraints.Number](x, y T) Point2D[T] {
	return Point2D[T]{x, y}
}

func (p Point2D[T]) Add(v Vector2D[T]) Point2D[T] {
	return Point2D[T]{p.X + v.X, p.Y + v.Y}
}

func (p Point2D[T]) Sub(q Point2D[T]) Vector2D[T] {
	return Vector2D[T]{p.X - q.X, p.Y - q.Y}
}

func (p Point2D[T]) Distance(q Point2D[T]) float64 {
	return math.Sqrt(float64(p.Sub(q).LenSq()))
}

func (p Point2D[T]) String() string {
	return fmt.Sprintf("(%v, %v)", p.X, p.Y)
}

