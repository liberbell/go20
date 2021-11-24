package controllers

import (
	"net/http"
	"pj79/config"
)

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.HandleFunc("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", top)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
