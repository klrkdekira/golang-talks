package jsonip

import (
	"testing"
)

// For the sake of example
func BenchmarkIPv4(b *testing.B) {
	setupServer(`{"ip": "127.0.0.1"}`)
	for i := 0; i < b.N; i++ {
		data, err := IPv4()
		if err != nil {
			b.Fatalf("Encountered unexpected error, %v", err)
		}
		_ = data
	}
}
