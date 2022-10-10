package models

/*
   @Auth: menah3m
   @Desc:
*/

import (
	"database/sql"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func init() {
	dbHost, _ := beego.AppConfig.String("mysql::dbHost")
	dbPort, _ := beego.AppConfig.String("mysql::dbPort")
	dbUser, _ := beego.AppConfig.String("mysql::dbUser")
	dbPass, _ := beego.AppConfig.String("mysql::dbPass")
	dbName, _ := beego.AppConfig.String("mysql::dbName")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=PRC&parseTime=true",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)
	db, _ = sql.Open("mysql", dsn)

	if err := db.Ping(); err != nil {
		log.Fatal(err)

	}
}
