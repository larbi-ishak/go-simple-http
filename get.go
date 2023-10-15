package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// NOTE: GET request
	resp, err := http.Get("https://quotes.toscrape.com/")
	if err != nil {
		fmt.Println(err)
	}
	// The caller must close the response body when finished with it:
	defer resp.Body.Close()

	// read the response body
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

// or in a simple way

// client := &http.client{}
// resp, _ := client.get("http://www.facebook.com")
//
// data, _ := io.readall(resp.body)
// fmt.println(string(data))
