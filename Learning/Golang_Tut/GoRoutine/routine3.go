package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Echo back the request method and URL path
	response := fmt.Sprintf("Method: %s, Path: %s", r.Method, r.URL.Path)
	// Write the response to the client
	w.Write([]byte(response))
}

func main() {
	// Register the handler function for the "/" route
	http.HandleFunc("/", handler) // function to be called when an HTTP request is made to the root path ("/") of the server.

	// Start the HTTP server on port 8080
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Println(err)
		}
	}()
	// func() { ... }: This is an anonymous function definition. When go is used before a function call, it instructs Go to run that function concurrently in a new goroutine.
	//parentheses () which represent the function call. This is necessary syntax in Go to execute the function.

	// Print a message indicating that the server has started
	fmt.Println("Server started on :8080")

	// Block the main goroutine indefinitely
	// This is done to keep the program running so that the HTTP server can handle incoming requests
	select {}
}
