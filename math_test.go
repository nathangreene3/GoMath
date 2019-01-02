package main

import (
	"math"
	"math/rand"
	"testing"
)

func TestFactorial(t *testing.T) {
	data := []struct {
		input int
		ans   int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
	}
	var result int
	for i := range data {
		result = factorial(data[i].input)
		if result != data[i].ans {
			t.Fatalf("expected %d, received %d\n", data[i].ans, result)
		}
	}
}

func TestPowInt(t *testing.T) {
	var x float64
	var n int
	var result float64
	var expected float64

	for i := 0; i < 1000; i++ {
		x, n = rand.ExpFloat64(), rand.Int() // Cant test negative x using math.Pow
		result = powInt(x, n)
		expected = math.Pow(x, float64(n))
		if tol <= abs(expected-result) {
			if expected < 0 || result < 0 {
				t.Fatalf("powInt(%v)\nexpected %0.15f\nreturned %0.15f\n   error  %0.15f\n", x, expected, result, abs(result-expected))
			}
			t.Fatalf("powInt(%v)\nexpected %0.15f\nreturned %0.15f\n   error %0.15f\n", x, expected, result, abs(result-expected))
		}
	}
}

func TestPow(t *testing.T) {
	var x, y float64
	var result float64
	var expected float64

	for i := 0; i < 1000; i++ {
		x, y = rand.ExpFloat64(), rand.ExpFloat64()
		result = pow(x, y)
		expected = math.Pow(x, y)
		if tol <= abs(expected-result) {
			if expected < 0 || result < 0 {
				t.Fatalf("pow(%v)\nexpected %0.15f\nreturned %0.15f\n   error  %0.15f\n", x, expected, result, abs(result-expected))
			}
			t.Fatalf("pow(%v)\nexpected %0.15f\nreturned %0.15f\n   error %0.15f\n", x, expected, result, abs(result-expected))
		}
	}
}

func TestExp(t *testing.T) {
	var x float64
	var result float64
	var expected float64

	for i := 0; i < 1000; i++ {
		x = math.Pow(-1, float64(rand.Intn(2)+1)) * rand.ExpFloat64()
		result = exp(x)
		expected = math.Exp(x)
		if tol < abs(expected-result) {
			t.Fatalf("exp(%0.15f) expected %0.15f, returned %0.15f\n", x, expected, result)
		}
	}
}
func TestLn(t *testing.T) {
	var x float64
	var result float64
	var expected float64

	for i := 0; i < 1000; i++ {
		x = rand.ExpFloat64()
		result = ln(x)
		expected = math.Log(x)
		if tol <= abs(expected-result) {
			if expected < 0 || result < 0 {
				t.Fatalf("ln(%v)\nexpected %0.15f\nreturned %0.15f\n   error  %0.15f\n", x, expected, result, abs(result-expected))
			}
			t.Fatalf("ln(%v)\nexpected %0.15f\nreturned %0.15f\n   error %0.15f\n", x, expected, result, abs(result-expected))
		}
	}
}

func BenchmarkPowInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		powInt(100, 100)
	}
}

func BenchmarkPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pow(100, 100)
	}
}
