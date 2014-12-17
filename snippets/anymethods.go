package main

import "fmt"

type Level int

func (l Level) Info() string {
	return fmt.Sprintf("Currently at level %d", l)
}

func main() {
	var i Level = 10
	fmt.Println(i.Info())

	var j Level = 20
	fmt.Println(j.Info())
}
