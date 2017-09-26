package main

import "testing"

// START BENCH

func BenchmarkWasteCycle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wasteCycle()
	}
}

// END BENCH
