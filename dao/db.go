package dao

import (
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

//_ "github.com/denisenkom/go-mssqldb"
//_ "github.com/go-sql-driver/mysql"

var Db *xorm.Engine

func InitDB(dialect, conn string) {
	var err error

	Db, err = xorm.NewEngine(dialect, conn)
	if err != nil {
		log.Fatal(err)
	}

	//Db.ShowSQL(true)
	Db.SetColumnMapper(core.SameMapper{})
	//Db.Sync2(new(Fruit))
}
