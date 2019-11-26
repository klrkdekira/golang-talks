package main

import (
	"testing"
)

//START OMIT
var p *int64
var q int64

func BenchmarkHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			a := int64(i)
			p = &a
		}
	}
}

func BenchmarkStack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			a := int64(i)
			q = a
		}
	}
}

//END OMIT
