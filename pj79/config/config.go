package config

import (
	"log"

	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port      string
	DbName    string
	SQLDriver string
	LogFile   string
}

var Config ConfigList

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}

	Config = ConfigList{
		Port: cfg.Section("web").Key("port").MustString("8080"),
	}
}
