package main

import (
	"log"
	"net/http"
)

func main() {
	// Multiplexor
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))
	// Handler
	mux.Handle("/", fs)
	log.Println("Server started on http://localhost:8080")
	log.Println(http.ListenAndServe(":8080", mux))
}
