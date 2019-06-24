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

//END A OMIT

//START B OMIT
type (
	Shape struct {
		Message   string          `json:"msg"`
		Count     int             `json:"count"`
		ShapeType string          `json:"shape_type"`
		Shape     json.RawMessage `json:"shape"`
	}

	Circle struct {
		Radius int `json:"radius"`
	}

	Square struct {
		Height int `json:"height"`
		Width  int `json:"width"`
	}
)

//END B OMIT

//START D OMIT
func parseShape(shape *Shape) error {
	if shape.ShapeType == "circle" {
		var circle *Circle
		err := json.Unmarshal(shape.Shape, &circle)
		if err != nil {
			return err
		}
		fmt.Printf("%#v\n", circle)
		// do work here
	} else if shape.ShapeType == "square" {
		var square *Square
		err := json.Unmarshal(shape.Shape, &square)
		if err != nil {
			return err
		}
		fmt.Printf("%#v\n", square)
		// do work here
	}
	return nil
}

//END D OMIT

func main() {
	//START C OMIT
	shape := new(Shape)
	err := json.Unmarshal([]byte(payloadA), &shape)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", shape)
	err = parseShape(shape)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(payloadB), &shape)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", shape)
	parseShape(shape)
	if err != nil {
		panic(err)
	}
	//END C OMIT
}
