package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//START OMIT
	data := ` {
	  "model_number": 20,
	  "count": 10,
	  "rating": 1.2
	}`

	var m map[string]json.Number
	err := json.Unmarshal([]byte(data), &m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", m)
	fmt.Println(m["count"].Int64())
	//END OMIT
}
