package main

import (
	"net/http"

	"github.com/codegangsta/negroni"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	// http.Handler
	mux := http.NewServeMux()
	mux.HandleFunc("/", SayHello)

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":8080")
}
