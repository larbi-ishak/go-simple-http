package main

import (
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static/"))
	http.Handle("/", fileServer))

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)
	
}
