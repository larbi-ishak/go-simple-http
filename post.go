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

	// Create an HTTP request with the URL, method (POST), and request body
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}

	// Set headers if needed
	req.Header.Set("Content-Type", "application/json")

	// Make the request and handle the response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Handle the response, e.g., check status code and read response body
	if resp.StatusCode == http.StatusCreated {
		// Request was successful
		// You can read the response body using resp.Body
		body, _ := io.ReadAll(resp.Body)
		// Successfully read the response body
		fmt.Println("Request was successful!")
		fmt.Println("Response Body:", string(body))
	} else {
		// Request failed, handle the error
		fmt.Println("Request failed with status code:", resp.Status)
	}
}
