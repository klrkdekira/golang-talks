package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func main() {
	//START A OMIT
	data := `{"title":"hello","content":"hello, world"}`

	var m map[string]interface{}
	err := json.Unmarshal([]byte(data), &m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", m)
	//END A OMIT

	//START B OMIT
	fields := make([]reflect.StructField, 0)
	for k, v := range m {
		fields = append(fields, reflect.StructField{
			Name: strings.Title(k),
			Type: reflect.TypeOf(v),
			Tag:  reflect.StructTag(fmt.Sprintf("json:\"%s\"", k)),
		})
	}

	typ := reflect.StructOf(fields)
	v := reflect.New(typ).Elem()
	newStruct := v.Addr().Interface()

	err = json.Unmarshal([]byte(data), &newStruct)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", newStruct)
	//END B OMIT
}
