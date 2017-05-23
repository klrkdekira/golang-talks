package jsonip

import (
	"testing"
)

func TestIPv4(t *testing.T) {
	expected := "127.0.0.1"
	data, err := IPv4()
	if err != nil {
		t.Fatalf("Encountered unexpected error, %v", err)
	}

	if data.IP != expected {
		t.Fatalf("Expected %s, got %s", expected, data.IP)
	}
}
