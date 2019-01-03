package main

import (
	"fmt"
	"math"
)

// Note: accuracy =/= precision. A precise, wrong answer is not accurate.
// An accurate, imprecise answer is a lucky guess that may not be
// duplicated or may fail under other circumstances. Here, all values
// are precise to the constant tol. Accuracy is only guarenteed by the
// algorithm of choice.

const (
	// Pi is the ratio of a circle's circumfrence to its diameter.
	Pi = float64(3.141592653589793)
	// E is Euler's number, the natural rate of growth of the exponential
	// function.
	E = float64(2.718281828459045)
	// LN2 is the natural logarithm of two (ln(2)).
	LN2 = float64(0.693147180559945)
	// tol (1e-16) is the tolerance required by all numerical functions.
	// All numerical methods must converge to a value determined by this
	// tolerance. This does not guarentee the value returned is correct
	// (accurate) to the number of digits represented here, but that the
	// value returned cannot be made more precise with the method used.
	tol = 5e-16
)

func main() {
	// x := 2.9237425767113105
	// actual := ln(x)
	// expected := math.Log(x)
	// err := abs(actual - expected)
	// table := tabby.New()
	// table.AddHeader("Function", "Input", "Value")
	// table.AddLine("ln", x, actual)
	// table.AddLine("math.Log", x, expected)
	// table.AddLine("Abs Error", "", err)
	// table.Print()

	x := 1234
	fmt.Println(
		roundUpToBase(int(x), int(math.Pow10(int(math.Log10(float64(x))-1)))), // Equivalent to roundToMag10(x,2)
	) // 1300

	for i := 0; i < 5; i++ {
		fmt.Println(roundToMag10(x, i))
	}
}

// roundToMag10 rounds x to the next power n of 10. Panics if n is
// negative. A few examples include:
// 1. x = 1234, n = 0 -->  1235
// 2. x = 1234, n = 1 -->  1240
// 3. x = 1234, n = 3 -->  2000
// 4. x = 1234, n = 5 --> 10000
func roundToMag10(x, n int) int {
	if n < 0 {
		panic("n must be non-negative")
	}
	return roundUpToBase(x, int(math.Pow10(n)))
}

// returns the highest power of 10 such that x >= 10^n
func mag10(x int) int {
	var n int
	return int(math.Log10(float64(x)))
}

// powInt returns x^n for any real x and any integer n.
func powInt(x float64, n int) float64 {
	if x == 0 {
		if n == 0 {
			panic("indeterminant form")
		}
		return 0 // Case 0^n = 0 for all n
	}

	y := 1.0
	if 0 < n {
		// Yacas' method
		for 0 < n {
			if n%2 != 0 {
				y *= x
			}
			x *= x
			n /= 2
		}
	} else if n < 0 {
		return 1 / powInt(x, -n)
	}
	return y
}

// pow returns x^y for real x and y.
func pow(x, y float64) float64 {
	// x^y = x^n * x^r
	// 1. Compute x^n
	n := int(y)
	r := y - float64(n) // 0 <= r < 1
	v := powInt(x, n)

	// 2. Compute x^r
	if tol < r {
		return v * exp(r*ln(x))
	}
	return v
}

// exp returns e^x.
func exp(x float64) float64 {
	// Calculate 2^k
	k := int(x / LN2)
	r := x - float64(k)*LN2
	v0 := powInt(2, k)
	if r == 0 {
		return v0
	}

	// GeeksForGeeks solution for set n
	// TODO: Choose smallest possible n to value that minimizes error
	v1 := 1.0
	for n := 20.0; 0 < n; n-- {
		v1 = 1 + r*v1/n
	}
	return v0 * v1
}

// sqrt returns +x^0.5.
func sqrt(x float64) float64 {
	return nthRoot(x, 2)
}

// nthRoot returns x^(1/n).
func nthRoot(x float64, n int) float64 {
	var chgSign bool // Indicates sign of x
	if x < 0 {
		if n%2 == 0 {
			panic("indeterminant form")
		}
		x = -x
		chgSign = true
	}

	// Solve v^n - x = 0 using Newton-Raphson's method
	v0 := 0.0
	v1 := 1.0
	p := float64(n)
	for tol < abs(v1-v0) {
		v0 = v1
		v1 = 1 / p * (v0*(p-1.0) + x/powInt(v0, n-1))
	}

	if chgSign {
		return -v1
	}
	return v1
}

// ln returns the natural logarithm of base e.
func ln(x float64) float64 {
	switch {
	case x == E:
		return 1 // e^1 = e <==> ln(e) = 1
	case x == 1:
		return 0 // e^0 = 1 <==> ln(1) = 0
	case x <= 0:
		panic("ln is undefined for non-positive values") // e^x = y > 0 for all x <==> ln(y) is defined for all y > 0
	}

	// Decompose ln(x) = k ln(2) + ln(y), where k is the largest integer
	// such that 2^k <= x and y is a real number such that 0 <= y < 1.
	// Compute k ln(2)
	var k int
	p := 1.0 // p = 2^k <= x
	for ; p <= x; p *= 2 {
		if x == p {
			return float64(k) * LN2 // x = p = 2^k
		}
		k++
	}

	// Compute ln(y) using Newton-Raphson's method. y is on the range
	// [1,2).
	y := x / powInt(2, k) // 1 <= y < 2
	v0 := 0.0
	v1 := 1.5 // Initialize v1 as the center of [1,2)
	for tol < abs(v1-v0) {
		v0 = v1
		v1 = v0 + y/exp(v0) - 1
	}
	return float64(k)*LN2 + v1
}

// log returns the base b-logarithm of x.
func log(x, b float64) float64 {
	return ln(x) / ln(b)
}

func log10(x float64) float64 {
	return ln(x) / ln(10)
}

func log2(x float64) float64 {
	return ln(x) / ln(2)
}

// newton (also newton-raphson) finds a local root assuming f is smooth
//  and continuous. Note, Newton's method is known to be unstable.
func newton(f func(x float64) float64, x0 float64) float64 {
	x1 := x0 + 2*tol
	for tol < abs(x1-x0) {
		x0 = x1
		x1 = x0 - f(x0)/diff(f, x0)
		if x0 == x1 {
			panic("x0 is a fixed point") // Fixed points map to themselves: f(x) = x
		}
	}
	return x1
}

// bisection: TODO
func bisection(f func(x float64) float64, x0, x1 float64) float64 {
	x := (x0 + x1) / 2.0
	y, y0, y1 := f(x), f(x0), f(x1)

	switch {
	case abs(y) < tol:
		return x
	case y0*y < 0:
		return bisection(f, x0, x)
	case y*y1 < 0:
		return bisection(f, x, x1)
	default:
		panic("no root on boundary")
	}
}

// diff returns the approximate value of df/dx at x0.
func diff(f func(x float64) float64, x0 float64) float64 {
	return (f(x0+tol) - f(x0-tol)) / (2 * tol)
}

// abs returns |x|.
func abs(x float64) float64 {
	// math.Abs(x) does this: return Float64frombits(Float64bits(x) &^ (1 << 63))
	if x < 0 {
		return -x
	}
	return x
}

// gcd returns the greatest common divisor of two non-negative integers.
func gcd(a, b int) int {
	// TODO: Optimize using least squares or something...
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// lcm returns the least common multiple of two integers.
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// factorial returns n! Panics if n < 0.
func factorial(n int) int {
	if n < 0 {
		panic("integer n must be non-negative")
	}
	f := 1
	for ; 1 < n; n-- {
		f *= n
	}
	return f
}

// isPrime reports whether a number is prime or not. Panics if n is less
// than two.
func isPrime(n int) bool {
	if n < 2 {
		panic("integer n must be greater than one")
	}

	if n == 2 {
		return true
	}

	if n%2 == 0 {
		return false
	}

	for i := 3; i <= int(sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// factor returns a collection of each divisor of a positive integer
// mapped to the number of times each divisor divides said integer. For
// example, 12 = 2^2 * 3^1, so [1:1, 2:2, 3:1] would be returned. Note
// the integer n passed is returned as a factor if and only if n is
// prime.
func factor(n int) map[int]int {
	if n < 1 {
		panic("cannot factor non-positive integer")
	}
	factors := make(map[int]int)
	factors[1] = 1
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
	if 1 < n {
		factors[n] = 1
	}
	return factors
}

// round returns x rounded to the nearest whole number as a float64.
func round(x float64) float64 {
	changedSign := false
	if x < 0 {
		x = -x
		changedSign = true
	}

	y := floor(x)
	if 0.5 <= x-y {
		y++
	}

	if changedSign {
		return -y
	}
	return y
}

// roundTo rounds a number to the nth decimal place. For example,
// x = 0.123 becomes x.12 when n = 2. Special cases are n = 0 and n < 0.
// If n = 0, then 1 is returned. If n < 0, then 0 is returned.
func roundTo(x float64, n int) float64 {
	v := powInt(10, n)
	return floor(x*v) / v
}

// roundUpToBase rounds a number up to the next multiple of b. For
// example, given x = 21, b = 5, 25 would be returned. Panics if base is
// not positive.
func roundUpToBase(n, b int) int {
	if b <= 0 {
		// TODO: Consider b < 0 or n < 0
		panic("base b must be positive")
	}
	return n + b - (n % b)
}

// floor returns the largest integer value less than x as a float64. For
// example, the floor of 1.9 is 1.0.
func floor(x float64) float64 {
	// TODO: Check for x < 0 case
	return float64(int(x))
}

// ceiling returns the smallest integer greater than x as a float64. If
// x is an integer value, then x is returned, not x+1. For example, the
// ceiling of 1.1 is 2.0, but the ceiling of 1.0 is 1.0.
func ceiling(x float64) float64 {
	y := floor(x)
	return y + floor(x-y)
}
