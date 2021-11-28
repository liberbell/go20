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
	GenerateHTML(w, "HeLLo", "layout", "public_navbar", "top")
}

func index(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		GenerateHTML(w, nil, "layout", "index", "private_navbar")
	}
}
