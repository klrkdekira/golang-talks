package main

import "fmt"

func sayHello(msg chan<- string) {
	msg <- "Hello, 世界!"
}

func main() {
	message := make(chan string, 1)
	go sayHello(message)
	fmt.Println(<-message)
}
