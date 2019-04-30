package main

import (
	"net/http"
	"time"
)

func main() {
	i := 0
	var err error
	for {
		// log.Println(i)
		if i == 0 {
			_, err = http.Get("http://127.0.0.1:8080/push")

		} else {
			_, err = http.Get("http://127.0.0.1:8080/pull")
		}
		if err != nil {
			panic(err)
		}

		time.Sleep(100 * time.Millisecond)
		i = (i + 1) % 2
	}
}
