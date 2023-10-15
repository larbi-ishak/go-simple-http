package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// you can't write to the response after a redirect
		// you can't redirect once you've written to the response

		// fmt.Fprintf(w, "Hello World")
		// NOTE redirect
		http.Redirect(w, r, "/head", 301)
	})

	http.HandleFunc("/head", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "another route ")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
