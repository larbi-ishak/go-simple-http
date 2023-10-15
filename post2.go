package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Create a Person struct and marshal it into JSON
	data := Person{
		Name: "John",
		Age:  30,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	// Set the URL for the POST request (JSONPlaceholder example)
	url := "https://jsonplaceholder.typicode.com/users"
	resp, _ := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Response Body:", string(body))
}
