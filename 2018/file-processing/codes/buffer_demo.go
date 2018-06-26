package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("2018/file-processing/codes/data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// START
	c := make(chan string, 10)

	var i int
	for scanner.Scan() {
		fmt.Printf("item %d\n", i)
		c <- scanner.Text()
		i++
	}
	// END
}
