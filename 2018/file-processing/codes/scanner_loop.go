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

	for i := 0; i < 10 && scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
	}
	// END
}
