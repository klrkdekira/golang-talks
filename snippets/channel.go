package main

import "fmt"

func main() {
	c := make(chan int, 5)
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c)

	for i := 0; i < 5; i++ {
		fmt.Printf("Popping value %d\n", <-c)
	}
}
