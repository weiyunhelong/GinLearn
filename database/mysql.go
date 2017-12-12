package database

import (
 "database/sql"
 _ "github.com/go-sql-driver/mysql"
 "log"
)

var SqlDB *sql.DB

func init() {
 var err error
 SqlDB, err = sql.Open("mysql", "root:123456@tcp(192.168.1.87:3306)/test?parseTime=true")
 if err != nil {
  log.Fatal(err.Error())
 }
 err = SqlDB.Ping()
 if err != nil {
  log.Fatal(err.Error())
 }
}

