package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
)

/*
   @Auth: menah3m
   @Desc:  鉴权控制器
*/

type AuthenticationController struct {
	BaseController
}

func (c *AuthenticationController) Prepare() {
	// 鉴权工作
	sessionKey := beego.AppConfig.DefaultString("auth::SessionKey", "user")
	loginAction := beego.AppConfig.DefaultString("auth::LoginAction", "AuthController.Login")
	sessionUser := c.GetSession(sessionKey)
	// 如果获取不到session信息，则跳转到登录页面
	if sessionUser == nil {
		c.Redirect(c.URLFor(loginAction), http.StatusFound)
		c.StopRun()
	}
}
