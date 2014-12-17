package main

import "fmt"

// START OMIT
func inventory(i int) func() int {
	// Double
	return func() int {
		return i * 2
	}
}

func main() {
	fmt.Println(inventory(10)())
}

// END OMIT
