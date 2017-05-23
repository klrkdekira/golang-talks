package main

import (
	"fmt"
	"net/http"
	"os"

	_ "expvar"
)

func main() {
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
