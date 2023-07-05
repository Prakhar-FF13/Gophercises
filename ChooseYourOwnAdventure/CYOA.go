package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	story := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}

	mux.HandleFunc("/", story)

	log.Fatal(http.ListenAndServe(":4000", mux))
}
