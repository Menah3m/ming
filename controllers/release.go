package controllers

/*
   @Auth: menah3m
   @Desc: 发布相关模块 控制器
*/

type ReleaseController struct {
	AuthenticationController
}

func (c *ReleaseController) ServiceList() {
	c.Data["name"] = "发布相关功能的首页"
	c.TplName = "release/index.html"
}
