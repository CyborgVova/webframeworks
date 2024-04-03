package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	mw "webframeworks/middleware"

	"github.com/gorilla/mux"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-type", "application/json")
	switch r.Method {
	case http.MethodGet:
		if name := r.URL.Query().Get("name"); name != "" {
			fmt.Fprintf(w, "Hello, %s\n", name)
			return
		}
	case http.MethodPost:
		b, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("error reading body request:", err)
			fmt.Fprintln(w, "Hello, Stranger")
			return
		}
		name := &struct{ Name string }{}
		err = json.Unmarshal(b, name)
		if err != nil {
			log.Println("error serialization:", err)
			fmt.Fprintln(w, "Hello, Stranger")
			return
		}
		fmt.Fprintf(w, "Hello, %s\n", name.Name)
		return
	}

	fmt.Fprintln(w, "Hello, Stranger")
}

func main() {
	router := mux.NewRouter()
	router.StrictSlash(true)

	router.HandleFunc("/hello", Hello).Methods("GET", "POST")

	router.Use(mw.Logging)

	fmt.Println("Starting server on localhost:8080 ...")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
