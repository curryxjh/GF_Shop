package model

import "github.com/gogf/gf/v2/os/gtime"

// RotationCreateUpdateBase 创建/修改内容基类
type RotationCreateUpdateBase struct {
	PicUrl string
	Link   string
	Sort   int
}

// RotationCreateInput 创建内容
type RotationCreateInput struct {
	RotationCreateUpdateBase
}

// RotationCreateOutput 创建内容返回结果
type RotationCreateOutput struct {
	RotationId int `json:"rotation_id"`
}
type RotationUpdateInput struct {
	RotationCreateUpdateBase
	Id uint
}
type RotationGetListInput struct {
	Page int
	Size int
	Sort int
}

type RotationGetListOutput struct {
	List  []RotationGetListOutputItem `json:"list" dc:"列表"`
	Page  int                         `json:"page" dc:"分页码"`
	Size  int                         `josn:"size" dc:"分页数量"`
	Total int                         `json:"total" dc:"数据总数"`
}

type RotationSearchInput struct {
	Key      string
	Type     string
	Category uint
	Page     int
	Size     int
	Sort     int
}

type RotationGetListOutputItem struct {
	Id        uint        `json:"id"`
	PicUrl    string      `json:"pic_url"`
	Link      string      `json:"link"`
	Sort      int         `json:"sort"`
	CreatedAt *gtime.Time `json:"created_at"`
	UpdatedAt *gtime.Time `json:"updated_at"`
}

type RotationSearchOutputItem struct {
	RotationGetListOutputItem
}
