package controllers

/*
   @Auth: menah3m
   @Desc: 首页控制器
*/

type IndexController struct {
	AuthenticationController
}

func (c *IndexController) Index() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}
