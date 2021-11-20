package models

import (
	"database/sql"

	"github.com/tommy-muehle/go-mnd/config"
)

var Db *sql.DB

var err error

const (
	tableNameUser = "users"
)

func init() {
	Db, err := sql.Open(config.Config.SQLDriver, config.Config.DbName)
}
