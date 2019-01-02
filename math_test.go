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
func TestLn(t *testing.T) {

}

func BenchmarkDivBy2v1(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n = math.MaxInt64
		n /= 2
	}
}

func BenchmarkDivBy2v2(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n = math.MaxInt64
		n = n / 2
	}
}

func BenchmarkBitShiftBy1(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n = math.MaxInt64
		n = n >> 1
	}
}

func TestPowInt(t *testing.T) {
	var result float64
	var expected float64

	tolerance := 0.000000001

	result = powInt(2, 3)
	expected = 8
	if tolerance < abs(expected-result) {
		t.Fatalf("expected %v, received %v\n", expected, result)
		t.FailNow()
	}

	result = powInt(-2, 3)
	expected = -8
	if tolerance < abs(expected-result) {
		t.Fatalf("expected %v, received %v\n", expected, result)
		t.FailNow()
	}

	result = powInt(2, -3)
	expected = 1.0 / 8.0
	if tolerance < abs(expected-result) {
		t.Fatalf("expected %v, received %v\n", expected, result)
		t.FailNow()
	}

	result = powInt(-2, -3)
	expected = -1.0 / 8.0
	if tolerance < abs(expected-result) {
		t.Fatalf("expected %v, received %v\n", expected, result)
		t.FailNow()
	}

	result = powInt(0, 3)
	expected = 0
	if tolerance < abs(expected-result) {
		t.Fatalf("expected %v, received %v\n", expected, result)
		t.FailNow()
	}

	result = powInt(2, 0)
	expected = 1
	if tolerance < abs(expected-result) {
		t.Fatalf("expected %v, received %v\n", expected, result)
		t.FailNow()
	}

	result = powInt(2.1, 3)
	expected = 2.1 * 2.1 * 2.1
	if tolerance < abs(expected-result) {
		t.Fatalf("expected %v, received %v\n", expected, result)
	}
}

func TestPow(t *testing.T) {
	var x, y float64
	var result float64
	var expected float64

	tolerance := 0.000000001
	for i := 0; i < 1000; i++ {
		x, y = rand.ExpFloat64(), rand.ExpFloat64()
		result = pow(x, y)
		expected = math.Pow(x, y)
		if tolerance < abs(expected-result) {
			t.Fatalf("i = %d input x = %v, y = %v, expected %v, received %v\n", i, x, y, expected, result)
		}
	}
}

func TestExp(t *testing.T) {
	var x float64
	var result float64
	var expected float64

	tolerance := 0.000000001
	for i := 0; i < 1000; i++ {
		x = math.Pow(-1, float64(rand.Intn(2)+1)) * rand.ExpFloat64()
		result = exp(x)
		expected = math.Exp(x)
		if tolerance < abs(expected-result) {
			t.Fatalf("i = %d input x = %v, expected %v, received %v\n", i, x, expected, result)
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
