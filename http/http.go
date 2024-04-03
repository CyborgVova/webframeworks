package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	mw "webframeworks/middleware"
)

// type Handler interface {
//     ServeHTTP(ResponseWriter, *Request)
// }

// type HandlerFunc func(ResponseWriter, *Request)

func Hello(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		name := r.URL.Query().Get("name")
		if name != "" {
			fmt.Fprintf(w, "Hello, %s\n", name)
			return
		}
	case http.MethodPost:
		b, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("error reading body request", err)
		}
		var user map[string]interface{}
		json.Unmarshal(b, &user)
		if name, ok := user["name"].(string); ok {
			fmt.Fprintf(w, "Hello, %s\n", name)
			return
		}
		log.Println("empty request")
	}
	fmt.Fprintln(w, "Hello, Stranger !")
}

func main() {
	// mux := http.NewServeMux()
	// mux.HandleFunc("/hello", Hello)
	// fmt.Println("Starting server localhost:8080 ...")
	// log.Fatal(http.ListenAndServe("localhost:8080", mw.Logging(mux)))

	http.Handle("/hello", mw.Logging(http.HandlerFunc(Hello)))
	http.Handle("/auth", mw.Logging(mw.Authorization(http.HandlerFunc(Hello))))

	fmt.Println("Starting server localhost:8080 ...")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
