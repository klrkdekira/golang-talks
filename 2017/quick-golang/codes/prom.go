package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())

	err := http.ListenAndServe(httpAddr, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
