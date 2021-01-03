package main

import (
	"fmt"
	"net/http"
)

func setupRoutes() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello Chat")
	})
}

func main() {
	fmt.Printf("Hello Chat")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
