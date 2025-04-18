package timing

import "math"

type EasingFunc func(t float64) float64

func Linear(t float64) float64 {
	return t
}

func EaseInOutQuad(t float64) float64 {
	if t < 0.5 {
		return 2 * t * t
	}
	return -1 + (4-2*t)*t
}

func CubicBezier(p1, p2 float64) EasingFunc {
	return func(t float64) float64 {

		u := 1 - t
		return 3*u*u*t*p1 + 3*u*t*t*p2 + t*t*t
	}
}

func Spring(t float64) float64 {
	return 1 - math.Exp(-6*t)*math.Cos(12*t)
}
