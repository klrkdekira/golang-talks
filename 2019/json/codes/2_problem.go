package main

import (
	"encoding/json"
	"fmt"
)

//START A OMIT
const payloadA = `{
		"msg": "hello, world",
		"count": 10,
		"shape_type": "circle",
		"shape": {
			"radius": 10
		}}`

const payloadB = `{
		"msg": "hello, world",
		"count": 10,
		"shape_type": "square",
		"shape": {
			"height": 10,
			"width": 10
		}}`

type Shape struct {
	Message   string                 `json:"msg"`
	Count     int                    `json:"count"`
	ShapeType string                 `json:"shape_type"`
	Shape     map[string]interface{} `json:"shape"`
}

//END A OMIT

func main() {
	//START B OMIT
	circle := new(Shape)
	err := json.Unmarshal([]byte(payloadA), &circle)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", circle)

	// Square here
	square := new(Shape)
	err = json.Unmarshal([]byte(payloadB), &square)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", square)
	fmt.Println(square.Shape["radius"])
	//END B OMIT
}
