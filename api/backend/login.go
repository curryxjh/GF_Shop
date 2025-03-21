package backend

import "github.com/gogf/gf/v2/frame/g"

type LoginDoReq struct {
	g.Meta   `path:"/login" method:"post" summary:"执行登录请求" tags:"登录"`
	Name     string `json:"name" v:"required#请输入账号" summary:"执行登录请求" tags:"登录"`
	Password string `json:"password" v:"required#请输入密码" dc:"密码(明文)"`
}

// for jwt
type LoginDoRes struct {
	Info interface{} `json:"info"`
	//Referer string `json:"referer" dc:"引导客户端跳转地址"`
	//Token  string    `json:"token"`
	//Expire time.Time `json:"expire"`
}

// for gtoken
type LoginRes struct{}
