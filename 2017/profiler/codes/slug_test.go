package main

import (
	"context"
	"testing"
)

// START BENCH

func BenchmarkWasteCycle(b *testing.B) {
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		wasteCycle(ctx)
	}
}

// END BENCH
