package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hello World.")
}

func top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tepl.html")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.ListenAndServe(":8000", nil)
}
