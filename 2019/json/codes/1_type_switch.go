package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//START OMIT
	data := `{
	  "msg":"hello, world",
	  "count":10
	}`

	var m map[string]interface{}
	err := json.Unmarshal([]byte(data), &m)
	if err != nil {
		panic(err)
	}

	for _, v := range m {
		switch t := v.(type) {
		case int:
			fmt.Println("This is int")
		case string:
			fmt.Println("This is string")
		default:
			fmt.Printf("No idea what this is, value %v\n", t)
		}
	}
	//END OMIT
}
