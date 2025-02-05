package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	start := server.ListenAndServe()
	if start != nil {
		fmt.Printf("Error starting server: %v", start)
	}
}
