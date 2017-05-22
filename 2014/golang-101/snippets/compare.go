package main

import "fmt"

// START OMIT
func compareString(x, y string) string {
	if x > y {
		return x
	}
	return y
}

func compareInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Printf("%v won!\n", compareString("a", "b"))
	fmt.Printf("%v won!\n", compareInt(5, 10))
}

// END OMIT
