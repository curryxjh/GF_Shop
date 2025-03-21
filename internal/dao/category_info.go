// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"GF_Shop/internal/dao/internal"
)

// internalCategoryInfoDao is an internal type for wrapping the internal DAO implementation.
type internalCategoryInfoDao = *internal.CategoryInfoDao

// categoryInfoDao is the data access object for the table category_info.
// You can define custom methods on it to extend its functionality as needed.
type categoryInfoDao struct {
	internalCategoryInfoDao
}

var (
	// CategoryInfo is a globally accessible object for table category_info operations.
	CategoryInfo = categoryInfoDao{
		internal.NewCategoryInfoDao(),
	}
)

// Add your custom methods and functionality below.
