package controllers

import (
	"log"
	"net/http"
	"pj79/apps/models"
)

func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		GenerateHTML(w, nil, "layout", "public_navbar", "signup")
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user := models.User{
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
		}
		if err := user.CreateUser(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/", 302)
	}

}

func login(w http.ResponseWriter, r *http.Request) {
	GenerateHTML(w, nil, "layout", "public_navbar", "login")
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	user, err := models.GetUserByEmail(r.PostFormValue("email"))
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/login", 302)
	}
	if user.Password == models.Encrypt(r.PostFormValue(password)) {
		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
		}
		cockie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cockie)
		http.Redirect(w, r, "/", 302)
	}
}
