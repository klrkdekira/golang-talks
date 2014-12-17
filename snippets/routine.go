package main

import (
	"fmt"
	"time"
)

func double(i int) {
	fmt.Printf("Doubled value %d to %d\n", i, i*2)
}

func main() {
	for i := 0; i < 10; i++ {
		go double(i)
	}
	time.Sleep(2 * time.Second)
}
