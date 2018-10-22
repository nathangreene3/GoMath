package main

import (
	"fmt"
	"math"
)

const (
	// E is Euler's number
	E float64 = 2.718281
	// Pi is the ratio of a circle's circumfrence to its diameter
	Pi float64 = 3.141592
)

func main() {
	// x := -2.0
	// r := 1.0 / 3.0
	// n := int(1 / r)
	// fmt.Println(powFloat64(x, r))
	// fmt.Println(nthRoot(x, n))
	// fmt.Println(math.Pow(x, r))
	var (
		bestPi float64
		pi     float64
		err    float64
		k      int
	)
	for i := 0; i < 20; i++ {
		pi = archimedes(i)
		fmt.Println(pi)
		if abs(math.Pi-pi) < abs(math.Pi-bestPi) {
			k = i
			bestPi = pi
		}
	}
	fmt.Println(k, bestPi, math.Pi, err)
}

// archimedes approximates pi. k = 10 is optimal, but still terrible. Math pkg gives k = 13 as optimal.
func archimedes(k int) float64 {
	b := sqrt(3.0) / 2.0
	for i := 0; i < k; i++ {
		b = sqrt((1.0 - sqrt(1.0-powFloat64(b, 2.0))) / 2.0)
	}
	return 3.0 * powInt(2.0, k) * b
}

// round returns x rounded to the nearest whole number as a float64.
func round(x float64) float64 {
	y := float64(int(x)) // Floor of x
	if 0.5 <= x-y {
		y++
	}
	return y
}

func exp(x float64) float64 {
	return powFloat64(E, x)
}

// ln returns the natural logarithm of a number
func ln(x float64) float64 {
	if x <= 0 {
		panic("x must be positive")
	}
	v0 := float64(0)
	v1 := float64(1)
	for 0.000001 < abs(v1-v0) {
		v0 = v1
		v1 = v0 + x*powFloat64(x, -v0) - 1
	}
	return v1
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
			y = 1
		}
	} else {
		if n == 0 {
			panic("indeterminant form")
		}
	}
	return y
}

// powFloat64 returns x^y for real x and y. Some special cases cause panic.
func powFloat64(x, y float64) float64 {
	n := int(y)
	r := y - float64(n)
	m := int(round(1 / r))
	v := powInt(x, n)
	if r != 0 {
		if x < 0 && m%2 == 0 {
			panic("indeterminant form")
		}
		// v *= newton(func(z float64) float64 { return powInt(z, m) - x }, 1, 0.000001)
		v *= nthRoot(x, m) // Solves v^m - x = 0 for set x and m
	}
	return v
}

// sqrt returns x^0.5.
func sqrt(x float64) float64 {
	return nthRoot(x, 2)
}

// nthRoot returns x^(1/n). This is Newton's method applied to the specific problem of solving v^n-x = 0 for set x and n.
func nthRoot(x float64, n int) float64 {
	chgSign := false
	if x < 0 {
		if n%2 == 0 {
			panic("indeterminant form")
		}
		x = -x
		chgSign = true
	}
	tol := 0.000001
	v0, v1 := 0.0, 1.0
	for tol < abs(v1-v0) {
		v0 = v1
		v1 = v0 * (1 - 1/float64(n)*(1-x/powInt(v0, n)))
	}
	if chgSign {
		v1 = -v1
	}
	return v1
}

// newton finds a local root assuming f is smooth and continuous. Note, Newton's method is known to be unstable.
func newton(f func(x float64) float64, x0, tol float64) float64 {
	x1 := x0 + 2*tol
	for tol < abs(x1-x0) {
		x0 = x1
		x1 = x0 - f(x0)/diff(f, x0, tol)
	}
	return x1
}

// diff returns the approximate value of df/dx at x0.
func diff(f func(x float64) float64, x0, tol float64) float64 {
	return (f(x0+tol) - f(x0-tol)) / (2 * tol)
}

// abs returns |x|.
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
