package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// START_a
	file, err := os.Open("2018/file-processing/codes/data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// END_a

	// START_b
	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	lines := bytes.Split(b, []byte("\n"))
	fmt.Printf("rows %d\n", len(lines))
	fmt.Printf("first row\n%s", lines[0])
	// END_b
}
