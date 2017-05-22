package main

import "net/http"

func SayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	// http.Handler
	mux := http.NewServeMux()
	mux.HandleFunc("/", SayHello)

	http.ListenAndServe(":8080", mux)
}
