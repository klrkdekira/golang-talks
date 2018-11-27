package main

import (
	"bytes"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	lines := bytes.Split(b, []byte("\n"))
	for i := 0; i < len(lines); i++ {
		item := lines[i]
		_ = item
	}
}
