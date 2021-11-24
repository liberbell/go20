package controllers

import (
	"fmt"
	"net/http"
	"pj79/config"
)

func GenerateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("apps/views/templates/%s.html", file))
	}
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", top)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
