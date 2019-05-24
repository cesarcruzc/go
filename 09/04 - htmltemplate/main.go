package main

import (
	"html/template"
	"log"
	"net/http"
)

//Person contains person's data
type Person struct {
	Name string
	Age  uint8
}

func render(w http.ResponseWriter, r *http.Request) {
	p := &Person{"Cesar", 28}
	t, err := template.ParseFiles("./views/index.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	t.Execute(w, p)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", render)

	log.Println("Server started on http://localhost:8080")
	log.Println(http.ListenAndServe(":8080", mux))
}
