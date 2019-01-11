package math

// GCD returns the greatest common divisor of two non-negative integers.
func GCD(a, b int) int {
	// TODO: Optimize using least squares or something...
	// TODO: Check for a,b <= 0
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM returns the least common multiple of two integers.
func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

// Factorial returns n! Panics if n < 0.
func Factorial(n int) int {
	if n < 0 {
		panic("integer n must be non-negative")
	}
	f := 1
	for ; 1 < n; n-- {
		f *= n
	}
	return f
}

// IsPrime reports whether a number is prime or not. Panics if n is less
// than two.
func IsPrime(n int) bool {
	if n < 2 {
		panic("integer n must be greater than one")
	}

	if n == 2 {
		return true
	}

	if n%2 == 0 {
		return false
	}

	for i := 3; i <= int(Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Factor returns a collection of each divisor of a positive integer
// mapped to the number of times each divisor divides said integer. For
// example, 12 = 2^2 * 3^1, so [1:1, 2:2, 3:1] would be returned. Note
// the integer n passed is returned as a Factor if and only if n is
// prime.
func Factor(n int) map[int]int {
	if n < 1 {
		panic("cannot factor non-positive integer")
	}
	factors := make(map[int]int)
	factors[1] = 1
	for n%2 == 0 {
		factors[2]++
		n /= 2
	}
	for i := 3; i < int(Round(Sqrt(float64(n)))); i += 2 {
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

// Max returns the maximum value of two numbers.
func Max(x, y float64) float64 {
	if x < y {
		return y
	}
	return x
}

// MaxIndex returns the index of the maximum value in an unordered list. Returns -1 if list is empty.
func MaxIndex(values []float64) int {
	n := len(values)
	if n == 0 {
		return -1
	}

	var m int
	for i := 1; i < n; i++ {
		if values[m] < values[i] {
			m = i
		}
	}
	return m
}

// Min returns the minimum value of two numbers.
func Min(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

// MinIndex returns the index of the minimum value in an unordered list. Returns -1 if list is empty.
func MinIndex(values []float64) int {
	n := len(values)
	if n == 0 {
		return -1
	}

	var m int
	for i := 1; i < n; i++ {
		if values[i] < values[m] {
			m = i
		}
	}
	return m
}
