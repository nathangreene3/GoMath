package math

// Abs returns |x|.
func Abs(x float64) float64 {
	// math.Abs(x) does this: return Float64frombits(Float64bits(x) &^ (1 << 63))
	if x < 0 {
		return -x
	}
	return x
}

// Round returns x rounded to the nearest whole number as a float64.
func Round(x float64) float64 {
	changedSign := false
	if x < 0 {
		x = -x
		changedSign = true
	}

	y := Floor(x)
	if 0.5 <= x-y {
		y++
	}

	if changedSign {
		return -y
	}
	return y
}

// RoundTo rounds a number to the nth decimal place. For example,
// x = 0.123 becomes x.12 when n = 2. Special cases are n = 0 and n < 0.
// If n = 0, then 1 is returned. If n < 0, then 0 is returned.
func RoundTo(x float64, n int) float64 {
	v := PowInt(10, n)
	return Floor(x*v) / v
}

// RoundToMag10 rounds x to the power n of 10. Panics if n is
// negative. A few examples include:
// 	1. x = 1234, n = 0 -->  1235 (rounds to 10^0 =     1)
// 	2. x = 1234, n = 1 -->  1240 (rounds to 10^1 =    10)
// 	3. x = 1234, n = 3 -->  2000 (rounds to 10^3 =  1000)
// 	4. x = 1234, n = 5 --> 10000 (rounds to 10^5 = 10000)
func RoundToMag10(x, n int) int {
	if n < 0 {
		panic("n must be non-negative")
	}
	return RoundUpToBase(x, int(Pow10(n)))
}

// OrderMag10 returns the order of magnitude (largest power) n of 10 such that x >= 10^n.
func OrderMag10(x int) int {
	return int(Log10(float64(x)))
}

// OrderMag10 returns the order of magnitude (largest power) n of 10 such
// that x >= 10^n. Panics if x <= 0.
// func OrderMag10(x float64) float64 {
// 	return float64(int(Log10(x)))
// }

// RoundUpToBase rounds a number up to the next multiple of b. For
// example, given x = 21, b = 5, 25 would be returned. Panics if base is
// not positive.
func RoundUpToBase(n, b int) int {
	if b <= 0 {
		// TODO: Consider b < 0 or n < 0
		panic("base b must be positive")
	}
	return n + b - (n % b)
}

// NumDigits returns the number of digits of any positive n in base b.
func NumDigits(n, b int) int {
	return int(Log(float64(n), float64(b))) + 1
}

// Trunc rounds x to the nearest integer value toward zero.
func Trunc(x float64) float64 {
	return float64(int(x))
}

// Floor returns the largest integer value less than x as a float64. For
// example, the Floor of 1.9 is 1.0.
func Floor(x float64) float64 {
	// a := abs(x)
	return float64(int(x))
}

// Ceiling returns the smallest integer greater than x as a float64. If
// x is an integer value, then x is returned, not x+1. For example, the
// Ceiling of 1.1 is 2.0, but the Ceiling of 1.0 is 1.0.
func Ceiling(x float64) float64 {
	// TODO: Check for x < 0 case
	y := Floor(x)
	return y + Floor(x-y)
}
