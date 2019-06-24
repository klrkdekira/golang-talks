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
	fmt.Printf("%#v\n", m)
	//END OMIT
}
