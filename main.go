package main

import (
	"fmt"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello, world!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	fmt.Println("Starting server at http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Server failed: %s\n", err)
	}

}
