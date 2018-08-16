package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println("32:", isPrime(32))
	//displayFactorization(32)
	printCollatz(5)
}

func newton(f func(x float64) float64, x0 float64) float64 {
	var (
		e  float64
		x1 float64
	)
	e = 0.001
	x1 = x0 + e
	for e < abs(x1-x0) {
		x1 = x0 - diff(f, x0)/f(x0)
	}
	return x1
}

func integrate(f func(x float64) float64, x0, x1 float64) float64 {
	return 0
}

func diff(f func(x float64) float64, x0 float64) float64 {
	return 0
}

func abs(x float64) float64 {
	var y float64
	if x < 0 {
		y = -x
	} else {
		y = x
	}
	return y
}

func factor(n int) map[int]int {
	f := make(map[int]int)
	for n%2 == 0 {
		if _, exists := f[2]; exists {
			f[2]++
		} else {
			f[2] = 1
		}
		n /= 2
	}
	for i := 3; i < int(math.Sqrt(float64(n)+1)); i += 2 {
		for n%i == 0 {
			if _, exists := f[i]; exists {
				f[i]++
			} else {
				f[i] = 1
			}
			n /= i
		}
	}
	if 2 < n {
		f[n] = 1
	}
	return f
}

func gcd(a, b int) int {
	return 0
}

func displayFactorization(n int) {
	//f := factor(n)
	fmt.Print(n, " = ")
	// for key, value := range f {
	// 	//
	// }
}

func isPrime(n int) bool {
	return len(factor(n)) == 1
}

func nthPrime(n int) int {
	//var p int

	return 0
}

func collatz(n int) map[int]int {
	c := make(map[int]int)
	c[0] = n
	i := 0
	for 1 < n {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		c[i] = n
		i++
	}
	return c
}

func printCollatz(n int) {
	c := collatz(n)
	for _, value := range c {
		fmt.Print(value, " ")
	}
	fmt.Println()
}
