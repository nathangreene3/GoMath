package main

import (
	"math"
	"math/rand"
	"testing"

	"github.com/cheynewallace/tabby"
)

// TODO: Add panic test scenarios

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
		{6, 720},
		{7, 5040},
		{8, 40320},
		{9, 362880},
		{10, 3628800},
	}
	var actual int
	for i := range data {
		actual = factorial(data[i].input)
		if actual != data[i].ans {
			table := tabby.New()
			table.AddHeader("Function", "Input", "Value")
			table.AddLine("exp", data[i].input, actual)
			table.AddLine("", "", data[i].ans)
			table.Print()
			t.Fatal("\n")
		}
	}
}

func TestPowInt(t *testing.T) {
	var x float64
	var n int
	var actual float64
	var expected float64
	var err float64

	for i := 0; i < 1000; i++ {
		x, n = rand.ExpFloat64(), rand.Int() // Cant test negative x using math.Pow
		actual = powInt(x, n)
		expected = math.Pow(x, float64(n))
		err = abs(expected - actual)
		if 10*tol < err {
			table := tabby.New()
			table.AddHeader("Function", "Input", "Value")
			table.AddLine("powInt", x, actual)
			table.AddLine("math.Pow", x, expected)
			table.AddLine("Abs Error", "", err)
			table.Print()
			t.Fatal("\n")
		}
	}
}

func TestPow(t *testing.T) {
	var x, y float64
	var actual float64
	var expected float64
	var err float64

	for i := 0; i < 1000; i++ {
		x, y = rand.ExpFloat64(), rand.ExpFloat64()
		actual = pow(x, y)
		expected = math.Pow(x, y)
		err = abs(expected - actual)
		if 10*tol < err {
			table := tabby.New()
			table.AddHeader("Function", "Input", "Value")
			table.AddLine("pow", x, actual)
			table.AddLine("math.Pow", x, expected)
			table.AddLine("Abs Error", "", err)
			table.Print()
			t.Fatal("\n")
		}
	}
}

func TestExp(t *testing.T) {
	var x float64
	var actual float64
	var expected float64
	var err float64

	for i := 0; i < 1000; i++ {
		x = math.Pow(-1, float64(rand.Intn(2)+1)) * rand.ExpFloat64()
		actual = exp(x)
		expected = math.Exp(x)
		err = abs(expected - actual)
		if 10*tol < err {
			table := tabby.New()
			table.AddHeader("Function", "Input", "Value")
			table.AddLine("exp", x, actual)
			table.AddLine("math.Exp", x, expected)
			table.AddLine("Abs Error", "", err)
			table.Print()
			t.Fatal("\n")
		}
	}
}

func TestLn(t *testing.T) {
	var x float64
	var actual float64
	var expected float64
	var err float64

	for i := 0; i < 5; i++ {
		x = rand.ExpFloat64()
		actual = ln(x)
		expected = math.Log(x)
		err = abs(expected - actual)
		if 10*tol < err {
			table := tabby.New()
			table.AddHeader("Function", "Input", "Value")
			table.AddLine("ln", x, actual)
			table.AddLine("math.Log", x, expected)
			table.AddLine("Abs Error", "", err)
			table.Print()
			t.Fatal("\n")
		}
	}
}

func BenchmarkMathPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.Pow(100, 100)
	}
}

func BenchmarkPowInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		powInt(100, 100)
	}
}

func BenchmarkMathPow10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.Pow10(100)
	}
}

func BenchmarkPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pow(100, 100)
	}
}

func BenchmarkPow2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pow2(100)
	}
}

func BenchmarkPow10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pow10(100)
	}
}
