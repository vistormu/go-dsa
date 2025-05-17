package geometry

import (
	"math"

	"github.com/vistormu/go-dsa/constraints"
)

type Segment3D[T constraints.Number] struct {
	A, B Point3D[T]
}

func NewSegment3D[T constraints.Number](a, b Point3D[T]) Segment3D[T] {
	return Segment3D[T]{a, b}
}

func (s Segment3D[T]) Direction() Vector3D[T] {
	return s.B.Sub(s.A)
}

func (s Segment3D[T]) Length() float64 {
	return s.A.Distance(s.B)
}

func (s Segment3D[T]) Midpoint() Point3D[T] {
	return Point3D[T]{(s.A.X + s.B.X) / 2, (s.A.Y + s.B.Y) / 2, (s.A.Z + s.B.Z) / 2}
}

func (s Segment3D[T]) Distance(p Point3D[T]) float64 {
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
		cp := Point3D[float64]{
			float64(s.A.X) + float64(d.X)*t,
			float64(s.A.Y) + float64(d.Y)*t,
			float64(s.A.Z) + float64(d.Z)*t,
		}
		return p.Distance(Point3D[T]{T(cp.X), T(cp.Y), T(cp.Z)})
	}
}

func (s Segment3D[T]) DistanceSegment(o Segment3D[T]) float64 {
	d1 := s.Direction()
	d2 := o.Direction()
	r := s.A.Sub(o.A)

	a := float64(d1.Dot(d1))
	e := float64(d2.Dot(d2))
	f := float64(d2.Dot(r))

	if a == 0 && e == 0 {
		return s.A.Distance(o.A)
	}
	if a == 0 {
		return o.Distance(s.A)
	}
	if e == 0 {
		return s.Distance(o.A)
	}

	b := float64(d1.Dot(d2))
	c := float64(d1.Dot(r))
	den := a*e - b*b

	sn, sd := den, den
	tn, td := den, den

	if den != 0 {
		sn = b*f - c*e
	} else {
		sn = 0
		sd = 1
	}
	tn = b*sn/a + f

	if sn < 0 {
		sn = 0
		tn = f
		td = e
	} else if sn > sd {
		sn = sd
		tn = f + b
		td = e
	}

	if tn < 0 {
		tn = 0
		switch {
		case -c < 0:
			sn = 0
		case -c > a:
			sn = sd
		default:
			sn = -c
			sd = a
		}
	} else if tn > td {
		tn = td
		switch {
		case -c+b < 0:
			sn = 0
		case -c+b > a:
			sn = sd
		default:
			sn = -c + b
			sd = a
		}
	}

	sc := 0.0
	if sd != 0 {
		sc = sn / sd
	}
	tc := 0.0
	if td != 0 {
		tc = tn / td
	}

	cp1 := Point3D[float64]{
		float64(s.A.X) + float64(d1.X)*sc,
		float64(s.A.Y) + float64(d1.Y)*sc,
		float64(s.A.Z) + float64(d1.Z)*sc,
	}
	cp2 := Point3D[float64]{
		float64(o.A.X) + float64(d2.X)*tc,
		float64(o.A.Y) + float64(d2.Y)*tc,
		float64(o.A.Z) + float64(d2.Z)*tc,
	}

	dx := cp1.X - cp2.X
	dy := cp1.Y - cp2.Y
	dz := cp1.Z - cp2.Z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}
