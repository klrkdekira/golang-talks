package main

import "fmt"

func inc(i int) {
	i++
}

func incPtr(i *int) {
	*i++
}

func main() {
	x := 10
	fmt.Printf("initial value: %d\n", x)
	inc(x)
	fmt.Printf("final value: %d\n", x)

	y := 10
	fmt.Printf("initial value: %d\n", y)
	incPtr(&y)
	fmt.Printf("final value: %d\n", y)
}
