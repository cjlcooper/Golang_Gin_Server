package Database

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

var SqlDB *sql.DB

func init() {
 var err error
 SqlDB, err = sql.Open("mysql", "root:cjl1992@tcp(127.0.0.1:3306)/TEST?charset=utf8")
 if err != nil {
  log.Fatal(err.Error())
 }
 err = SqlDB.Ping()
 if err != nil {
  log.Fatal(err.Error())
 }
}
