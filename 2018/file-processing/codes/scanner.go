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

	// START
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	fmt.Printf("first row\n%s", scanner.Text())
	// END
}
