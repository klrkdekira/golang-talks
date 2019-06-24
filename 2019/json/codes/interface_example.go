package main

import (
	"fmt"
)

func main() {
	fmt.Println(parse(1))
	fmt.Println(parse("hello"))
	fmt.Println(parse([]string{"one"}))
}

func parse(i interface{}) string {
	switch i.(type) {
	case int:
		return "this is integer"
	case string:
		return "this is string"
	default:
		return "value with unhandled type found"
	}
}
