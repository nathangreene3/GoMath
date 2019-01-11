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
	n := int(y)
	return PowInt(x, n) * Exp((y-float64(n))*Ln(x)) // x^y = x^n * x^r for integer n and r on range [0,1)
}

// Exp returns e^x.
func Exp(x float64) float64 {
	// Calculate 2^k
	k := int(x / Ln2)
	r := x - float64(k)*Ln2
	if r == 0 {
		return PowInt(2, k)
	}

	// GeeksForGeeks solution for set n
	// TODO: Choose smallest possible n that minimizes error
	v := 1.0
	for n := 20.0; 0 < n; n-- {
		v = 1 + r*v/n
	}
	return PowInt(2, k) * v
}

// Sqrt returns +x^0.5.
func Sqrt(x float64) float64 {
	return NthRoot(x, 2)
}

// NthRoot returns x^(1/n). Panics if n is zero or x is negative and n is
// even simultaneously.
func NthRoot(x float64, n int) float64 {
	if n == 0 {
		panic("indeterminant form")
	}
	// TODO: Figure out what the largest n is before overflow occurs
	// if int(Pow2(1023)) <= n {
	// 	// fmt.Println(int(Pow2(1023)), n)
	// 	return 1
	// }

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
	e0 := 0.0
	e1 := 1.0
	p := float64(n)
	for tol < Abs(e1-e0) {
		e0 = e1
		v0 = v1
		v1 = 1 / p * (v0*(p-1) + x/PowInt(v0, n-1))
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
		panic("undefined for non-positive values") // e^x = y > 0 for all x <==> ln(y) is defined for all y > 0
	}

	// Decompose ln(x) = n ln(2) + ln(r), where n is the largest integer
	// such that 2^n <= x and r is a real number such that 0 <= r < 1.
	// Compute n ln(2)
	var n int
	p := 1.0 // p = 2^k <= x
	for ; p <= x; p *= 2 {
		if x == p {
			return float64(n) * Ln2 // x = p = 2^k, r = 1
		}
		n++
	}

	// For r close to one, this works well.
	// Compute ln(r) = 2*artanh((r-1)/(r+1)), 0 < r < 2 (if r = 1, then n*Ln2 would have already been returned)
	r := x / PowInt(2, n)
	r2 := r*r + 1
	v := 0.0
	q := 1.0 - 2/(r+1)       // (r - 1) / (r + 1) // r = x / 2^n
	q2 := 1.0 - 4*r/(r2+2*r) // q^2 = (r2-2*r)/(r2+2*r)
	for k := 1.0; k < 25; k += 2 {
		v += q / k
		q *= q2
	}
	return float64(n)*Ln2 + 2*v
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

// Choose returns n-choose-k or n!/((n-k)!*k!). 0 <= k <= n and 0 <= n.
func Choose(n, k int) int {
	switch {
	case k < 0:
		panic("cannot choose negative quantity k")
	case n < k:
		panic("cannot choose k greater than n")
	case n < 0:
		panic("cannot choose from negative quantity n")
	}
	return Factorial(n) / (Factorial(n-k) * Factorial(k))
}

// binomCoefs: TODO
func binomCoefs(n int) []int {
	coef := make([]int, 0, n+1)
	for i := 0; i < n/2; i++ {
		coef = append(coef, Choose(n, i))
	}
	if n%2 == 0 {
		coef = append(coef, Choose(n, n/2))
	}
	for i := 0; i < n/2; i++ {
		coef = append(coef, coef[n-i+n/2])
	}
	return coef
}
