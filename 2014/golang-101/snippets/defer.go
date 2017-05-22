package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	defer func() {
		t := time.Now().String()
		fmt.Printf("Finished at %v\n", t)
	}()

	fmt.Println("Working...")
}

// END OMIT
