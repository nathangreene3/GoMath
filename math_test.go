package main

import "testing"

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
		t.FailNow()
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

func BenchmarkPowYacas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		powYacas(100, 100)
	}
}
