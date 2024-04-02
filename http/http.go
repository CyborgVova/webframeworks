package main

import (
	"fmt"
	"log"
	"net/http"
	mw "webframeworks/http/middleware"
)

// type Handler interface {
//     ServeHTTP(ResponseWriter, *Request)
// }

// type HandlerFunc func(ResponseWriter, *Request)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.UserAgent())
}

func main() {
	mux := http.NewServeMux()
	// http.Handle("/hello", mw.Logging(http.HandlerFunc(Hello)))
	mux.HandleFunc("/hello", Hello)
	fmt.Println("Starting server localhost:8080 ...")

	handler := mw.Logging(mux)
	log.Fatal(http.ListenAndServe("localhost:8080", handler))
}
