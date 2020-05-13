package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello From Golang!")
	})
	s := &http.Server{
		Addr:        ":8080",
		Handler:     r,
		ReadTimeout: 10 * time.Second,
	}
	fmt.Println("Starting server on port 8080")
	log.Fatal(s.ListenAndServe())
}