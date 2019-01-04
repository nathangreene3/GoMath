package math

// PowInt returns x^n for any real x and any integer n.
func PowInt(x float64, n int) float64 {
	switch {
	case x == 0:
		if n == 0 {
			panic("indeterminant form") // 0^0 is undefined
		}
		return 0 // 0^n = 0 for all n
	case n < 0:
		return 1 / PowInt(x, -n) // x^-n = 1/x^n while x != 0
	case n == 0:
		return 1 // x^0 = 1 while x != 0
	}

	// Yacas' method
	// x^n where x != 0 and n > 0
	y := 1.0
	for ; 0 < n; n /= 2 {
		if n%2 != 0 {
			y *= x
		}
		x *= x
	}
	return y
}

// Pow2 returns 2^n for any integer n.
func Pow2(n int) float64 {
	if n < 0 {
		return 1 / Pow2(-n)
	}

	// Yaca's method
	// This is equivalent to powInt(2,n), but skips the special cases
	v := 1.0
	x := 2.0
	for ; 0 < n; n /= 2 {
		if n%2 != 0 {
			v *= x
		}
		x *= x
	}
	return v
}

// Pow10 returns 10^n for any integer n.
func Pow10(n int) float64 {
	// TODO: pre-compute 10^n as math.Pow10 does
	if n < 0 {
		return 1 / Pow10(-n)
	}

	// Yaca's method
	// This is equivalent to powInt(10,n), but skips the special cases
	v := 1.0
	x := 10.0
	for ; 0 < n; n /= 2 {
		if n%2 != 0 {
			v *= x
		}
		x *= x
	}
	return v
}

// Pow returns x^y for real x and y.
func Pow(x, y float64) float64 {
	// x^y = x^n * x^r for integer n and r on range [0,1)
	// Compute x^n
	n := int(y)
	r := y - float64(n) // 0 <= r < 1
	if r == 0 {
		return PowInt(x, n) // y is an integer
	}

	// Compute x^r = e^(r ln x)
	return PowInt(x, n) * Exp(r*Ln(x))
}

// Exp returns e^x.
func Exp(x float64) float64 {
	// Calculate 2^k
	k := int(x / Ln2)
	r := x - float64(k)*Ln2
	v0 := PowInt(2, k)
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

// Sqrt returns +x^0.5.
func Sqrt(x float64) float64 {
	return NthRoot(x, 2)
}

// NthRoot returns x^(1/n).
func NthRoot(x float64, n int) float64 {
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
	for tol < Abs(v1-v0) {
		v0 = v1
		v1 = 1 / p * (v0*(p-1.0) + x/PowInt(v0, n-1))
	}

	if chgSign {
		return -v1
	}
	return v1
}

// Ln returns the natural logarithm of base e.
func Ln(x float64) float64 {
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
			return float64(k) * Ln2 // x = p = 2^k
		}
		k++
	}

	// Compute ln(y) using Newton-Raphson's method. y is on the range
	// (1,2), not [1,2) since we would have already returned the case
	// x = 2^k.
	y := x / PowInt(2, k) // 1 < y < 2
	v0 := 0.0
	v1 := -1.0 // v1 will be negative as ln(y) < 0 for all y on (0,1)
	for tol < Abs(v1-v0) {
		v0 = v1
		v1 = v0 + y/Exp(v0) - 1
	}
	return float64(k)*Ln2 + v1
}

// Log returns the base b-logarithm of x.
func Log(x, b float64) float64 {
	return Ln(x) / Ln(b)
}

// Log10 returns the base-10 logarithm of x.
func Log10(x float64) float64 {
	return Ln(x) / Ln(10)
}

// Log2 returns the base-2 logarithm of x.
func Log2(x float64) float64 {
	return Ln(x) / Ln(2)
}