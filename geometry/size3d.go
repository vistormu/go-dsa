package geometry

import (
	"fmt"

	"github.com/vistormu/go-dsa/constraints"
)

type Size3D[T constraints.Number] struct {
	W, H, D T
}

func NewSize3D[T constraints.Number](w, h, d T) Size3D[T] {
	return Size3D[T]{w, h, d}
}

func (s Size3D[T]) Add(o Size3D[T]) Size3D[T] {
	return Size3D[T]{s.W + o.W, s.H + o.H, s.D + o.D}
}

func (s Size3D[T]) Sub(o Size3D[T]) Size3D[T] {
	return Size3D[T]{s.W - o.W, s.H - o.H, s.D - o.D}
}

func (s Size3D[T]) MulScalar(k T) Size3D[T] {
	return Size3D[T]{s.W * k, s.H * k, s.D * k}
}

func (s Size3D[T]) Volume() T {
	return s.W * s.H * s.D
}

func (s Size3D[T]) IsZero() bool {
	return s.W == 0 && s.H == 0 && s.D == 0
}

func (s Size3D[T]) String() string {
	return fmt.Sprintf("(%v × %v × %v)", s.W, s.H, s.D)
}

