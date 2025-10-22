package control

import (
	"github.com/vistormu/go-dsa/constraints"
)

type Pid[T constraints.Float] struct {
	kp, ki, kd T
	alpha      T
	prevValue  T
	integral   T
	derivative T
}

func NewPid[T constraints.Float](kp, ki, kd, alpha T) *Pid[T] {
	return &Pid[T]{kp, ki, kd, alpha, 0, 0, 0}
}

func (p *Pid[T]) Compute(value T, dt T) T {
	p.integral += value * dt

	unfiltDerivative := (value - p.prevValue) / dt
	p.derivative = p.alpha*unfiltDerivative + (1-p.alpha)*p.derivative

	output := p.kp*value + p.ki*p.integral + p.kd*p.derivative

	p.prevValue = value

	return output
}
