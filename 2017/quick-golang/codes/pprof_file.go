package main

import (
	"fmt"
	"os"
	"runtime/pprof"
)

func main() {
	file, err := os.Create("cpu.pprof")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	pprof.StartCPUProfile(file)
	defer pprof.StopCPUProfile()
}
