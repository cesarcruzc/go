package main

import (
	"fmt"
	"log"
	"net/http"
)

type messageHandler struct {
	message string
}

func (m messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.message)
}

func main() {
	mux := http.NewServeMux()
	m1 := &messageHandler{message: "<h1>Hola desde un Handler 1</h1>"}
	m2 := &messageHandler{message: "<h1>Hola desde un Handler 2</h1>"}
	m3 := &messageHandler{message: "<h1>Hola desde un Handler 3</h1>"}

	mux.Handle("/greet", m1)
	mux.Handle("/say-goodbye", m2)
	mux.Handle("/talk", m3)
	log.Println("Server started on http://localhost:8080")
	log.Println(http.ListenAndServe(":8080", mux))
}
