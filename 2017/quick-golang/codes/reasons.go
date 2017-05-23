package main

import "fmt"

func main() {
	prev := 0
	next := 1
	for i := 0; i < 10; i++ {
		now := prev + next
		fmt.Println(now)

		prev, next = next, now
	}
}
