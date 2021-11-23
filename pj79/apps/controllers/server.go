package controllers

import (
	"net/http"
	"pj79/config"
)

func StartMainServer() error {
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
