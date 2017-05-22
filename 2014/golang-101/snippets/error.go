package main

import (
	"fmt"
	"strconv"
)

// START OMIT
func printError(i string) {
	fmt.Printf("converting string (%s) to int\n", i)
	result, err := strconv.Atoi(i)
	if err != nil {
		fmt.Printf("error encountered, %v\n", err)
	}
	fmt.Printf("converted result, %d\n", result)
}

func ignoreError(i string) {
	fmt.Printf("converting string (%s) to int\n", i)
	result, _ := strconv.Atoi(i)
	fmt.Printf("converted result, %d\n", result)
}

func main() {
	printError("123")
	ignoreError("abc")
}

// END OMIT
