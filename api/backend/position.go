package backend

import "github.com/gogf/gf/v2/frame/g"

type PositionReq struct {
	g.Meta    `path:"/position/add" tags:"Position" method:"post" summary:"position add api"`
	PicUrl    string `json:"pic_url" v:"required#图片链接不能为空" dc:"图片链接"`
	Link      string `json:"link" v:"required#跳转链接为空" dc:"跳转链接"`
	GoodsName string `json:"goods_name" v:"required#商品名不能为空" dc:"商品名称"`
	GoodsId   uint   `json:"goods_id" v:"required#商品Id不能为kon" dc:"商品Id"`
	Sort      int    `json:"sort" dc:"排序"`
}

type PositionRes struct {
	PositionId int `json:"position_id"`
}

type PositionDeleteReq struct {
	g.Meta `path:"/position/delete" method:"delete" tags:"手工位图" summary:"position delete api"`
	Id     uint `json:"id" v:"min:1#请选择需要删除的手工位图" dc:"手工位图id"`
}

type PositionDeleteRes struct{}

type PositionUpdateReq struct {
	g.Meta    `path:"/position/update/{Id}" method:"post" tags:"手工位图" summary:"修改手工位图接口"`
	Id        uint   `json:"id"      v:"min:1#请选择需要修改的手工位图" dc:"手工位图Id"`
	PicUrl    string `json:"pic_url" v:"required#手工位图图片链接不能为空" dc:"图片链接"`
	Link      string `json:"link"    v:"required#跳转链接不能为空" dc:"跳转链接"`
	Sort      int    `json:"sort"    dc:"跳转链接"`
	GoodsName string `json:"goods_name" v:"required#商品名称不能为空" dc:"商品名称"` //冗余设计
	GoodsId   uint   `json:"goods_id"  v:"required#商品Id不能为空" dc:"商品ID"`  //mysql三范式
}

type PositionUpdateRes struct {
	Id uint `json:"id"`
}

type PositionGetListCommonReq struct {
	g.Meta `path:"/position/list" method:"get" tags:"手工位图" summary:"手工位列表接口"`
	Sort   int `json:"sort" in:"query" dc:"排序类型"`
	CommonPaginationReq
}

type PositionGetListCommonRes struct {
	List  interface{} `json:"list" dc:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
