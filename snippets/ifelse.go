package main

import "fmt"

// START OMIT
func main() {
	x := "x"
	if x == "x" {
		fmt.Println("It is x!")
	} else if x == "y" {
		fmt.Println("It is y!")
	} else {
		fmt.Println("Unknown")
	}
}

// END OMIT
