package tool

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func GetDb() *sql.DB {
	return Db
}

func init() {
	cfg := GetCfg().Database

	db, err := sql.Open(cfg.Driver, cfg.User+":"+cfg.Password+"@tcp("+cfg.Host+":"+cfg.Port+")/"+cfg.DbName+"?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		fmt.Println(err)
		return
	}
	Db = db
}
