package controllers

import (
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"ming/common/errors"
	"ming/forms"
	"ming/models"
	"net/http"
	"strings"
	"time"
)

/*
   @Auth: menah3m
   @Desc: 认证控制器
*/

type AuthController struct {
	BaseController
	o orm.Ormer
}

//认证登录
func (c *AuthController) Login() {
	//检查session，如果已经登录，则直接返回用户
	sessionKey := beego.AppConfig.DefaultString("auth::SessionKey", "user")
	indexAction := beego.AppConfig.DefaultString("auth::IndexAction", "IndexController.Get")
	sessionUser := c.GetSession(sessionKey)
	if sessionUser != nil {

		c.Redirect(beego.URLFor(indexAction), http.StatusFound)
		return
	}
	form := &forms.LoginForm{}
	errs := errors.New()
	// POST请求，则验证
	if c.Ctx.Input.IsPost() {
		// 获取用户数据
		remember := strings.TrimSpace(c.GetString("remember"))
		if err := c.ParseForm(form); err != nil {
			errs.AddError("default", "参数解析错误")
		} else {
			// 验证（此处可自定义验证逻辑）
			user := models.GetUserByUsername(form.Username)
			if user == nil {
				// 用户不存在
				errs.AddError("default", "用户名或密码错误")
			} else if !user.ValidPassword(form.Password) {
				// 密码不正确
				errs.AddError("default", "用户名或密码错误")
			} else {
				// 验证通过
				//如果点击了 记住我 按钮，则保存userid到cookie
				if remember == "on" {

					c.Ctx.SetCookie("username", user.Username, time.Second*60*60*24)

				} else {
					c.Ctx.SetCookie("username", user.Username, -1)
				}

				// 记录用户状态（记录到session）
				c.SetSession("user", user.ID)

				// 将sessionID 写入cookie 以便于返回给client，以后每次请求都可以携带该sessionID

				// 如果cookie为空，则不让查询，如果有cookie并验证用户是已经登录的状态，则允许查询
				c.Redirect(beego.URLFor(indexAction), http.StatusFound)

			}
		}
		c.Data["form"] = form
		c.Data["errors"] = errs

	}
	// GET请求，则返回页面
	if c.Ctx.Request.Method == "GET" {
		//先判断cookie中是否有用户数据
		username := c.Ctx.GetCookie("username")

		if username != "" {
			c.Data["username"] = username
			c.Data["checked"] = "checked"
		} else {
			c.Data["username"] = ""
			c.Data["checked"] = ""
		}

	}
	c.TplName = "auth/login.html"
}

func (c *AuthController) Signup() {
	// POST请求，则验证
	if c.Ctx.Input.IsPost() {

		// 验证成功
		// 验证失败
	}
	// 其他方法请求，则返回页面
	c.TplName = "auth/signup.html"
}

func (c *AuthController) Logout() {
	loginAction := beego.AppConfig.DefaultString("auth::LogoutAction", "AuthController.Login")
	c.DestroySession()
	c.Redirect(beego.URLFor(loginAction), http.StatusFound)
}
