package control

import (
	"github.com/vistormu/go-dsa/constraints"
)

type Pid[T constraints.Float] struct {
	kp, ki, kd T
	dt         T
	alpha      T
	prevValue  T
	integral   T
	derivative T
}

func NewPid[T constraints.Float](kp, ki, kd, dt, alpha T) *Pid[T] {
	return &Pid[T]{kp, ki, kd, dt, alpha, 0, 0, 0}
}

func (p *Pid[T]) Compute(value T) T {
	p.integral += value * p.dt

	unfiltDerivative := (value - p.prevValue) / p.dt
	p.derivative = p.alpha*unfiltDerivative + (1-p.alpha)*p.derivative

	output := p.kp*value + p.ki*p.integral + p.kd*p.derivative

	p.prevValue = value

	return output
}
