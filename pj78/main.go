package main

import "gopkg.in/go-ini/ini.v1"

type ConfigList struct {
	Port      int
	DbName    string
	SQLDriver string
}

var Config ConfigList

func init() {
	cfg, _ := ini.Load("config.ini")
}

func main() {

}
