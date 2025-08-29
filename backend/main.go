package main

import (
	"log"
	"net/http"
)

func main() {
	// Initialize DB
	InitDB()

	// Routes
	http.Handle("/", http.FileServer(http.Dir("../frontend")))
	http.HandleFunc("/api/shorten", ShortenHandler)
	http.HandleFunc("/", RedirectHandler)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
