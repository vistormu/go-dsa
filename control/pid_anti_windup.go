package control

import (
	"github.com/vistormu/go-dsa/constraints"
	"github.com/vistormu/go-dsa/math"
)

type PidAntiWindup[T constraints.Float] struct {
	kp, ki, kd     T
	dt             T
	alpha          T
	prevValue      T
	integral       T
	derivative     T
	integralBounds [2]T
}

func NewPidAntiWindup[T constraints.Float](kp, ki, kd, dt, alpha T, integralBounds [2]T) *PidAntiWindup[T] {
	return &PidAntiWindup[T]{kp, ki, kd, dt, alpha, 0, 0, 0, integralBounds}
}

func (p *PidAntiWindup[T]) Compute(value T) T {
	p.integral += value * p.dt
	p.integral = math.Clip(p.integral, p.integralBounds[0], p.integralBounds[1])

	unfiltDerivative := (value - p.prevValue) / p.dt
	p.derivative = p.alpha*unfiltDerivative + (1-p.alpha)*p.derivative

	output := p.kp*value + p.ki*p.integral + p.kd*p.derivative

	p.prevValue = value

	return output
}
