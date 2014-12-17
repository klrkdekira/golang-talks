package main

import "fmt"

// START OMIT
func main() {
	func() {
		fmt.Println("Hello, World")
	}()
}

// END OMIT
