package main

func main() {

}

// isPrime factors a positive integer and determines if it
// is prime or not.
func isPrime(n int) bool {
	if len(*factor(n)) == 2 {
		return true
	}
	return false
}

// factor returns a collection of each divisor of a positive
// integer mapped to the number of times each divisor
// divides said integer. See GeeksForGeeks.
func factor(n int) *map[int]int {
	if n == 0 {
		panic("cannot factor zero")
	}
	factors := map[int]int{1: 1}
	m := n
	for n%2 == 0 {
		factors[2]++
		n /= 2
	}
	for i := 3; i < int(round(sqrt(float64(n)))); i += 2 {
		for n%i == 0 {
			factors[i]++
			n /= i
		}
	}
	if 2 < n {
		factors[n]++
	}
	factors[m] = 1
	return &factors
}

// round returns x rounded to the nearest whole number as a float64.
func round(x float64) float64 {
	changedSign := false
	if x < 0 {
		x = -x
		changedSign = true
	}
	y := float64(int(x)) // Floor of x
	if 0.5 <= x-y {
		y++
	}
	if changedSign {
		return -y
	}
	return y
}

// roundTo rounds a number to the nth decimal place.
// For example, x = 0.123 becomes x.1 when n = 1.
func roundTo(x float64, n int) float64 {
	// TODO
	return x
}

// powInt returns x^n for any real x and any integer n.
func powInt(x float64, n int) float64 {
	var y float64
	if x != 0 {
		switch {
		case 0 < n:
			// Yacas' method
			y = 1
			for 0 < n {
				if n%2 != 0 {
					y *= x
				}
				x *= x
				n = n >> 1 // Divide by 2
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
	p := float64(n)
	for tol < abs(v1-v0) {
		v0 = v1
		v1 = ((p-1)*v0 + x) / p * powInt(v0, 1-n) // v0 * (1 - 1/float64(n)*(1-x/powInt(v0, n)))
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

// bisection: TODO
func bisection(f func(x float64) float64, x0, x1, tol float64) float64 {
	x := (x0 + x1) / 2.0
	y, y0, y1 := f(x), f(x0), f(x1)

	switch {
	case y == 0:
		return x
	case y0*y < 0:
		x = bisection(f, x0, x, tol)
	case y*y1 < 0:
		x = bisection(f, x, x1, tol)
	default:
		panic("no root on boundary")
	}

	return x
}

// diff returns the approximate value of df/dx at x0.
func diff(f func(x float64) float64, x0, tol float64) float64 {
	return (f(x0+tol) - f(x0-tol)) / (2 * tol)
}

// abs returns |x|.
func abs(x float64) float64 {
	// x = math.Abs(x) // Check out what the math package does: return Float64frombits(Float64bits(x) &^ (1 << 63))
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
