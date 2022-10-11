package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"ming/controllers"
)

func init() {
	beego.Router("/", &controllers.IndexController{}, "get:Index")
	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.ReleaseController{})
}
