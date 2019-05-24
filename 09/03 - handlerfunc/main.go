package main

import (
	"fmt"
	"log"
	"net/http"
)

// Controller
func messageHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hola handlerfunc</h1>")
}

// Pass parameters
func messageHFCustom(message string) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, message)
		},
	)
}

func main() {
	mux := http.NewServeMux()
	// Routes
	mux.HandleFunc("/", messageHandlerFunc)
	mux.HandleFunc("/hello", messageHFCustom("<h1>Hello</h1>"))
	mux.HandleFunc("/bye", messageHFCustom("<h1>Bye</h1>"))

	log.Println("Server started on http://localhost:8080")
	log.Println(http.ListenAndServe(":8080", mux))
}
