package main

import (
	"encoding/json"
	"fmt"
)

const document = `
{
	"msg":"hello, world",
	"count": 10,
	"shape_type": "circle",
	"shape": {
		"radius": 10
	}
}
`

func main() {
	var data interface{}
	err := json.Unmarshal([]byte(document), &data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", data)
	switch t := data.(type) {
	case map[string]interface{}:
		fmt.Println("map[string]interface{}")
		casted := data.(map[string]interface{})
		for k, v := range casted {
			fmt.Println(k)
			fmt.Println(v)
		}
	default:
		fmt.Println(t)
	}
}
