package main

import (
	_ "ming/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

