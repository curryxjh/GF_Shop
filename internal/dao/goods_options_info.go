// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"GF_Shop/internal/dao/internal"
)

// internalGoodsOptionsInfoDao is an internal type for wrapping the internal DAO implementation.
type internalGoodsOptionsInfoDao = *internal.GoodsOptionsInfoDao

// goodsOptionsInfoDao is the data access object for the table goods_options_info.
// You can define custom methods on it to extend its functionality as needed.
type goodsOptionsInfoDao struct {
	internalGoodsOptionsInfoDao
}

var (
	// GoodsOptionsInfo is a globally accessible object for table goods_options_info operations.
	GoodsOptionsInfo = goodsOptionsInfoDao{
		internal.NewGoodsOptionsInfoDao(),
	}
)

// Add your custom methods and functionality below.
