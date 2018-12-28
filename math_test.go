package main

import (
	"math"
	"math/rand"
	"testing"
)

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

func TestPowFloat64(t *testing.T) {
	var x, y float64
	var result float64
	var expected float64

	tolerance := 0.000000001
	for i := 0; i < 1000; i++ {
		x, y = rand.ExpFloat64(), rand.ExpFloat64()
		result = powFloat64(x, y)
		expected = math.Pow(x, y)
		if tolerance < abs(expected-result) {
			t.Fatalf("i = %d input x = %v, y = %v, expected %v, received %v\n", i, x, y, expected, result)
		}
	}
}

func BenchmarkPowInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		powInt(100, 100)
	}
}

func BenchmarkPowFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		powFloat64(100, 100)
	}
}
