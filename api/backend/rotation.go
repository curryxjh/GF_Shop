package backend

import "github.com/gogf/gf/v2/frame/g"

type RotationReq struct {
	g.Meta `path:"/rotation/add" tags:"Rotation" method:"post" summary:"You first rotation api"`
	PicUrl string `json:"pic_url" v:"required#图像链接不能为空" dc:"图像链接"`
	Link   string `json:"link" v:"required#跳转链接不能为空" dc:"跳转链接"`
	Sort   int    `json:"sort" dc:"排序"`
}

type RotationRes struct {
	RotationId int `json:"rotation_id"`
}

type RotationDeleteReq struct {
	g.Meta `path:"/rotation/delete" method:"delete" tags:"Rotation" summary:"delete rotation api"`
	Id     uint `v:"min:1#请选择需要删除的轮播图" dc:"轮播图id"`
}
type RotationDeleteRes struct{}

type RotationUpdateReq struct {
	g.Meta `path:"/rotation/update/{Id}" method:"post" tags:"rotation" summary:"Update rotation api"`
	Id     uint   `json:"id" v:"min:1#请选择需要修改的轮播图" dc:"轮播图Id"`
	PicUrl string `json:"pic_url" v:"required#轮播图图片链接不能为空" dc:"图片链接"`
	Link   string `json:"link" v:"required#轮播图跳转链接不能为空" dc:"跳转链接"`
	Sort   int    `json:"sort" dc:"排序"`
}

type RotationUpdateRes struct {
	Id uint `json:"id"`
}

type RotationGetListCommonReq struct {
	g.Meta `path:"/rotation/list" method:"get" tags:"轮播图" summary:"List rotation api"`
	Sort   int `json:"sort" in:"query" dc:"排序类型"`
	CommonPaginationReq
}

type RotationGetListCommonRes struct {
	List  interface{} `json:"list" ds:"列表"`
	Page  int         `json:"page" ds:"分页码"`
	Size  int         `json:"size" ds:"分页数量"`
	Total int         `json:"total" ds:"数据总数"`
}
