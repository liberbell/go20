package models

import (
	"database/sql"

	"pj79/config"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const (
	tableNameUser = "users"
)

func init() {
	Db, err := sql.Open(config.Config.SQLDriver, config.Config.DbName)
}
