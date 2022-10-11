package main

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"

	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"log"
	_ "ming/routers"
)

func main() {

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

	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册数据库
	orm.RegisterDataBase("default", "mysql", dsn)
	//测试数据库连接
	if db, err := orm.GetDB("default"); err != nil {
		log.Fatal(err)
	} else if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	// 打开debug模式
	orm.Debug = true

	beego.Run()
}
