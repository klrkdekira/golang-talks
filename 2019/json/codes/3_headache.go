package main

import (
	"encoding/json"
	"fmt"
)

//START A OMIT
type (
	Circle struct {
		Radius int    `json:"radius"`
		Color  string `json:"color"`
	}

	Message struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	Combined struct {
		*Circle
		*Message
	}
)

//END A OMIT

func main() {
	//START B OMIT
	data := `{"title":"hello","content":"hello, world", "color":"red"}`

	combined := new(Combined)
	err := json.Unmarshal([]byte(data), &combined)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", combined)
	fmt.Println(combined.Title)
	fmt.Println(combined.Radius)
	//END B OMIT
}
