package controllers

import (
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	// t, err := template.ParseFiles("apps/views/templates/top.html")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// t.Execute(w, "Hello")
	GenerateHTML(w, "HeLLo", "layout", "top")
}