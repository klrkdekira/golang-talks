package main

// START IMPORT

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // HL
	"time"
)

// END IMPORT

// START MAIN

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		wasteCycle()
		fmt.Fprint(w, "Hello, World")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// END MAIN

// START WASTE

func wasteCycle() {
	timeout := time.After(60 * time.Second)
	var i int64
	var s []int64
	for {
		select {
		case <-timeout:
			return
		default:
			s = append(s, i)
			i++
		}
	}
}

// END WASTE
