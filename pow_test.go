package main

import (
	"testing"
)

func BenchmarkPowInt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		powInt(2, 10)
	}
}
