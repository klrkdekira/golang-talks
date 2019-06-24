package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	simple := `{"msg":"hello, world"}`

	var m map[string]string
	err := json.Unmarshal([]byte(simple), &m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", m)
}
