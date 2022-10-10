package utils

/*
   @Auth: menah3m
   @Desc: 返回json数据
*/

type JsonMsg struct {
	Code int
	Msg  string
}

func Msg(code int, msg string) *JsonMsg {
	return &JsonMsg{
		Code: code,
		Msg:  msg,
	}
}
