package geometry

import "github.com/vistormu/go-dsa/constraints"

type Rect[T constraints.Number] struct {
	Min, Max Point2D[T] // inclusive lower-left, exclusive upper-right
}

func NewRect[T constraints.Number](x0, y0, x1, y1 T) Rect[T] {
	if x0 > x1 {
		x0, x1 = x1, x0
	}
	if y0 > y1 {
		y0, y1 = y1, y0
	}
	return Rect[T]{Point2D[T]{x0, y0}, Point2D[T]{x1, y1}}
}

func (r Rect[T]) Width() T {
	return r.Max.X - r.Min.X
}

func (r Rect[T]) Height() T {
	return r.Max.Y - r.Min.Y
}

func (r Rect[T]) Area() T {
	return r.Width() * r.Height()
}

func (r Rect[T]) ContainsPoint(p Point2D[T]) bool {
	return p.X >= r.Min.X && p.X < r.Max.X &&
		p.Y >= r.Min.Y && p.Y < r.Max.Y
}

func (r Rect[T]) Intersects(o Rect[T]) bool {
	return r.Min.X < o.Max.X && r.Max.X > o.Min.X &&
		r.Min.Y < o.Max.Y && r.Max.Y > o.Min.Y
}

func (r Rect[T]) Intersection(o Rect[T]) (Rect[T], bool) {
	if !r.Intersects(o) {
		return Rect[T]{}, false
	}
	return Rect[T]{
		Min: Point2D[T]{max(r.Min.X, o.Min.X), max(r.Min.Y, o.Min.Y)},
		Max: Point2D[T]{min(r.Max.X, o.Max.X), min(r.Max.Y, o.Max.Y)},
	}, true
}

func (r Rect[T]) Union(o Rect[T]) Rect[T] {
	return Rect[T]{
		Min: Point2D[T]{min(r.Min.X, o.Min.X), min(r.Min.Y, o.Min.Y)},
		Max: Point2D[T]{max(r.Max.X, o.Max.X), max(r.Max.Y, o.Max.Y)},
	}
}
