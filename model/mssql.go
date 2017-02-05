package model

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-adodb"
)

var msdb *sqlx.DB

func InitMssql(dialect, conn string) (err error) {
	msdb, err = sqlx.Connect(dialect, conn)
	return
}
