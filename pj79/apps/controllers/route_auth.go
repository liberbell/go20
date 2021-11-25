package controllers

import "net/http"

func signup(w http.ResponseWriter, r *http.Request) {
	GenerateHTML(w, nil, "layout", "public_navbar", "signup")
}
