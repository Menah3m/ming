package forms

/*
   @Auth: menah3m
   @Desc:  认证模块 表单
*/

//用户登录表单
type LoginForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
