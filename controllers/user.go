package controllers

import (
	"ming/models"
)

/*
   @Auth: menah3m
   @Desc: 用户控制器
*/

//	UserController 用户管理控制器
type UserController struct {
	AuthenticationController
}

// List 获取用户列表
func (c *UserController) List() {
	//鉴权

	//q为查询参数
	q := c.GetString("q")

	users := models.GetUserList(q)
	c.Data["users"] = users
	c.Data["q"] = q
	c.TplName = "user/index.html"
}
