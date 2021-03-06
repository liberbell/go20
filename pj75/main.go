package main

import (
	"html/template"
	"log"
	"net/http"
)

type MyHandler struct{}

// func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello World.")
// }

func top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tepl.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, "Hello Go World.")
}

func main() {
	http.HandleFunc("/top", top)
	http.ListenAndServe(":8000", nil)
}
