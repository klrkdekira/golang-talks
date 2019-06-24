package main

import (
	"encoding/json"
	"fmt"
)

//START A OMIT
type Simple struct {
	Message string `json:"msg"`
	Count   int    `json:"count"`
}

//END A OMIT

func main() {
	//START OMIT
	data := `
	{
		"msg":"hello, world",
		"count":10
	}`

	var simple *Simple
	err := json.Unmarshal([]byte(data), &simple)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", simple)
	fmt.Println(simple.Message)
	//END OMIT
}
