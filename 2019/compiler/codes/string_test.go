package main

import (
	"strings"
	"testing"
)

//START OMIT
func BenchmarkStringConcat(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s += "foo"
	}
	_ = s
}

func BenchmarkStringBuilder(b *testing.B) {
	var s strings.Builder
	for i := 0; i < b.N; i++ {
		s.WriteString("foo")
	}
	_ = s.String()
}

//END OMIT
