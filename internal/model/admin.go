package model

import "github.com/gogf/gf/v2/os/gtime"

type AdminCreateUpdateBase struct {
	Name     string
	Password string
	RoleIds  string
	UserSalt string
	IsAdmin  int
}

type AdminCreateInput struct {
	AdminCreateUpdateBase
}

type AdminCreateOutput struct {
	AdminId int `json:"admin_id"`
}

type AdminUpdateInput struct {
	AdminCreateUpdateBase
	Id uint
}

type AdminGetListInput struct {
	Page int
	Size int
	Sort int
}

type AdminGetListOutput struct {
	List  []AdminGetListOutputItem `json:"list" description:"列表"`
	Page  int                      `json:"page" description:"分页码"`
	Size  int                      `json:"size" description:"分页数量"`
	Total int                      `json:"total" description:"数据总数"`
}

type AdminGetListOutputItem struct {
	Id         int         `json:"id"`
	Name       string      `json:"name"`
	RoleIds    string      `json:"role_ids"`
	IsAdmin    int         `json:"is_admin"`
	Created_at *gtime.Time `json:"created_at"`
	Updated_at *gtime.Time `json:"updated_at"`
}

type AdminSearchOutputItem struct {
	AdminGetListOutputItem
}
