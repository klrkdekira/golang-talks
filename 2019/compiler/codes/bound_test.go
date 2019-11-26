package main

import "testing"

//START OMIT
func BenchmarkBound(b *testing.B) {
	list := []int{0, 1, 2, 3, 4}
	var sum int
	for i := 0; i < b.N; i++ {
		sum += list[0]
		sum += list[1]
		sum += list[2]
		sum += list[3]
		sum += list[4]
	}
	_ = sum
}

//END OMIT

//REMOVED OMIT
func BenchmarkBoundRemoved(b *testing.B) {
	list := []int{0, 1, 2, 3, 4}
	var sum int
	for i := 0; i < b.N; i++ {
		_ = list[4]
		sum += list[0]
		sum += list[1]
		sum += list[2]
		sum += list[3]
		sum += list[4]
	}
	_ = sum
}

//REMOVED_END OMIT
