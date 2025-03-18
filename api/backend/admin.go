package backend

import "github.com/gogf/gf/v2/frame/g"

type AdminReq struct {
	g.Meta   `path:"/admin/add" tags:"Admin" method:"post" summary:"admin api"`
	Name     string `json:"name" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password" v:"required#密码不能为空" dc:"密码"`
	RoleIds  string `json:"role_ids"`
	IsAdmin  int    `json:"is_admin"`
}

type AdminRes struct {
	AdminId int `json:"admin_id"`
}

type AdminUpdateReq struct{}

type AdminUpdateRes struct{}

type AdminGetListCommonReq struct{}

type AdminGetListCommonRes struct{}

type AdminGetInfoReq struct{}

// jwt
//type AdminGetInfoRes struct {
//	Id          int    `json:"id"`
//	IdentityKey string `json:"identity_key"`
//	Payload     string `json:"payload"`
//}

// gtoken
type AdminGetInfoRes struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	RoleIds string `json:"role_ids"`
	IsAdmin int    `json:"is_admin"`
}
