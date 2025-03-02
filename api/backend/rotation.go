package backend

import "github.com/gogf/gf/v2/frame/g"

type RotationReq struct {
	g.Meta `path:"/rotation/add" tags:"Rotation" method:"post" summay:"You first rotation api"`
	PicUrl string `json:"pic_url" v:"required#图像链接不能为空" dc:"图像链接"`
	Link   string `json:"link" v:"required#跳转链接不能为空" dc:"跳转链接"`
	Sort   int    `json:"sort" dc:"排序"`
}

type RotationRes struct {
	RotationId int `json:"rotation_id"`
}
