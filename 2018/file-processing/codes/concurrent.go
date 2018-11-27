package main

import (
	"bufio"
	"os"
	"sync"
)

func main() {
	file, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// START
	c := make(chan string, 10)

	workerCount := 12
	wg := new(sync.WaitGroup)
	wg.Add(workerCount)
	// Starts the consumer

	for i := 0; i < workerCount; i++ {
		go consume(c, wg)
	}

	for scanner.Scan() {
		c <- scanner.Text()
	}

	wg.Done()
	// END
}

func consume(in <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range in {
		// just consume, do nothing
		_ = item
	}
}
