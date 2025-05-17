package filter

import (
	"github.com/vistormu/go-dsa/constraints"
)

type KalmanFilter[T constraints.Number] struct {
	q    float64
	r    float64
	xHat float64
	p    float64
	f    float64
	h    float64
}

func NewKalmanFilter[T constraints.Number](processVariance, measurementVariance, initialErrorCovariance float64, initialEstimate T) *KalmanFilter[T] {
	return &KalmanFilter[T]{
		q:    processVariance,
		r:    measurementVariance,
		p:    initialErrorCovariance,
		f:    1.0,
		h:    1.0,
		xHat: float64(initialEstimate),
	}
}

func (kf *KalmanFilter[T]) Compute(measurement T) T {
	xHatPredict := kf.f * kf.xHat
	pPredict := kf.f*kf.p*kf.f + kf.q

	k := pPredict * kf.h / (kf.h*pPredict*kf.h + kf.r)
	kf.xHat = xHatPredict + k*(float64(measurement)-kf.h*xHatPredict)
	kf.p = (1 - k*kf.h) * pPredict

	return T(kf.xHat)
}
