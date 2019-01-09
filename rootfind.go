package main

// Newton (also Newton-raphson) finds a local root assuming f is smooth
// and continuous. Note, Newton's method is known to be unstable.
func Newton(f func(x float64) float64, x0 float64) float64 {
	x1 := x0 + 2*tol
	for tol < Abs(x1-x0) {
		x0 = x1
		x1 = x0 - f(x0)/Diff(f, x0)
		if x0 == x1 {
			panic("x0 is a fixed point") // Fixed points map to themselves: f(x) = x
		}
	}
	return x1
}

// Bisection finds a local root assuming at least one root exists on the
// range [x0, x1]. Multiple roots may exist, but only one will be
// returned. Panics if no root exists. Bisection is more stable than
// Newton, but it is significantly less efficient.
func Bisection(f func(x float64) float64, x0, x1 float64) float64 {
	x := (x0 + x1) / 2.0
	y, y0, y1 := f(x), f(x0), f(x1)

	switch {
	case Abs(y) < tol:
		return x
	case y0*y < 0:
		return Bisection(f, x0, x)
	case y*y1 < 0:
		return Bisection(f, x, x1)
	default:
		panic("no root on boundary")
	}
}

// Diff returns the approximate value of df/dx at x0.
func Diff(f func(x float64) float64, x0 float64) float64 {
	return (f(x0+tol) - f(x0-tol)) / (2 * tol)
}
