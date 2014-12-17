package main

import "fmt"

// START OMIT
func determine(input interface{}) {
	switch input.(type) {
	case string:
		fmt.Printf("%v is string\n", input)
	case int:
		fmt.Printf("%v is int\n", input)
	default:
		fmt.Printf("%v is unknown\n", input)
	}
}

func main() {
	var x, y interface{}
	x = "Hello, World!"
	y = 42

	determine(x)
	determine(y)
}

// END OMIT
