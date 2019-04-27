package main

// START IMPORT

import (
	"context"
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
		// context.WithValue(parent context.Context, key interface{}, val interface{})
		// message, ok := iobject.(string)
		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()

		wasteCycle(ctx)
		fmt.Fprint(w, "Hello, World")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// END MAIN

// START WASTE

func wasteCycle(ctx context.Context) {
	var i int64
	var s []int64
	for {
		select {
		case <-ctx.Done():
			log.Println("client disconnected")
			return
		default:
			s = append(s, i)
			i++
		}
	}
}

// END WASTE
