package main

import "fmt"

func main() {
	fmt.Println(powFloat64(2, 2.1))
}

// powInt returns x^n for any real x and any integer n.
func powInt(x float64, n int) float64 {
	var y float64
	if x != 0 {
		switch {
		case 0 < n:
			y = powInt(x, n/2)
			y *= y
			if n%2 != 0 {
				y *= x
			}
		case n < 0:
			y = 1 / powInt(x, -n)
		case n == 0:
			fallthrough
		default:
			y = 1
		}
	} else {
		if n == 0 {
			panic("indeterminant form")
		}
	}
	return y
}

// powFloat64 returns x^y for any real x and y.
//
func powFloat64(x, y float64) float64 {
	n := int(y)
	r := y - float64(n)
	m := int(1 / r)
	v := powInt(x, n)
	if r != 0 {
		switch {
		case 0 < x:
			v *= newton(func(z float64) float64 { return powInt(z, m) - x }, 1, 0.000001)
		case x < 0:
			if m%2 != 0 {
				v *= newton(func(z float64) float64 { return powInt(z, m) }, 1, 0.000001)
			} else {
				panic("indeterminant form")
			}
		case x == 0:
			fallthrough
		default:
			v = 0
		}
	}
	return v
}

func sqrt(x float64) float64 {
	return powFloat64(x, 0.5)
}

func nthRoot(x float64, n int) float64 {
	return powFloat64(x, 1/float64(n))
}

// newton finds a local root assuming f is smooth and continuous. Note, this algorithm is known to be unstable.
func newton(f func(x float64) float64, x0, tol float64) float64 {
	x1 := 2 * x0
	for tol < abs(x1-x0) {
		x0 = x1
		x1 = x0 - f(x0)/diff(f, x0, tol)
	}
	return x1
}

func diff(f func(x float64) float64, x0, tol float64) float64 {
	return (f(x0+tol) - f(x0-tol)) / (2 * tol)
}

func abs(x float64) float64 {
	if x < 0 {
		x *= -1
	}
	return x
}

// gcd returns the greatest common divisor of two non-negative integers.
// TODO: Optimize using least squares or something...
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
