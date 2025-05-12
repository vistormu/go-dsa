package geometry

import (
	"fmt"
	"math"

	"github.com/vistormu/go-dsa/constraints"
)

type Vector2D[T constraints.Number] struct {
	X, Y T
}

func NewVector2D[T constraints.Number](x, y T) Vector2D[T] {
	return Vector2D[T]{x, y}
}

func (v Vector2D[T]) Add(o Vector2D[T]) Vector2D[T] {
	return Vector2D[T]{v.X + o.X, v.Y + o.Y}
}

func (v Vector2D[T]) Sub(o Vector2D[T]) Vector2D[T] {
	return Vector2D[T]{v.X - o.X, v.Y - o.Y}
}

func (v Vector2D[T]) Mul(s T) Vector2D[T] {
	return Vector2D[T]{v.X * s, v.Y * s}
}

func (v Vector2D[T]) Dot(o Vector2D[T]) T {
	return v.X*o.X + v.Y*o.Y
}

func (v Vector2D[T]) LenSq() T {
	return v.Dot(v)
}

func (v Vector2D[T]) Len() float64 {
	return math.Sqrt(float64(v.LenSq()))
}

func (v Vector2D[T]) Distance(o Vector2D[T]) float64 {
	return v.Sub(o).Len()
}

func (v Vector2D[T]) Norm() Vector2D[T] {
	l := v.Len()
	if l == 0 {
		return v
	}
	inv := 1 / l
	return Vector2D[T]{T(float64(v.X) * inv), T(float64(v.Y) * inv)}
}

func (v Vector2D[T]) String() string {
	return fmt.Sprintf("(%v, %v)", v.X, v.Y)
}

